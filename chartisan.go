package golang

import "encoding/json"

// Chartisan is the main chartisan struct.
type Chartisan struct {
	serverData ServerData
}

// Build creates a new instance of a chartisan chart.
func Build() *Chartisan {
	return &Chartisan{
		serverData: ServerData{
			Chart:    ChartData{Extra: nil},
			Datasets: []DatasetData{},
		},
	}
}

// Labels sets the chart labels.
func (chart *Chartisan) Labels(labels []string) *Chartisan {
	chart.serverData.Chart.Labels = labels
	return chart
}

// Extra adds extra information to the chart.
func (chart *Chartisan) Extra(value map[string]interface{}) *Chartisan {
	chart.serverData.Chart.Extra = value
	return chart
}

// AdvancedDataset appends a new dataset to the chart or modifies an existing one.
// If the ID has already been used, the dataset will be replaced with this one.
func (chart *Chartisan) AdvancedDataset(
	name string,
	values []int,
	extra map[string]interface{},
) *Chartisan {
	// Get or create the given dataset.
	dataset, isNew := chart.getOrCreateDataset(name, values, extra)
	if isNew {
		// Append the new dataset.
		chart.serverData.Datasets = append(chart.serverData.Datasets, *dataset)
		return chart
	}
	// Modify the existing dataset.
	dataset.Name = name
	dataset.Values = values
	dataset.Extra = extra
	return chart
}

// Dataset adds a new simple dataset to the chart. If more advanced control is
// needed, consider using `AdvancedDataset` instead.
func (chart *Chartisan) Dataset(name string, values []int) *Chartisan {
	chart.AdvancedDataset(name, values, nil)
	return chart
}

// ToJSON transforms the chart into the JSON representation needed.
func (chart *Chartisan) ToJSON() string {
	json, err := json.Marshal(chart.serverData)
	if err != nil {
		return `{"error": "Error converting chart to JSON"}`
	}
	return string(json)
}

// ToObject transforms the chart into the ServerData needed.
func (chart *Chartisan) ToObject() ServerData {
	return chart.serverData
}

// getOrCreateDataset returns a dataset from the chart or creates a new one given the data.
func (chart *Chartisan) getOrCreateDataset(
	name string,
	values []int,
	extra map[string]interface{},
) (*DatasetData, bool) {
	for i := 0; i < len(chart.serverData.Datasets); i++ {
		if chart.serverData.Datasets[i].Name == name {
			return &chart.serverData.Datasets[i], false
		}
	}
	return &DatasetData{
		Name:   name,
		Values: values,
		Extra:  extra,
	}, true
}
