package cronjob

import (
	"flare-fsc/client/shared"
	"flare-fsc/utils"
	"time"

	"gitlab.com/flarenetwork/libs/go-flare-common/pkg/logger"
)

type Cronjob interface {
	Name() string
	Enabled() bool
	Timeout() time.Duration
	RandomTimeoutDelta() time.Duration
	Call() error
	OnStart() error

	// Set health status of cronjob.
	// (can be implemented to ignore the status based on other conditions)
	UpdateCronjobStatus(status shared.HealthStatus)
}

func RunCronjob(c Cronjob) {
	if !c.Enabled() {
		logger.Debugf("%s cronjob disabled", c.Name())
		c.UpdateCronjobStatus(shared.HealthStatusOk)
		return
	}

	err := c.OnStart()
	if err != nil {
		logger.Errorf("%s cronjob on start error %v", c.Name(), err)
		return
	}

	logger.Debugf("starting %s cronjob", c.Name())

	ticker := utils.NewRandomizedTicker(c.Timeout(), c.RandomTimeoutDelta())
	for {
		<-ticker

		err := c.Call()
		if err == nil {
			c.UpdateCronjobStatus(shared.HealthStatusOk)
		} else {
			logger.Errorf("%s cronjob error %s", c.Name(), err.Error())
			c.UpdateCronjobStatus(shared.HealthStatusError)
		}
	}
}
