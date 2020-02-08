package golang

// ChartData represents the chart information.
type ChartData struct {
	Labels []string               `json:"labels"`
	Extra  map[string]interface{} `json:"extra"`
}

// DatasetData represents the dataset information.
type DatasetData struct {
	Name   string                 `json:"name"`
	Values []float64              `json:"values"`
	Extra  map[string]interface{} `json:"extra"`
}

// ServerData represents how the server is expected
// to send the data to the chartisan client.
type ServerData struct {
	Chart    ChartData     `json:"chart"`
	Datasets []DatasetData `json:"datasets"`
}
