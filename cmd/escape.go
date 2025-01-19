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
	Use:   "escape",
	Short: "interpret the escape sequences",
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// escapeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// escapeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
