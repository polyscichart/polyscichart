package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Post struct {
	XTitle    string
	XText     string
	XSponsor  string
	XSource   []string
	ChartType string
	ImagePath string
	Title     string
	ChartAlt  string
	Style     map[string]string
	ChartData ChartData
}

const PSCHeaderMin = 3 // Minimum number of rows for a valid post

type ChartData struct {
	Labels []string
	Values [][]float64
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: checkcsv <psc_csv_file>")
		os.Exit(1)
	}

	post, err := parseCSV(os.Args[1])
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

func parseCSV(filePath string) (Post, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Post{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()
	if err != nil {
		return Post{}, err
	}

	if len(records) < PSCHeaderMin { // Minimum: header + delimiter + data
		return Post{}, fmt.Errorf("PSC-CSV too short: %d rows, minimum %d", len(records), PSCHeaderMin)
	}

	// Skip optional header row if present
	startIdx := 0
	if records[0][0] == "Key" && records[0][1] == "Value1" {
		startIdx = 1
	}

	// Find delimiters
	postEndIdx := -1
	chartStartIdx := -1
	for i, row := range records[startIdx:] {
		if len(row) > 0 {
			if row[0] == "+++" {
				postEndIdx = i + startIdx
			} else if row[0] == "---" {
				chartStartIdx = i + startIdx
				break
			}
		}
	}
	if postEndIdx == -1 || chartStartIdx == -1 || postEndIdx >= chartStartIdx {
		return Post{}, fmt.Errorf("invalid delimiters: +++ o	r --- missing or misplaced")
	}

	// Parse header section
	post := Post{Style: make(map[string]string)}
	for _, row := range records[startIdx:postEndIdx] {
		if len(row) < 2 {
			continue
		}
		key := row[0]
		switch key {
		case "x_title":
			post.XTitle = row[1]
		case "x_text":
			post.XText = strings.ReplaceAll(row[1], "\n", " ") // Strip newlines
		case "x_sponsor":
			post.XSponsor = row[1]
		case "x_source":
			for _, v := range row[1:] {
				if v != "" {
					post.XSource = append(post.XSource, v)
				}
			}
		case "chart_type":
			post.ChartType = row[1]
		case "image_path":
			post.ImagePath = row[1]
		case "title":
			post.Title = row[1]
		case "chart_alt":
			post.ChartAlt = row[1]
		case "style":
			for _, v := range row[1:] {
				if v != "" {
					kv := strings.SplitN(v, "=", 2)
					if len(kv) == 2 {
						post.Style[kv[0]] = kv[1]
					}
				}
			}
		}
	}

	// Parse chart data
	if post.ImagePath == "" && chartStartIdx+1 < len(records) {
		post.ChartData.Labels = records[chartStartIdx+1]
		for _, row := range records[chartStartIdx+2:] {
			values := make([]float64, len(row))
			for i, v := range row {
				fmt.Sscanf(strings.TrimSpace(v), "%f", &values[i])
			}
			post.ChartData.Values = append(post.ChartData.Values, values)
		}
	}

	return post, nil
}
