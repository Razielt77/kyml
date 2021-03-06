package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmdOptions struct {
	Verbose bool
}

var rootCmd = &cobra.Command{
	Use:  "kubectl yaml-writer",
	Long: "Use kubectl yaml-writer to manipulate k8s GitOps yaml files.\nClone GitOps repo first then run kubectl yaml-writer on the cloned directory to update the k8s resources to its desired state.\nLastly commit your changes and trigger your GitOps sync tool",
}

// Execute - execute the root command
func Execute() {
	err := rootCmd.Execute()
	dieOnError(err)
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&rootCmdOptions.Verbose, "verbose", false, "Set to get more detailed output")
}
