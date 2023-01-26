package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wen-long/bdpass"
	"github.com/wen-long/bdpass/encoder"
	"io/fs"
	"os"
	"path/filepath"
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
	for _, entryName := range args {
		relName, err := filepath.Rel(".", entryName)
		if err != nil {
			fmt.Printf("# %s is not valid path, skip: %v\n", entryName, err)
			continue
		}

		filepath.Walk(relName, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("# %s is not accessable, skip: %v\n", path, err)
				return nil
			}
			if info.IsDir() {
				return nil
			}
			filename := path
			meta, err := bdpass.Stat(filename)
			if err != nil {
				fmt.Printf("#: %s: %s\n", filename, err)
			}
			fmt.Println(enc.Encode(meta))
			return nil
		})
	}
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
