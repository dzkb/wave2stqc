package cmd

import (
	"fmt"
	"os"

	"github.com/dzkb/wave2stqc/pkg/wave"
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

		ch := make(chan []byte)
		go wave.ReadWave(os.Stdin, ch, 128)

		for {
			data := <-ch
			if len(data) > 0 {
				fmt.Println(data)
			}
		}
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
