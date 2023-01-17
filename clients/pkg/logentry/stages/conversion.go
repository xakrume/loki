package main

type conversionStage struct {
	target       string
	conversionMap map[string]float64
}

func (c *conversionStage) Process(labels api.Labels, entry *api.Entry) (*api.Labels, *api.Entry) {
	if val, ok := labels[c.target]; ok {
		if unit, ok := labels["latency_unit"]; ok {
			if conversionFactor, ok := c.conversionMap[unit]; ok {
				labels["latency_seconds"] = val * conversionFactor
			}
		}
	}

	return &labels, entry
}

func (c *conversionStage) Name() string {
	return "conversion"
}

func init() {
	pipeline.Register(func() pipeline.Stage {
		return &conversionStage{
			target: "latency_value",
			conversionMap: map[string]float64{
				"ns": 1e-9,
				"us": 1e-6,
				"ms": 1e-3,
				"s": 1,
				"m": 60,
				"h": 3600,
			},
		}
	})
}
