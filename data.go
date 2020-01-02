package chartisan

// ChartData represents the chart information.
type ChartData struct {
	Labels []string    `json:"labels"`
	Extra  interface{} `json:"extra"`
}

// DatasetData represents the dataset information.
type DatasetData struct {
	ID     int         `json:"id"`
	Name   string      `json:"name"`
	Values []int       `json:"values"`
	Extra  interface{} `json:"extra"`
}

// ServerData represents how the server is expected
// to send the data to the chartisan client.
type ServerData struct {
	Chart    ChartData     `json:"chart"`
	Datasets []DatasetData `json:"datasets"`
}
