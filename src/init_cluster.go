package src

import (
	"fmt"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// returns the config and clientset and error
func (p *PipelineController) InitCluster() (*rest.Config, *kubernetes.Clientset, error) {
	path := p.Input.ConfigPath
	if path != "" {
		if _, err := os.Stat(path); err != nil {
			return nil, nil, fmt.Errorf("cluster config not exits. %s", path)
		}
	}
	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		return nil, nil, err
	}
	client, err := kubernetes.NewForConfig(config)
	return config, client, err
}
