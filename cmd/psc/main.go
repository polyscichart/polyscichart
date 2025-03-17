package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"

	"github.com/polyscichart/polyscichart/datastore" // Replace with your actual import path
)

var command2Args = map[string]int{
	"chart": 3,  // create an image from the indicated pscpost
	"post":  3,  // dump the pscpost text data as would be sent to X to stdout
	"xpost": 3,  // post to X
}

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Usage: psc <command> <pscpost>")
		return
	}
	// now pull the command from cmdline
	cmd := os.Args[1]
	if checkArgs(cmd) == false {
		return
	}

	switch cmd {
	case "chart":
		pscPost, err := datastore.FromFile(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading datastore file: %v\n", err)
			return
		}
		generatePSCPostJSON(pscPost)
	}

}

func checkArgs(cmd string) bool {

	if requiredArgs, exists := command2Args[cmd]; exists {
		if len(os.Args)-2 >= requiredArgs {
			return true
		}
		fmt.Fprintf(os.Stderr, "Error: Command '%s' requires at least %d arguments, but %d were provided.\n", cmd, requiredArgs, len(os.Args)-2)
	} else {
		fmt.Fprintf(os.Stderr, "Error: Unknown command '%s'.\n", cmd)
	}
	return false
}

func generatePSCPostJSON(pscPost datastore.PSCPost) error {
	// Marshal the PSCPost data to JSON
	jsonData, err := json.MarshalIndent(pscPost, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	// Write the JSON data to a file
	err = os.WriteFile("pscpost.json", jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing JSON to file: %w", err)
	}

	fmt.Println("pscpost.json created successfully.")
	return nil
}

func generateBarChartPNGFromPSCPost(data datastore.BarChartData) {
	pscPost, ok := data.(datastore.PSCPost)
	if !ok {
		fmt.Println("Error: Invalid data type. Expected PSCPost")
		return
	}
	chartData, ok := pscPost.ChartData.(datastore.ChartBar
	Or
	PIe )

	graph := chart.BarChart{
		Title: "GDP by Country (Trillions USD)",
		TitleStyle: chart.Style{
			FontSize:  16,
			FontColor: drawing.ColorBlack,
		},
		XAxis: chart.Style{
			Show: true,
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		Bars: []chart.Value{},
	}

	for i := 0; i < len(data.Labels); i++ {
		graph.Bars = append(graph.Bars, chart.Value{
			Value: float64(data.Values[i]),
			Label: data.Labels[i],
		})
	}

	graph.Background = chart.Style{
		Padding: chart.Box{
			Top: 20,
		},
	}

	// Write to a file
	f, err := os.Create("bar_chart.png")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	err = graph.Render(chart.PNG, f)
	if err != nil {
		fmt.Println("Error rendering chart:", err)
		return
	}

	fmt.Println("bar_chart.png created successfully.")
}

func generateDataAndChart() {
	// Sample bar chart data
	barChartData := datastore.BarChartData{
		Labels: []string{"US", "China", "Ukraine", "EU", "Russia"},
		Values: []float64{23.0, 17.7, 0.2, 17.1, 1.7}, // GDP in Trillions (made up numbers
	}

	// Create a PSCPost with the bar chart data
	pscPost := datastore.PSCPost{
		ChartPostID: "20240701-001",
		XPostURL:    "https://x.com/example", // Placeholder
		ChartData:   barChartData,
		CreatedAt:   time.Now(),
		Description: "GDP by Country (Trillions USD)",
		Tags:        []string{"#GDP", "#Economy", "#Data"},
	}

	err := generatePSCPostJSON(pscPost)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Generate a PNG chart using go-chart
	generateBarChartPNGFromPSCPost(barChartData)
}
func generateBarChartPNG(data datastore.BarChartData) {
	graph := chart.BarChart{
		Title: "GDP by Country (Trillions USD)",
		TitleStyle: chart.Style{
			FontSize:  16,
			FontColor: drawing.ColorBlack,
		},
		XAxis: chart.Style{
			Show: true,
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		Bars: []chart.Value{},
	}

	for i := 0; i < len(data.Labels); i++ {
		graph.Bars = append(graph.Bars, chart.Value{
			Value: float64(data.Values[i]),
			Label: data.Labels[i],
		})
	}

	graph.Background = chart.Style{
		Padding: chart.Box{
			Top: 20,
		},
	}

	// Write to a file
	f, err := os.Create("bar_chart.png")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	err = graph.Render(chart.PNG, f)
	if err != nil {
		fmt.Println("Error rendering chart:", err)
		return
	}

	fmt.Println("bar_chart.png created successfully.")
}
