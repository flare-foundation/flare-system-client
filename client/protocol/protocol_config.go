package protocol

import (
	"crypto/ecdsa"

	"flare-tlc/client/config"
	globalConfig "flare-tlc/config"
	"flare-tlc/utils/chain"
	"flare-tlc/utils/contracts/submission"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type protocolCredentials struct {
	submitPrivateKey       *ecdsa.PrivateKey  // sign tx for submit1, submit2, submit3
	submitSignaturesTxOpts *bind.TransactOpts // submitSignatures
	signerPrivateKey       *ecdsa.PrivateKey  // sign data for submitSignatures
}

type protocolAddresses struct {
	SubmitContractAddress common.Address
	signingAddress        common.Address // address of signerPrivateKey
}

type contractSelectors struct {
	submit1          []byte
	submit2          []byte
	submit3          []byte
	submitSignatures []byte
}

func newProtocolCredentials(
	chainID int,
	cfg *config.CredentialsConfig,
) (*protocolCredentials, error) {
	signerPkString, err := globalConfig.ReadFileToString(cfg.SigningPolicyPrivateKeyFile)
	if err != nil {
		return nil, errors.Wrap(err, "error reading signer private key")
	}
	signerPk, err := chain.PrivateKeyFromHex(signerPkString)
	if err != nil {
		return nil, errors.Wrap(err, "error creating signer private key")
	}

	submitPkString, err := globalConfig.ReadFileToString(cfg.ProtocolManagerSubmitPrivateKeyFile)
	if err != nil {
		return nil, errors.Wrap(err, "error reading submit private key")
	}
	submitPk, err := chain.PrivateKeyFromHex(submitPkString)
	if err != nil {
		return nil, errors.Wrap(err, "error creating submit private key")
	}

	submitSignaturesTxOpts, _, err := chain.CredentialsFromPrivateKey(cfg.ProtocolManagerSubmitSignaturesPrivateKeyFile, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "error creating submit signatures tx opts")
	}

	return &protocolCredentials{
		submitPrivateKey:       submitPk,
		submitSignaturesTxOpts: submitSignaturesTxOpts,
		signerPrivateKey:       signerPk,
	}, nil
}

func newProtocolAddresses(
	pc *protocolCredentials,
	cfg *globalConfig.ContractAddresses,
) (*protocolAddresses, error) {
	signingAddress, err := chain.PrivateKeyToEthAddress(pc.signerPrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "error getting signing address")
	}
	return &protocolAddresses{
		SubmitContractAddress: cfg.Submission,
		signingAddress:        signingAddress,
	}, nil
}

func newContractSelectors() contractSelectors {
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
