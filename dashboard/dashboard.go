/*
Copyright © 2021 Peter Parada <parad.peter@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package dashboard

type GrafanaDashboardParams struct {
	Title string
	Style string
}

type GrafanaDashboardConfigTime struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

type GrafanaDashboardConfigTimepicker struct {
	RefreshIntervals []string `json:"refresh_intervals"`
	TimeOptions      []string `json:"time_options"`
}

type GrafanaDashboardConfigTemplatingListItemCurrent struct {
	Text  string `json:"text,omitempty"`
	Value string `json:"value,omitempty"`
}

type GrafanaDashboardConfigTemplatingListItemOptions struct {
	Selected bool   `json:"selected,omitempty"`
	Text     string `json:"text,omitempty"`
	Value    string `json:"value,omitempty"`
}

type GrafanaDashboardConfigTemplatingListItem struct {
	Current GrafanaDashboardConfigTemplatingListItemCurrent   `json:"current"`
	Name    string                                            `json:"name"`
	Query   string                                            `json:"query"`
	Type    string                                            `json:"type"`
	Options []GrafanaDashboardConfigTemplatingListItemOptions `json:"options"`
}

type GrafanaDashboardConfigTemplating struct {
	List []GrafanaDashboardConfigTemplatingListItem `json:"list,omitempty"`
}

type GrafanaDashboardConfigPanelGridPos struct {
	H int `json:"h,omitempty"`
	W int `json:"w,omitempty"`
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
}

type GrafanaDashboardConfigPanelGauge struct {
	MaxValue         int  `json:"maxValue,omitempty"`
	MinValue         int  `json:"minValue,omitempty"`
	Show             bool `json:"show,omitempty"`
	ThresholdLabels  bool `json:"thresholdLabels,omitempty"`
	ThresholdMarkers bool `json:"thresholdMarkers,omitempty"`
}

type GrafanaDashboardConfigPanelTarget struct {
	Expr           string `json:"expr,omitempty"`
	Format         string `json:"format,omitempty"`
	IntervalFactor int    `json:"intervalFactor,omitempty"`
	RefId          string `json:"refId,omitempty"`
}

type GrafanaDashboardConfigPanelSparkline struct {
	FillColor string `json:"fillColor,omitempty"`
	Full      bool   `json:"full"`
	LineColor string `json:"lineColor,omitempty"`
	Show      bool   `json:"show"`
}

type GrafanaDashboardConfigPanelRangeMap struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
	Text string `json:"text,omitempty"`
}

type GrafanaDashboardConfigPanelOptionsReduceOptions struct {
	Calcs  []string `json:"calcs,omitempty"`
	Fields string   `json:"fields,omitempty"`
	Values bool     `json:"values"`
}

type GrafanaDashboardConfigPanelOptions struct {
	Orientation          string                                          `json:"orientation,omitempty"`
	ReduceOptions        GrafanaDashboardConfigPanelOptionsReduceOptions `json:"reduceOptions,omitempty"`
	ShowThresholdLabels  bool                                            `json:"showThresholdLabels"`
	ShowThresholdMarkers bool                                            `json:"showThresholdMarkers"`
}

type GrafanaDashboardConfigPanel struct {
	Id            int                                 `json:"id,omitempty"`
	GridPos       GrafanaDashboardConfigPanelGridPos  `json:"gridPos,omitempty"`
	Title         string                              `json:"title,omitempty"`
	Type          string                              `json:"type,omitempty"`
	MaxDataPoints int                                 `json:"maxDataPoints,omitempty"`
	Targets       []GrafanaDashboardConfigPanelTarget `json:"targets,omitempty"`
	Options       GrafanaDashboardConfigPanelOptions  `json:"options,omitempty"`
}

type GrafanaDashboardConfigInput struct {
	Name        string `json:"name,omitempty"`
	Label       string `json:"label,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	PluginId    string `json:"pluginId,omitempty"`
	PluginName  string `json:"pluginName,omitempty"`
}

type GrafanaDashboardConfig struct {
	Id            string                           `json:"id,omitempty"`
	Uid           string                           `json:"uid,omitempty"`
	Title         string                           `json:"title"`
	Style         string                           `json:"style"`
	Timezone      string                           `json:"timezone"`
	Editable      bool                             `json:"editable"`
	GraphTooltip  int                              `json:"graphTooltip"`
	Time          GrafanaDashboardConfigTime       `json:"time"`
	Timepicker    GrafanaDashboardConfigTimepicker `json:"timepicker"`
	Templating    GrafanaDashboardConfigTemplating `json:"templating"`
	Refresh       string                           `json:"refresh"`
	SchemaVersion int                              `json:"schemaVersion"`
	Version       int                              `json:"version"`
	Panels        []GrafanaDashboardConfigPanel    `json:"panels"`
	Annotations   []string                         `json:"annotations"`
	Inputs        []GrafanaDashboardConfigInput    `json:"__inputs,omitempty"`
}

func GenerateDashboard(params *GrafanaDashboardParams) interface{} {
	greenColor := "#299c46"
	redColor := "#d44a3a"
	orangeColor := "rgba(237, 129, 40, 0.89)"
	lightBlueColor := "rgba(31, 118, 189, 0.18)"
	whiteColor := "rgb(31, 120, 193)"

	return &GrafanaDashboardConfig{
		Title:        params.Title,
		Style:        params.Style,
		Timezone:     "browser",
		Editable:     true,
		GraphTooltip: 1,
		Time: GrafanaDashboardConfigTime{
			From: "now-6h",
			To:   "now",
		},
		Timepicker: GrafanaDashboardConfigTimepicker{
			RefreshIntervals: []string{
				"5s",
				"10s",
				"30s",
				"1m",
				"5m",
				"15m",
				"30m",
				"1h",
				"2h",
				"1d"},
			TimeOptions: []string{
				"5m",
				"15m",
				"1h",
				"6h",
				"12h",
				"24h",
				"2d",
				"7d",
				"30d"},
		},
		Templating: GrafanaDashboardConfigTemplating{
			List: []GrafanaDashboardConfigTemplatingListItem{
				{
					Current: GrafanaDashboardConfigTemplatingListItemCurrent{
						Text:  ".*",
						Value: ".*",
					},
					Name: "node",
					Options: []GrafanaDashboardConfigTemplatingListItemOptions{
						{
							Selected: true,
							Text:     ".*",
							Value:    ".*",
						},
					},
					Query: ".*",
					Type:  "constant",
				},
			},
		},
		Refresh:       "5s",
		SchemaVersion: 1,
		Version:       1,
		Panels: []GrafanaDashboardConfigPanel{
			{
				Id: 1,
				GridPos: GrafanaDashboardConfigPanelGridPos{
					H: 1,
					W: 24,
					X: 0,
					Y: 0,
				},
				Title: "Cluster Health",
				Type:  "row",
			},
			{
				Id: 2,
				GridPos: GrafanaDashboardConfigPanelGridPos{
					H: 6,
					W: 4,
					X: 0,
					Y: 1,
				},
				Title: "Cluster Pod Usage",
				Type:  "gauge",
				Options: GrafanaDashboardConfigPanelOptions{
					Orientation: "horizontal",
					ReduceOptions: GrafanaDashboardConfigPanelOptionsReduceOptions{
						Calcs:  []string{"mean"},
						Fields: "",
						Values: false,
					},
					ShowThresholdLabels:  false,
					ShowThresholdMarkers: true,
				},
				Colors:     []string{greenColor, orangeColor, redColor},
				DataSource: "${DS_PROMETHEUS}",
				Format:     "percentunit",
				Gauge: GrafanaDashboardConfigPanelGauge{
					MaxValue:         100,
					MinValue:         0,
					Show:             true,
					ThresholdMarkers: true,
				},
				MaxDataPoints: 100,
				Targets: []GrafanaDashboardConfigPanelTarget{
					{
						Expr:           "sum(kube_pod_info{node=~\"$node\"}) / sum(kube_node_status_allocatable_pods{node=~\".*\"})",
						Format:         "time_series",
						IntervalFactor: 1,
						RefId:          "A",
					},
				},
				Sparkline: GrafanaDashboardConfigPanelSparkline{
					FillColor: lightBlueColor,
					Full:      false,
					LineColor: whiteColor,
					Show:      false,
				},
				RangeMaps: []GrafanaDashboardConfigPanelRangeMap{
					{
						From: "null",
						Text: "N/A",
						To:   "null",
					},
				},
			},
		},
		Annotations: []string{},
		Inputs: []GrafanaDashboardConfigInput{
			{
				Name:       "DS_PROMETHEUS",
				Label:      "prometheus",
				Type:       "datasource",
				PluginId:   "prometheus",
				PluginName: "Prometheus",
			},
		},
	}
}
