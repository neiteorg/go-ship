package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"neite.dev/go-ship/internal/runner"
)

func init() {
	AppCmd.AddCommand(stopCmd)
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop app container on servers",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := runner.New()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return
		}

		if err := r.StopContainer(); err != nil {
			fmt.Fprint(os.Stderr, err)
			return

		}
	},
}
