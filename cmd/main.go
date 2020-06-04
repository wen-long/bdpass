package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/winterssy/bdpass"
	"github.com/winterssy/bdpass/encoder"
)

var (
	format string
)

var cmd = &cobra.Command{
	Use: "bdpass",
	Long: `
  _      _     _ ___ ___   _   ___ ___  _   _ 
 | |    /_\ _ | |_ _| _ ) /_\ |_ _|   \| | | |
 | |__ / _ \ || || || _ \/ _ \ | || |) | |_| |
 |____/_/ \_\__/|___|___/_/ \_\___|___/ \___/

`,
}

func init() {
	cmd.Flags().StringVarP(&format, "format", "f", "std", "coded format (std|pdl|pcs)")
	cmd.Run = Run
}

func Run(_ *cobra.Command, args []string) {
	var enc encoder.Encoder
	switch format {
	case "pdl":
		enc = &encoder.PDL{}
	case "pcs":
		enc = &encoder.PCS{}
	default:
		enc = &encoder.STD{}
	}
	for _, filename := range args {
		meta, err := bdpass.Stat(filename)
		if err != nil {
			fmt.Printf("bdpass: %s: %s\n", filename, err)
			continue
		}
		fmt.Println(enc.Encode(meta))
	}
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
