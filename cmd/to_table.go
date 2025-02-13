package cmd

import (
	"encoding/json"
	"fmt"
	"io"

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

		for _, v := range inputJson {
			ob, ok := v.(map[string]interface{})
			if !ok {
				return merry.New("failed to cast inputJson to map[string]interface{}")
			}
			for ob_key := range ob {
				headers = append(headers, ob_key)
			}
		}

		if maxColumns > 0 {
			headers = headers[:maxColumns]
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

var maxColumns int

func init() {
	rootCmd.AddCommand(toTableCmd)
	toTableCmd.Flags().IntVarP(&maxColumns, "max-columns", "m", 0, "maximum number of columns allowed in the table")
}
