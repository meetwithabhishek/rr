package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// escapeCmd represents the escape command
var escapeCmd = &cobra.Command{
	Use:     "escape",
	Aliases: []string{"e"},
	Short:   "interpret the escape sequences",
	RunE: func(cmd *cobra.Command, args []string) error {
		replacer := strings.NewReplacer(
			`\\`, `\`,
			`\n`, "\n",
			`\t`, "\t",
			`\r`, "\r",
			`\b`, "\b",
			`\f`, "\f",
		)

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Print(replacer.Replace(scanner.Text()))
		}
		err := scanner.Err()
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(escapeCmd)
}
