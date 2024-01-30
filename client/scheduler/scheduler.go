package scheduler

import (
	"flare-tlc/logger"
	"flare-tlc/utils"
	"time"
)

type ScheduledJob interface {
	// Run runs the job for the given epoch
	// Returns false if the job was skipped (e.g. because it was already executed) and true otherwise
	Run(epoch int64) (bool, error)
}

type EpochScheduler struct {
	Epoch       *utils.Epoch
	StartOffset time.Duration

	Job ScheduledJob

	// We use a time provider to make testing easier (can mock time)
	TimeProvider utils.TimeProvider
}

type RetryableScheduler struct {
	EpochScheduler

	// RetryDelay is the time to wait before retrying a failed try
	RetryDelay time.Duration
}

func NewEpochScheduler(epoch *utils.Epoch, startOffset time.Duration, job ScheduledJob) *EpochScheduler {
	return &EpochScheduler{
		Epoch:        epoch,
		StartOffset:  startOffset,
		Job:          job,
		TimeProvider: utils.RealTimeProvider{},
	}
}

func (s *EpochScheduler) Run() {
	et := utils.NewEpochTicker(s.StartOffset, s.Epoch)
	for {
		epoch := <-et.C
		_, err := s.Job.Run(epoch)
		if err != nil {
			// Do not fail or try to repeat the job, just log the error
			logger.Error("error running scheduled job: %v", err)
		}
	}
}

func NewRetryableScheduler(epoch *utils.Epoch, startOffset time.Duration, retryDelay time.Duration, job ScheduledJob) *RetryableScheduler {
	return &RetryableScheduler{
		EpochScheduler: *NewEpochScheduler(epoch, startOffset, job),
		RetryDelay:     retryDelay,
	}
}

func (s *RetryableScheduler) Run() {
	ret := utils.NewRetriableEpochTicker(s.StartOffset, s.RetryDelay, s.Epoch)
	for {
		epoch := <-ret.C
		didRun, err := s.Job.Run(epoch)
		if err != nil {
			logger.Error("error running scheduled job %w", err)
		} else if didRun {
			logger.Info("scheduled job completed successfully for epoch %d", epoch)
		}
	}
}
