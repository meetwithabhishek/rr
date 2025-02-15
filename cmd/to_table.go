package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	merry "github.com/ansel1/merry/v2"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
)

var style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA"))

// toTableCmd represents the to-table command
var toTableCmd = &cobra.Command{
	Use:     "to-table",
	Aliases: []string{"tt"},
	Short:   "Convert the JSON array to table",
	Long:    "Take the JSON from the stdin and give a table in the output",
	RunE: func(cmd *cobra.Command, args []string) error {

		// this does the trick
		var inputReader io.Reader = cmd.InOrStdin()

		inputJson := make([]interface{}, 0)

		err := json.NewDecoder(inputReader).Decode(&inputJson)
		if err != nil {
			return err
		}

		var headers []string

		// gettings the list of headers from the first object in the JSON array
		v := inputJson[0]

		ob, ok := v.(map[string]interface{})
		if !ok {
			return merry.New("failed to cast inputJson to map[string]interface{}")
		}
		for ob_key := range ob {
			headers = append(headers, ob_key)
		}

		if maxColumnsOpt > 0 {
			if len(headersAtleastOpt) > 0 {

				var chosenHeaders []string
				var leftHeaders []string
				for _, hdr := range readCommaSeparatedString(headersAtleastOpt) {
					for _, v := range headers {
						if hdr == v {
							chosenHeaders = append(chosenHeaders, v)
						} else {
							leftHeaders = append(leftHeaders, v)
						}
					}
				}

				headers = append(chosenHeaders, leftHeaders...)
			}

			headers = headers[:maxColumnsOpt]
		}

		if len(headersOpt) > 0 {
			var recomputedHeaders []string

			for _, v := range readCommaSeparatedString(headersOpt) {
				for _, hdr := range headers {
					if v == hdr {
						recomputedHeaders = append(recomputedHeaders, hdr)
					}
				}
			}

			headers = recomputedHeaders
		}

		t := table.New().
			Border(lipgloss.NormalBorder()).
			BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
			StyleFunc(func(row, col int) lipgloss.Style {
				switch {
				case row == 0:
					return style
				default:
					return style
				}
			}).Headers(headers...)

		for _, v := range inputJson {
			ob, ok := v.(map[string]interface{})
			if !ok {
				return merry.New("failed to cast inputJson to map[string]interface{}")
			}

			var row []string
			for _, v := range headers {
				row = append(row, fmt.Sprintf("%v", ob[v]))
			}

			t.Row(row...)
		}

		fmt.Println(t)

		return nil
	},
}

func readCommaSeparatedString(s string) []string {
	values := strings.Split(s, ",")

	var finalValues []string
	for _, v := range values {
		finalValues = append(finalValues, strings.TrimSpace(v))
	}

	return finalValues
}

var maxColumnsOpt int
var headersOpt string
var headersAtleastOpt string

func init() {
	rootCmd.AddCommand(toTableCmd)
	toTableCmd.Flags().IntVarP(&maxColumnsOpt, "max-columns", "m", 0, "maximum number of columns allowed in the table")
	toTableCmd.Flags().StringVarP(&headersOpt, "headers", "", "", "headers/columns of the table, its a comma separated list")
	toTableCmd.Flags().StringVarP(&headersAtleastOpt, "headers-atleast", "", "", `if max-columns option is used, then atleast include 
following headers in the table, its a comma separated list`)
}
