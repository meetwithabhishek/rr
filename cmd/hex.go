package cmd

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/ansel1/merry/v2"
	"github.com/spf13/cobra"
)

// hextCmd represents the hex command
var hexCmd = &cobra.Command{
	Use:     "hex",
	Aliases: []string{"x", "xxd"},
	Short:   "hex encode/decode",
	Long:    "This command performs hex encode by default. Specify --decode flag to perform hex decode. Input is taken from stdin",
	RunE: func(cmd *cobra.Command, args []string) error {
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			return merry.Prepend(err, "error while reading from stdin")
		}
		if decode {
			o, err := hex.DecodeString(string(b))
			if err != nil {
				return merry.Prepend(err, "failed to decode hex")
			}
			print(string(o))
		} else {
			o := hex.EncodeToString(b)
			print(o)
		}
		return nil
	},
}

var decode bool

func init() {
	rootCmd.AddCommand(hexCmd)
	hexCmd.Flags().BoolVarP(&decode, "decode", "d", false, "decode hex")
}

func print(a ...any) {
	if avoidNewLine {
		fmt.Print(a...)
	} else {
		fmt.Println(a...)
	}
}
