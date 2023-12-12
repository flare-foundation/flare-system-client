package shared

type metrics struct {
	MetricsBase
	// Add additional metrics here
}

func newMetrics(namespace string) *metrics {
	return &metrics{
		MetricsBase: *NewMetricsBase(namespace),
	}
}
