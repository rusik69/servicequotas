package types

// QuotasConfig represents the quotas configuration
type QuotasConfig struct {
	Quotas []struct {
		ServiceCode string  `json:"ServiceCode"`
		ServiceName string  `json:"ServiceName"`
		QuotaArn    string  `json:"QuotaArn"`
		QuotaCode   string  `json:"QuotaCode"`
		QuotaName   string  `json:"QuotaName"`
		Value       float64 `json:"Value"`
		Unit        string  `json:"Unit"`
		Adjustable  bool    `json:"Adjustable"`
		GlobalQuota bool    `json:"GlobalQuota"`
		UsageMetric struct {
			MetricNamespace  string `json:"MetricNamespace"`
			MetricName       string `json:"MetricName"`
			MetricDimensions struct {
				Class    string `json:"Class"`
				Resource string `json:"Resource"`
				Service  string `json:"Service"`
				Type     string `json:"Type"`
			} `json:"MetricDimensions"`
			MetricStatisticRecommendation string `json:"MetricStatisticRecommendation"`
		} `json:"UsageMetric,omitempty"`
	} `json:"Quotas"`
}