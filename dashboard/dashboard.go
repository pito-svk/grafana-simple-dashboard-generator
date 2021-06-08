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

type GrafanaDashboardConfigPanelFieldConfigDefaultsColor struct {
	Mode string `json:"mode,omitempty"`
}

type GrafanaDashboardConfigPanelFieldConfigDefaultsThresholdsStep struct {
	Color string `json:"color,omitempty"`
	Value int    `json:"value,omitempty"`
}

type GrafanaDashboardConfigPanelFieldConfigDefaultsThresholds struct {
	Mode  string                                                         `json:"mode,omitempty"`
	Steps []GrafanaDashboardConfigPanelFieldConfigDefaultsThresholdsStep `json:"steps,omitempty"`
	Max   int                                                            `json:"max,omitempty"`
	Unit  string                                                         `json:"unit,omitempty"`
}

type GrafanaDashboardConfigPanelFieldConfigDefaults struct {
	Color      GrafanaDashboardConfigPanelFieldConfigDefaultsColor      `json:"color,omitempty"`
	Thresholds GrafanaDashboardConfigPanelFieldConfigDefaultsThresholds `json:"thresholds,omitempty"`
	Unit       string                                                   `json:"unit,omitempty"`
}

type GrafanaDashboardConfigPanelFieldConfig struct {
	Defaults GrafanaDashboardConfigPanelFieldConfigDefaults `json:"defaults,omitempty"`
}

type GrafanaDashboardConfigPanel struct {
	Id            int                                    `json:"id,omitempty"`
	GridPos       GrafanaDashboardConfigPanelGridPos     `json:"gridPos,omitempty"`
	Title         string                                 `json:"title,omitempty"`
	Type          string                                 `json:"type,omitempty"`
	MaxDataPoints int                                    `json:"maxDataPoints,omitempty"`
	Targets       []GrafanaDashboardConfigPanelTarget    `json:"targets,omitempty"`
	Options       GrafanaDashboardConfigPanelOptions     `json:"options,omitempty"`
	FieldConfig   GrafanaDashboardConfigPanelFieldConfig `json:"fieldConfig,omitempty"`
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
	clusterCPUExpr := "sum(rate(container_cpu_usage_seconds_total{container!=\"POD\",container!=\"\"}[5m]))"
	// clusterMemoryExpr := ""

	templating := GrafanaDashboardConfigTemplating{
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
	}

	clusterMetricsRow := GrafanaDashboardConfigPanel{
		Id: 1,
		GridPos: GrafanaDashboardConfigPanelGridPos{
			H: 1,
			W: 24,
			X: 0,
			Y: 0,
		},
		Title: "Cluster metrics",
		Type:  "row",
	}

	clusterCPU := GrafanaDashboardConfigPanel{
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
		FieldConfig: GrafanaDashboardConfigPanelFieldConfig{
			Defaults: GrafanaDashboardConfigPanelFieldConfigDefaults{
				Color: GrafanaDashboardConfigPanelFieldConfigDefaultsColor{
					Mode: "thresholds",
				},
				Thresholds: GrafanaDashboardConfigPanelFieldConfigDefaultsThresholds{
					Mode: "percentage",
					Steps: []GrafanaDashboardConfigPanelFieldConfigDefaultsThresholdsStep{
						{
							Color: GREEN_COLOR,
						},
						{
							Color: ORANGE_COLOR,
							Value: 80,
						},
						{
							Color: RED_COLOR,
							Value: 90,
						},
					},
					Max:  100,
					Unit: "percentunit",
				},
				Unit: "percentunit",
			},
		},
		MaxDataPoints: 100,
		Targets: []GrafanaDashboardConfigPanelTarget{
			{
				Expr:           clusterCPUExpr,
				Format:         "time_series",
				IntervalFactor: 1,
				RefId:          "A",
			},
		},
	}

	return &GrafanaDashboardConfig{
		Title:        params.Title,
		Timezone:     "browser",
		Editable:     true,
		GraphTooltip: 1,
		Time:         CONFIG_TIME,
		Timepicker: GrafanaDashboardConfigTimepicker{
			RefreshIntervals: REFRESH_INTERVALS,
			TimeOptions:      TIME_OPTIONS,
		},
		Templating:    templating,
		Refresh:       "5s",
		SchemaVersion: 1,
		Version:       1,
		Panels: []GrafanaDashboardConfigPanel{
			clusterMetricsRow,
			clusterCPU,
			// clusterMemory
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
