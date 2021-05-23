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
	From string `json:"from"`
	To   string `json:"to"`
}

type GrafanaDashboardConfigTimepicker struct {
	RefreshIntervals []string `json:"refresh_intervals"`
	TimeOptions      []string `json:"time_options"`
}

type GrafanaDashboardConfigTemplatingListItemCurrent struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

type GrafanaDashboardConfigTemplatingListItemOptions struct {
	Selected bool   `json:"selected"`
	Text     string `json:"text"`
	Value    string `json:"value"`
}

type GrafanaDashboardConfigTemplatingListItem struct {
	Current *GrafanaDashboardConfigTemplatingListItemCurrent `json:"current"`
	Name    string                                           `json:"name"`
	Query   string                                           `json:"query"`
	Type    string                                           `json:"type"`
	Options *GrafanaDashboardConfigTemplatingListItemOptions `json:"options"`
}

type GrafanaDashboardConfigTemplating struct {
	List []GrafanaDashboardConfigTemplatingListItem `json:"list"`
}

type GrafanaDashboardConfigPanelsGridPos struct {
	H int `json:"h"`
	W int `json:"w"`
	X int `json:"x"`
	Y int `json:"y"`
}

type GrafanaDashboardConfigPanel struct {
	GridPos GrafanaDashboardConfigPanelsGridPos `json:"gridPos"`
	Rest_   map[string]interface{}              `json:"-"`
}

type GrafanaDashboardConfig struct {
	Id           *string                          `json:"id"`
	Uid          *string                          `json:"uid"`
	Title        string                           `json:"title"`
	Style        string                           `json:"style"`
	Timezone     string                           `json:"timezone"`
	Editable     bool                             `json:"editable"`
	GraphTooltip int                              `json:"graphTooltip"`
	Time         GrafanaDashboardConfigTime       `json:"time"`
	Timepicker   GrafanaDashboardConfigTimepicker `json:"timepicker"`
	Templating   GrafanaDashboardConfigTemplating `json:"templating"`
	// Annotations
	Refresh       string                        `json:"refresh"`
	SchemaVersion int                           `json:"schemaVersion"`
	Version       int                           `json:"version"`
	Panels        []GrafanaDashboardConfigPanel `json:"panels"`
}

func GenerateDashboard(params *GrafanaDashboardParams) interface{} {
	return &GrafanaDashboardConfig{
		Id:           nil,
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
		Templating:    GrafanaDashboardConfigTemplating{List: []GrafanaDashboardConfigTemplatingListItem{}},
		Refresh:       "5s",
		SchemaVersion: 1,
		Version:       1,
		Panels:        []GrafanaDashboardConfigPanel{},
	}
}
