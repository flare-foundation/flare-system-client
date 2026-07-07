package protocol

import (
	"math"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/flare-foundation/go-flare-common/pkg/payload"
	"github.com/stretchr/testify/require"
)

// SignatureSubmitterDataVerifier carries a real off-by-one boundary
// (MaxUint16-104) and a hard length-must-be-38 invariant. Both worth pinning.

func TestSignatureSubmitterDataVerifier_OkRequiresLength38(t *testing.T) {
	for _, length := range []int{0, 1, 37, 39, 100} {
		t.Run("len="+strString(length), func(t *testing.T) {
			err := SignatureSubmitterDataVerifier(&SubProtocolResponse{
				Status: payload.Ok,
				Data:   make([]byte, length),
			})
			require.Error(t, err)
			require.Contains(t, err.Error(), "is not 38")
		})
	}
}

func TestSignatureSubmitterDataVerifier_AdditionalDataBoundary(t *testing.T) {
	// 1 (type) + 38 (message) + 65 (signature) = 104. AdditionalData may
	// not exceed MaxUint16 - 104. Off-by-one trap.
	maxAllowed := math.MaxUint16 - 104

	t.Run("accepts exactly MaxUint16-104", func(t *testing.T) {
		err := SignatureSubmitterDataVerifier(&SubProtocolResponse{
			Status:         payload.Ok,
			Data:           make([]byte, 38),
			AdditionalData: make([]byte, maxAllowed),
		})
		require.NoError(t, err)
	})
	t.Run("rejects one byte over", func(t *testing.T) {
		err := SignatureSubmitterDataVerifier(&SubProtocolResponse{
			Status:         payload.Ok,
			Data:           make([]byte, 38),
			AdditionalData: make([]byte, maxAllowed+1),
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "additional data too long")
	})
}

func TestSignatureSubmitterDataVerifier_NonOkStatusSkipsLengthCheck(t *testing.T) {
	// Empty status must NOT inspect Data — pins the early-return contract.
	// (Wrong-length Data with Empty must be a no-op.)
	err := SignatureSubmitterDataVerifier(&SubProtocolResponse{
		Status: payload.Empty,
		Data:   []byte{0x01},
	})
	require.NoError(t, err)
}

// --- fetchData: robustness to upstream sending garbage ----------------------
//
// Threat model: ftso-client / fdc-client run alongside system-client on a
// trusted host. They aren't malicious, but a parser bug there must not crash
// system-client. These tests pin "we react cleanly to unexpected response
// shapes". Each case is a category of malformed response.

func makeSubProtocol(t *testing.T, srv *httptest.Server) *SubProtocol {
	t.Helper()
	return &SubProtocol{ID: 100, APIUrl: srv.URL, Type: 0}
}

func fetchURL(t *testing.T, srv *httptest.Server) *url.URL {
	t.Helper()
	u, err := url.Parse(srv.URL)
	require.NoError(t, err)
	return u
}

func TestFetchData_StripsHexPrefixOnOkResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"status":"OK","data":"0x1234","additionalData":"0xabcd"}`))
	}))
	defer srv.Close()

	resp, err := makeSubProtocol(t, srv).fetchData(fetchURL(t, srv), time.Second)
	require.NoError(t, err)
	require.Equal(t, payload.Ok, resp.Status)
	require.Equal(t, []byte{0x12, 0x34}, resp.Data)
	require.Equal(t, []byte{0xab, 0xcd}, resp.AdditionalData)
}

func TestFetchData_NonOkStatusSkipsHexDecode(t *testing.T) {
	// Non-OK status must NOT attempt to hex-decode Data. fdc-client emits
	// EMPTY / RETRY (go-flare-common payload enum); ftso-scaling emits
	// NOT_AVAILABLE from its own TS enum, which isn't a go-flare-common
	// constant — fetchData only checks `!= Ok`, so any non-OK string takes
	// the skip-decode path. Pin that contract for all three.
	for _, st := range []string{"EMPTY", "RETRY", "NOT_AVAILABLE"} {
		t.Run(st, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write([]byte(`{"status":"` + st + `","data":"not-hex","additionalData":""}`))
			}))
			defer srv.Close()

			resp, err := makeSubProtocol(t, srv).fetchData(fetchURL(t, srv), time.Second)
			require.NoError(t, err)
			require.Equal(t, payload.ResponseStatus(st), resp.Status)
			require.Empty(t, resp.Data)
		})
	}
}

func TestFetchData_RejectsNonHexDataOnOk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"status":"OK","data":"not-hex","additionalData":""}`))
	}))
	defer srv.Close()

	_, err := makeSubProtocol(t, srv).fetchData(fetchURL(t, srv), time.Second)
	require.Error(t, err)
	require.Contains(t, err.Error(), "decoding protocol client response body")
}

func TestFetchData_RejectsNonHexAdditionalDataOnOk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"status":"OK","data":"00","additionalData":"not-hex"}`))
	}))
	defer srv.Close()

	_, err := makeSubProtocol(t, srv).fetchData(fetchURL(t, srv), time.Second)
	require.Error(t, err)
	require.Contains(t, err.Error(), "additional data")
}

func TestFetchData_RejectsNon200(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	_, err := makeSubProtocol(t, srv).fetchData(fetchURL(t, srv), time.Second)
	require.Error(t, err)
	require.Contains(t, err.Error(), "http status")
}

func TestFetchData_RejectsMalformedJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{this is not json`))
	}))
	defer srv.Close()

	_, err := makeSubProtocol(t, srv).fetchData(fetchURL(t, srv), time.Second)
	require.Error(t, err)
}

func TestFetchData_RejectsUnknownJSONFields(t *testing.T) {
	// DisallowUnknownFields() is set — extra fields are a hard failure.
	// Defends against silent schema drift on the upstream side.
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"status":"OK","data":"00","additionalData":"","extra":"surprise"}`))
	}))
	defer srv.Close()

	_, err := makeSubProtocol(t, srv).fetchData(fetchURL(t, srv), time.Second)
	require.Error(t, err)
}

func TestFetchData_RejectsEmptyBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// no body written
	}))
	defer srv.Close()

	_, err := makeSubProtocol(t, srv).fetchData(fetchURL(t, srv), time.Second)
	require.Error(t, err)
}

func TestFetchData_RespectsTimeout(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-r.Context().Done()
	}))
	defer srv.Close()

	start := time.Now()
	_, err := makeSubProtocol(t, srv).fetchData(fetchURL(t, srv), 100*time.Millisecond)
	elapsed := time.Since(start)
	require.Error(t, err)
	require.Less(t, elapsed, 2*time.Second, "should bail near the timeout deadline")
}

// strString avoids importing strconv just for sub-test names.
func strString(n int) string {
	if n == 0 {
		return "0"
	}
	digits := []byte{}
	negative := n < 0
	if negative {
		n = -n
	}
	for n > 0 {
		digits = append([]byte{byte('0' + n%10)}, digits...)
		n /= 10
	}
	if negative {
		digits = append([]byte{'-'}, digits...)
	}
	return string(digits)
}
