package protocol

type EpochRunner interface {
	RunEpoch(currentEpoch int64)
}
