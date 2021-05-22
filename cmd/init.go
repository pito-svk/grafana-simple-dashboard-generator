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
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/pito-svk/grafana-simple-dashboard-generator/dashboard"
	"github.com/spf13/cobra"
)

var (
	mode     string
	output   string
	filepath string
	title    string
	style    string
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate new grafana dashboard",
	RunE: func(cmd *cobra.Command, args []string) error {
		if output != "console" && output != "file" {
			return fmt.Errorf("invalid value for output. Allowed values: [\"console\", \"file\"]")
		}

		if mode != "basic" {
			return fmt.Errorf("invalid value for mode. Allowed values: [\"basic\"]")
		}

		if output == "file" && filepath == "" {
			return fmt.Errorf("filepath is required when output is \"file\"")
		}

		dashboardParams := dashboard.GrafanaDashboardParams{Title: title}

		dashboard := dashboard.GenerateDashboard(&dashboardParams)
		jsonRes, _ := json.MarshalIndent(dashboard, "", "	")

		fmt.Println(string(jsonRes))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&mode, "mode", "m", "basic", "dashboard mode")
	initCmd.Flags().StringVarP(&output, "output", "o", "console", "specify output mode")
	initCmd.Flags().StringVarP(&filepath, "filepath", "f", "", "specify a filepath to generate dashboards")
	initCmd.Flags().StringVarP(&title, "title", "t", "Server monitoring", "specify a title for dashboard")
	initCmd.Flags().StringVarP(&style, "style", "s", "light", "specify a style for dashboard")
}
