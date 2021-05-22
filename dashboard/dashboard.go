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

type GrafanaDashboardConfig struct {
	Id           *string `json:"id"`
	Uid          *string `json:"uid"`
	Title        *string `json:"title"`
	Style        *string `json:"style"`
	Timezone     string  `json:"timezone"`
	Editable     bool    `json:"editable"`
	GraphTooltip int     `json:"graphTooltip"`
}

func GenerateDashboard(params *GrafanaDashboardParams) interface{} {
	return &GrafanaDashboardConfig{
		Id:           nil,
		Title:        &params.Title,
		Style:        &params.Style,
		Timezone:     "browser",
		Editable:     true,
		GraphTooltip: 1,
	}
}
