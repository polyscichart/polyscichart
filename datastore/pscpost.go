// datastore/datastore.go
package datastore

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// PSCPost represents the data structure for a PolySciChart post.
type PSCPost struct {
	ChartPostID string    `json:"chart_post_id"`         // Unique identifier: YYYYMMDD-###
	XPostURL    string    `json:"x_post_url"`            // URL of the X post
	ChartData   ChartData `json:"chart_data"`            // JSON data for chart generation
	CreatedAt   time.Time `json:"created_at"`            // Timestamp of creation (ISO 8601)
	Sponsor     string    `json:"sponsor,omitempty"`     // Optional sponsor name
	Description string    `json:"description,omitempty"` // Optional description
	SourceURLs  []string  `json:"source_urls,omitempty"` // Optional array of source URLs
	Tags        []string  `json:"tags,omitempty"`        // Optional array of tags
}

// UnmarshalJSON unmarshals the ChartData field from JSON.
func (p *PSCPost) UnmarshalJSON(data []byte) error {
	// Define a temporary struct to hold the unmarshaled data.
	type Alias PSCPost
	aux := &struct {
		ChartData json.RawMessage `json:"chart_data"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}

	// Unmarshal the data into the temporary struct.
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Determine the chart type from the ChartData field.
	var chartType struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(aux.ChartData, &chartType); err != nil {
		return err
	}

	// Create a new ChartData object based on the chart type.
	switch chartType.Type {
	case "pie":
		var chart PieChartData
		if err := json.Unmarshal(aux.ChartData, &chart); err != nil {
			return err
		}
		p.ChartData = chart
	case "bar":
		var chart BarChartData
		if err := json.Unmarshal(aux.ChartData, &chart); err != nil {
			return err
		}
		p.ChartData = chart
	case "line":
		var chart LineChartData
		if err := json.Unmarshal(aux.ChartData, &chart); err != nil {
			return err
		}
		p.ChartData = chart
	case "pareto":
		var chart ParetoChartData
		if err := json.Unmarshal(aux.ChartData, &chart); err != nil {
			return err
		}
		p.ChartData = chart
	case "scatter":
		var chart ScatterChartData
		if err := json.Unmarshal(aux.ChartData, &chart); err != nil {
			return err
		}
		p.ChartData = chart
	case "bubble":
		var chart BubbleChartData
		if err := json.Unmarshal(aux.ChartData, &chart); err != nil {
			return err
		}
		p.ChartData = chart
	default:
		// Return an error if the chart type is not supported.
		return fmt.Errorf("unsupported chart type: %s", chartType.Type)
	}

	return nil
}

// MarshalJSON marshals the ChartData field to JSON.
func (p PSCPost) MarshalJSON() ([]byte, error) {
	type Alias PSCPost
	aux := &struct {
		ChartData interface{} `json:"chart_data"`
		*Alias
	}{
		Alias:     (*Alias)(&p),
		ChartData: p.ChartData,
	}
	return json.Marshal(aux)
}

func FromFile(filePath string) (*PSCPost, error) {
	// Read the file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	// Unmarshal the JSON data into a PSCPost object
	var p PSCPost
	err = json.Unmarshal(data, &p)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}
	return &p, nil
}
