package scheduler

import (
	"flare-tlc/logger"
	"flare-tlc/utils"
	"time"
)

type SchedulerBase struct {
	Job ScheduledJob

	// We use a time provider to make testing easier (can mock time)
	TimeProvider utils.TimeProvider
}

type EpochScheduler struct {
	SchedulerBase

	Config EpochSchedulerConfig
}

type RetryableScheduler struct {
	SchedulerBase

	Config RetryableSchedulerConfig
}

type EpochSchedulerConfig struct {
	Epoch *utils.Epoch

	// SendOffset is the time *before* the end of the epoch to send the transaction
	SendOffset time.Duration
}

type RetryableSchedulerConfig struct {
	Epoch *utils.Epoch

	// SendOffset is the time after the start of the epoch to send the transaction
	SendOffset time.Duration

	// RetryDelay is the time to wait before retrying a failed try
	RetryDelay time.Duration
}

type ScheduledJob interface {
	// Run runs the job for the given epoch
	// Returns false if the job was skipped (e.g. because it was already executed) and true otherwise
	Run(epoch int64) (bool, error)
}

func NewEpochScheduler(config EpochSchedulerConfig, job ScheduledJob) *EpochScheduler {
	return &EpochScheduler{
		SchedulerBase: SchedulerBase{
			Job:          job,
			TimeProvider: utils.RealTimeProvider{},
		},
		Config: config,
	}
}

func (s *EpochScheduler) nextScheduledTimeInEpoch(epoch int64) <-chan time.Time {
	now := s.TimeProvider.Now()
	sendTime := s.Config.Epoch.EndTime(epoch).Add(-s.Config.SendOffset)
	sendDuration := sendTime.Sub(now)

	if sendDuration >= 0 {
		return time.NewTimer(sendDuration).C
	} else {
		return time.NewTimer(sendDuration + s.Config.Epoch.Period).C
	}
}

func (s *EpochScheduler) Run() {
	et := utils.NewEpochTicker(s.Config.Epoch.Period-s.Config.SendOffset, s.Config.Epoch)
	for {
		epoch := <-et.C
		_, err := s.Job.Run(epoch)
		if err != nil {
			// Do not fail or try to repeat the job, just log the error
			logger.Error("error running scheduled job: %v", err)
		}
	}
}

func NewRetryableScheduler(config RetryableSchedulerConfig, job ScheduledJob) *RetryableScheduler {
	return &RetryableScheduler{
		SchedulerBase: SchedulerBase{
			Job:          job,
			TimeProvider: utils.RealTimeProvider{},
		},
		Config: config,
	}
}

func (s *RetryableScheduler) Run() {
	ret := utils.NewRetriableEpochTicker(s.Config.SendOffset, s.Config.RetryDelay, s.Config.Epoch)
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
