package finalizer

// Transform signature to be used by go-ethereum crypto.SigToPub:
// transforms [V || R || S] to [R || S || V - 27]
// No checks are performed, we assume that signature array has length 65
func transformSignature(signature []byte) (RSV [65]byte) {
	copy(RSV[:], signature[1:33])
	copy(RSV[32:], signature[33:65])
	RSV[64] = signature[0] - 27
	return RSV
}
