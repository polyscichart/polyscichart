// datastore/chartdata.go

package datastore

// ChartData is an interface for all chart data types.
type ChartData interface {
	Type() string
}

// PieChartData represents the data for a pie chart.
type PieChartData struct {
	Labels []string  `json:"labels"`
	Values []float64 `json:"values"`
}

// Type returns the chart type.
func (p PieChartData) Type() string {
	return "pie"
}

// BarChartData represents the data for a bar chart.
type BarChartData struct {
	Labels   []string  `json:"labels"`
	Values   []float64 `json:"values"`
	Category []string  `json:"category,omitempty"` // Optional category for grouped bar charts
}

// Type returns the chart type.
func (b BarChartData) Type() string {
	return "bar"
}

// LineChartData represents the data for a line chart.
type LineChartData struct {
	Labels []string  `json:"labels"`
	Values []float64 `json:"values"`
}

// Type returns the chart type.
func (l LineChartData) Type() string {
	return "line"
}

// ParetoChartData represents the data for a Pareto chart.
type ParetoChartData struct {
	Labels []string  `json:"labels"`
	Values []float64 `json:"values"`
}

// Type returns the chart type.
func (p ParetoChartData) Type() string {
	return "pareto"
}

// ScatterChartData represents the data for a scatter chart.
type ScatterChartData struct {
	XValues []float64 `json:"x_values"`
	YValues []float64 `json:"y_values"`
}

// Type returns the chart type.
func (s ScatterChartData) Type() string {
	return "scatter"
}

// BubbleChartData represents the data for a bubble chart.
type BubbleChartData struct {
	XValues []float64 `json:"x_values"`
	YValues []float64 `json:"y_values"`
	Sizes   []float64 `json:"sizes"`
	Labels  []string  `json:"labels"`
}

// Type returns the chart type
func (b BubbleChartData) Type() string {
	return "bubble"
}