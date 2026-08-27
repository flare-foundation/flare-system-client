package chain

import "github.com/ethereum/go-ethereum/crypto"

// relayErrorNames are the Relay's custom errors, copied verbatim from
// IRelay.sol in flare-smart-contracts-v2. All are zero-argument, so the name is
// the signature. The new Relay's relay() is hand-written assembly that reverts
// with a bare 4-byte selector, so a revert carries no Error(string) to unpack.
var relayErrorNames = []string{
	"AlreadyRelayed()",
	"BadS()",
	"BadV()",
	"DelayedSignPolicy()",
	"DuplicateProtocolId()",
	"EcrecoverError()",
	"EcrecoverReturnedBadData()",
	"FeeCollectionAddressZero()",
	"FeeConfigNotAllowed()",
	"FeeExemptAddressZero()",
	"FeeExemptionsNotAllowed()",
	"FeeTokenActive()",
	"FeeTransferFailed()",
	"HistoryBeforeStart()",
	"IncorrectMerkleProof()",
	"IndexOutOfOrder()",
	"IndexOutOfRange()",
	"InitialSigningPolicyHashZero()",
	"InvalidConfigHash()",
	"InvalidInitialStartingVotingRoundId()",
	"InvalidProtocolId()",
	"InvalidRandomNumberProof()",
	"InvalidRandomNumberProtocolId()",
	"InvalidSignPolicyLength()",
	"InvalidSignPolicyMetadata()",
	"InvalidVotingRoundId()",
	"MerkleProofInvalid()",
	"MessageTooOld()",
	"MsgValueNotAllowed()",
	"MustUseNewSignPolicy()",
	"NoAccessToMerkleRoots()",
	"NoAccessToSigningPolicyHashes()",
	"NoNewSignPolicySize()",
	"NoRandomNumber()",
	"NoSignatureCount()",
	"NotEnoughSignatures()",
	"NotEnoughWeight()",
	"NotFinalized()",
	"NotNextRewardEpoch()",
	"NotWithLastInitialized()",
	"OldRelayIncompatible()",
	"OldRelayNotAllowedInRelayMode()",
	"OldRelayVerificationFailed()",
	"OldRelayWrongFirstRewardEpochStart()",
	"OldRelayWrongRewardEpochDuration()",
	"OldRelayWrongStartTs()",
	"OldRelayWrongVotingEpochDuration()",
	"OnlySigningPolicySetterRole()",
	"ProtocolFeeZero()",
	"RefundFailed()",
	"RewardEpochDurationZero()",
	"SignPolicyRelayDisabled()",
	"SigningPolicyEmpty()",
	"SigningPolicyHashMismatch()",
	"SigningPolicySetterNotAllowed()",
	"SigningPolicySetterZero()",
	"SourceChainIdMismatchOnHomeDeploy()",
	"SourceChainIdZero()",
	"ThresholdIncreaseTooSmall()",
	"ThresholdTooHigh()",
	"ThresholdTooLow()",
	"TooLowFee()",
	"TooManyVoters()",
	"TooShortMessage()",
	"TotalWeightTooBig()",
	"UnreachableCode()",
	"VerificationFailed()",
	"VotersWeightsSizeMismatch()",
	"VotingEpochDurationZero()",
	"WrongMessageFormat()",
	"WrongMessageFormat2()",
	"WrongSignPolicyRewardEpoch()",
	"WrongSignature()",
	"WrongSizeForNewSignPolicy()",
	"WrongVerificationData()",
	"ZeroMerkleRoot()",
	"ZeroSigner()",
}

// relayErrors maps a custom-error selector to its signature. Derived from the
// names, so it cannot drift from what the contract folds into its assembly.
var relayErrors = func() map[[4]byte]string {
	m := make(map[[4]byte]string, len(relayErrorNames))
	for _, name := range relayErrorNames {
		m[[4]byte(crypto.Keccak256([]byte(name))[:4])] = name
	}
	return m
}()

// relayErrorName resolves revert data that is a bare custom-error selector.
func relayErrorName(data []byte) (string, bool) {
	if len(data) < 4 {
		return "", false
	}
	name, ok := relayErrors[[4]byte(data[:4])]
	return name, ok
}
