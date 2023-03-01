package src

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned/typed/pipeline/v1alpha1"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned/typed/pipeline/v1beta1"
	resourceversioned "github.com/tektoncd/pipeline/pkg/client/resource/clientset/versioned"
	resourcev1alpha1 "github.com/tektoncd/pipeline/pkg/client/resource/clientset/versioned/typed/resource/v1alpha1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type Clients struct {
	KubeClient             kubernetes.Interface
	PipelineClient         v1beta1.TektonV1beta1Interface
	ClusterTaskClient      v1beta1.TektonV1beta1Interface
	TaskClient             v1beta1.TektonV1beta1Interface
	TaskRunClient          v1beta1.TektonV1beta1Interface
	PipelineRunClient      v1beta1.TektonV1beta1Interface
	PipelineResourceClient resourcev1alpha1.TektonV1alpha1Interface
	RunClient              v1alpha1.TektonV1alpha1Interface
}

func NewKubeConfig() (*rest.Config, error) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	cfg, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatalf("failed to create  config file  :: %v", err)
		return nil, err
	}
	return cfg, nil
}

func NewTektonClient() (*Clients, error) {
	c := &Clients{}
	cfg, err := NewKubeConfig()
	if err != nil {
		log.Fatalf("failed to create  config file  :: %v", err)
		return nil, err
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("failed to create kubeclient :: %v", err)
		return nil, err
	}
	c.KubeClient = kubeClient
	cs, err := versioned.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("failed to create pipeline clientset :: %v", err)
		return nil, err
	}
	rcs, err := resourceversioned.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("failed to create pipeline clientset :: %v", err)
		return nil, err
	}
	c.PipelineClient = cs.TektonV1beta1()
	c.ClusterTaskClient = cs.TektonV1beta1()
	c.TaskClient = cs.TektonV1beta1()
	c.TaskRunClient = cs.TektonV1beta1()
	c.PipelineRunClient = cs.TektonV1beta1()
	c.PipelineResourceClient = rcs.TektonV1alpha1()
	c.RunClient = cs.TektonV1alpha1()
	return c, nil
}
