package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/openshift/machine-config-operator/pkg/controller/bootstrap"
	"github.com/openshift/machine-config-operator/pkg/version"
)

var (
	bootstrapCmd = &cobra.Command{
		Use:   "boostrap",
		Short: "Starts Machine Config Controller in bootstrap mode",
		Long:  "",
		Run:   runbootstrapCmd,
	}

	bootstrapOpts struct {
		manifestsDir   string
		destinationDir string
	}
)

func init() {
	rootCmd.AddCommand(bootstrapCmd)
	bootstrapCmd.PersistentFlags().StringVar(&bootstrapOpts.destinationDir, "dest-dir", "", "The destination dir where MCC writes the generated machineconfigs and machineconfigpools.")
	bootstrapCmd.PersistentFlags().StringVar(&bootstrapOpts.manifestsDir, "mainfest-dir", "", "The dir where MCC reads the controllerconfig, machineconfigpools and user-defined machineconfigs.")

}

func runbootstrapCmd(cmd *cobra.Command, args []string) {
	flag.Set("logtostderr", "true")
	flag.Parse()

	// To help debugging, immediately log version
	glog.Infof("Version: %+v", version.Version)

	if bootstrapOpts.manifestsDir == "" || bootstrapOpts.destinationDir == "" {
		glog.Fatalf("--dest-dir or --mainfest-dir not set")
	}

	if err := bootstrap.New(rootOpts.templates, bootstrapOpts.manifestsDir).Run(bootstrapOpts.destinationDir); err != nil {
		glog.Fatalf("error running MCC[BOOTSTRAP]: %v", err)
	}
}
