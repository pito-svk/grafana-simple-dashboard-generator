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

const (
	GREEN_COLOR  = "rgb(41, 156, 70)"
	ORANGE_COLOR = "rgba(237, 129, 40, 0.89)"
	RED_COLOR    = "rgb(212, 74, 58)"
)

var REFRESH_INTERVALS = []string{
	"5s",
	"10s",
	"30s",
	"1m",
	"5m",
	"15m",
	"30m",
	"1h",
	"2h",
	"1d"}

var TIME_OPTIONS = []string{
	"5m",
	"15m",
	"1h",
	"6h",
	"12h",
	"24h",
	"2d",
	"7d",
	"30d"}

var CONFIG_TIME = GrafanaDashboardConfigTime{
	From: "now-6h",
	To:   "now",
}
