package cmd

import (
	"github.com/sirupsen/logrus"

	"github.com/thomas-maurice/thermal-printer/common"
	proto "github.com/thomas-maurice/thermal-printer/go"
)

func getClient() (proto.PrintServiceClient, error) {
	if useTLS {
		tlsConfig, err := common.GetTLSConfig(caCert, certFile, keyFile, insecureSkipVerify)

		if err != nil {
			logrus.WithError(err).Fatal("Could not setup TLS client")
		}
		return common.GetClient(controllerAddress, useTLS, tlsConfig)
	}

	return common.GetClient(controllerAddress, useTLS, nil)
}
