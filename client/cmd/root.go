package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	marshaller         string
	controllerAddress  string
	useTLS             bool
	certFile           string
	caCert             string
	keyFile            string
	insecureSkipVerify bool
)

var rootCmd = &cobra.Command{
	Use:   "printctl",
	Short: "Printctl is used to interact with the thermal printer",
	Long:  ``,
}

// Execute is the main Cobra func
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Error("Could not execute")
		os.Exit(1)
	}
}

func init() {
	initPrintCmd()

	rootCmd.AddCommand(printCmd)
	rootCmd.PersistentFlags().StringVarP(&marshaller, "output", "o", "json", "Output marshaller, json or yaml")
	rootCmd.PersistentFlags().StringVarP(&controllerAddress, "controller", "c", "localhost:8069", "Controller address")
	rootCmd.PersistentFlags().StringVar(&caCert, "ca", "", "File containing the CA certificate")
	rootCmd.PersistentFlags().StringVar(&certFile, "cert", "", "File containing the client certificate")
	rootCmd.PersistentFlags().StringVar(&keyFile, "key", "", "File containing the client key")
	rootCmd.PersistentFlags().BoolVar(&useTLS, "tls", false, "Wether or not use TLS authentication")
	rootCmd.PersistentFlags().BoolVar(&insecureSkipVerify, "insecure-skip-verify", false, "Wether or not verify the CA certificates")
}
