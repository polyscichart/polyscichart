package chartdata

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

type ChartData struct {
	Labels []string
	Values [][]float64
}
