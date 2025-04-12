package main

import (
	"fmt"
	"os"

	"github.com/polyscichart/polyscichart/chartdata"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: checkcsv <psc_csv_file>")
		os.Exit(1)
	}

	post, err := chartdata.ParseCSV(os.Args[1])
	if err != nil {
		fmt.Printf("Error parsing PSC-CSV: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Post Summary:")
	fmt.Printf("  X Title: %s\n", post.XTitle)
	fmt.Printf("  X Text: %s\n", post.XText)
	fmt.Printf("  X Sponsor: %s\n", post.XSponsor)
	fmt.Printf("  X Source: %v\n", post.XSource)
	fmt.Printf("  Chart Type: %s\n", post.ChartType)
	fmt.Printf("  Image Path: %s\n", post.ImagePath)
	fmt.Printf("  Title: %s\n", post.Title)
	fmt.Printf("  Chart Alt: %s\n", post.ChartAlt)
	fmt.Printf("  Style: %v\n", post.Style)
	fmt.Printf("  Chart Data:\n")
	fmt.Printf("    Labels: %v\n", post.ChartData.Labels)
	for i, row := range post.ChartData.Values {
		fmt.Printf("    Row %d: %v\n", i+1, row)
	}
}
