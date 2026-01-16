package cmd

import (
	"bufio"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// replaceCmd represents the replace command
var replaceCmd = &cobra.Command{
	Use:     "replace [old_string] [new_string]",
	Aliases: []string{"r"},
	Short:   "replace strings",
	Args:    cobra.ExactArgs(2),
	Long:    "This command replaces strings. Input is taken from stdin.",
	Example: `rr replace old_string new_string < input.txt`,
	RunE: func(cmd *cobra.Command, args []string) error {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			print(strings.ReplaceAll(scanner.Text(), args[0], args[1]))
		}

		err := scanner.Err()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(replaceCmd)
}
