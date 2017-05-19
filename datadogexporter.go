package eventexporter

import kodingmetrics "github.com/koding/metrics"

type DatadogExporter struct {
	datadog *kodingmetrics.DogStatsD
}

// NewDatadogExporter initializes DatadogExporter struct
// and NewDatadogExporter implements Exporter interface with Send and Close functions
func NewDatadogExporter(d *kodingmetrics.DogStatsD) *DatadogExporter {
	return &DatadogExporter{datadog: d}
}

// Send publishes Events to Datad
func (d *DatadogExporter) Send(m *Event) error {
	eventName, tags := eventSeperator(m)

	if m.Duration > 0 {
		_ = d.datadog.Timing(eventName, m.Duration, tags, 1)
		return nil
	}

	count := m.Count
	if count == 0 {
		count = 1
	}
	_ = d.datadog.Count(eventName, count, tags, 1)
	return nil
}

func (d *DatadogExporter) Close() error {
	return nil
}
