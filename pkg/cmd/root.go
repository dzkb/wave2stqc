package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	sampleRate int
)

var rootCmd = &cobra.Command{
	Use:   "wave2stqc",
	Short: "wave2stqc processes raw waveform and detects STQC tones",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		fmt.Println(sampleRate)
	},
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&sampleRate, "sample-rate", "s", 0, "Set sample rate")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
