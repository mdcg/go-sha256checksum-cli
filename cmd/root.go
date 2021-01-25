package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "sha256checksum",
		Short: "Simple CLI to compare the checksum of two files.",
		Long: `Checksums are a useful way to ensure that a file does not have an error.
If a random error occurs due to download or hard drive problems, the resulting
checksum will be different, even if it is just a small error. In this way, this
CLI will compare two files in order to ensure that they are the same or not.`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}
