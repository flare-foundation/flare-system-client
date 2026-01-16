package protocol

import (
	"crypto/ecdsa"

	"github.com/flare-foundation/flare-system-client/client/config"
	globalConfig "github.com/flare-foundation/flare-system-client/config"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/submission"
)

// Private keys and addresses needed for protocol voter
type protocolContext struct {
	submitPrivateKey           *ecdsa.PrivateKey // sign tx for submit1, submit2, submit3
	submitSignaturesPrivateKey *ecdsa.PrivateKey // submitSignatures
	signerPrivateKey           *ecdsa.PrivateKey // sign data for submitSignatures

	submitContractAddress   common.Address
	signingAddress          common.Address // address of signerPrivateKey
	submitAddress           common.Address // address of submitPrivateKey
	submitSignaturesAddress common.Address // address of submitSignaturesPrivateKey
}

type contractSelectors struct {
	submit1          []byte
	submit2          []byte
	submit3          []byte
	submitSignatures []byte
}

func newProtocolContext(cfg *config.Client) (*protocolContext, error) {
	ctx := &protocolContext{}

	var err error

	// Credentials
	ctx.signerPrivateKey, err = globalConfig.PrivateKeyFromConfig(cfg.Credentials.SigningPolicyPrivateKeyFile,
		cfg.Credentials.SigningPolicyPrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "error creating signer private key")
	}

	ctx.submitPrivateKey, err = globalConfig.PrivateKeyFromConfig(cfg.Credentials.ProtocolManagerSubmitPrivateKeyFile,
		cfg.Credentials.ProtocolManagerSubmitPrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "error creating submit private key")
	}

	ctx.submitSignaturesPrivateKey, err = globalConfig.PrivateKeyFromConfig(cfg.Credentials.ProtocolManagerSubmitSignaturesPrivateKeyFile,
		cfg.Credentials.ProtocolManagerSubmitSignaturesPrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "error reading submit signatures private key")
	}

	// Addresses
	ctx.signingAddress = crypto.PubkeyToAddress(ctx.signerPrivateKey.PublicKey)
	ctx.submitAddress = crypto.PubkeyToAddress(ctx.submitPrivateKey.PublicKey)
	ctx.submitSignaturesAddress = crypto.PubkeyToAddress(ctx.submitSignaturesPrivateKey.PublicKey)

	ctx.submitContractAddress = cfg.ContractAddresses.Submission

	return ctx, nil
}

// ContractSelectors return function selectors for submission contract.
func ContractSelectors() contractSelectors {
	submissionABI, err := submission.SubmissionMetaData.GetAbi()
	if err != nil {
		// panic, this error is fatal
		panic(err)
	}
	return contractSelectors{
		submit1:          submissionABI.Methods["submit1"].ID,
		submit2:          submissionABI.Methods["submit2"].ID,
		submit3:          submissionABI.Methods["submit3"].ID,
		submitSignatures: submissionABI.Methods["submitSignatures"].ID,
	}
}
