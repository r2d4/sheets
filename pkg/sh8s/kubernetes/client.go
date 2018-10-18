package kubernetes

import (
	"fmt"

	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	// Initialize all known client auth plugins
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func GetClientset() (kubernetes.Interface, error) {
	config, err := getClientConfig()
	if err != nil {
		return nil, errors.Wrap(err, "getting client config for kubernetes client")
	}
	return kubernetes.NewForConfig(config)
}

func getClientConfig() (*restclient.Config, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, &clientcmd.ConfigOverrides{})
	clientConfig, err := kubeConfig.ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("Error creating kubeConfig: %s", err)
	}
	return clientConfig, nil
}
