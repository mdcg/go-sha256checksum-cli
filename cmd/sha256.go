package cmd

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	var cmdFileSHA256 = &cobra.Command{
		Use:   "file [path to the file]",
		Short: "Returns the checksum using SHA256 for a given file.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(FileSHA256(args[0]))
		},
	}
	var cmdCompareFiles = &cobra.Command{
		Use:   "compare [file paths]",
		Short: "Checks whether all files reported have the same checksum.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			haveTheSameChecksum := CompareFiles(args)
			if haveTheSameChecksum {
				fmt.Println("All files have the same checksum!")
			} else {
				fmt.Println("The files do not have the same checksum.")
			}
		},
	}
	rootCmd.AddCommand(cmdFileSHA256)
	rootCmd.AddCommand(cmdCompareFiles)
}

func FileSHA256(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func CompareFiles(paths []string) bool {
	checksums := make([]string, 0)
	for _, path := range paths {
		f256 := FileSHA256(path)
		checksums = append(checksums, f256)
	}
	for _, cs := range checksums {
		if checksums[0] != cs {
			return false
		}
	}
	return true
}
