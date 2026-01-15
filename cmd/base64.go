package cmd

import (
	"encoding/base64"
	"io"
	"os"

	"github.com/ansel1/merry/v2"
	"github.com/spf13/cobra"
)

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:     "base64",
	Aliases: []string{"b"},
	Short:   "base64 encode/decode",
	Long: `This command performs base64 encode by default. Specify --decode flag to perform base64 decode. 
Input is taken from stdin by default. Input can be specified with --input flag as well.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var b []byte
		var err error
		if input != "" {
			b = []byte(input)
		} else {
			b, err = io.ReadAll(os.Stdin)
			if err != nil {
				return merry.Prepend(err, "error while reading from stdin")
			}
		}
		if decode {
			o, err := base64.StdEncoding.DecodeString(string(b))
			if err != nil {
				return merry.Prepend(err, "failed to decode base64")
			}
			print(string(o))
		} else {
			o := base64.StdEncoding.EncodeToString(b)
			print(o)
		}
		return nil
	},
}

var input string

func init() {
	rootCmd.AddCommand(base64Cmd)
	base64Cmd.Flags().BoolVarP(&decode, "decode", "d", false, "decode base64")
	base64Cmd.Flags().StringVarP(&input, "input", "i", "", "input string (if not provided, stdin is used)")
}
