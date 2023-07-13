package tekton_helper

import (
	"os"

	"github.com/sirupsen/logrus"
	v1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/scale/scheme"
)

func DecodeYamlToTask() (*v1.Task, error) {
	sch := runtime.NewScheme()
	_ = scheme.AddToScheme(sch)
	_ = v1.AddToScheme(sch)
	decode := serializer.NewCodecFactory(sch).UniversalDeserializer().Decode
	yamlInput := os.Getenv("TEKTON_FILE_PATH") + "/task.yaml"
	stream, err := os.ReadFile(yamlInput)
	if err != nil {
		logrus.Error("error occur on read file :: ", err)
		return nil, err
	}
	obj, gKV, err := decode(stream, nil, nil)
	if err != nil {
		return nil, err
	}
	if gKV.Kind == "Task" {
		task := obj.(*v1.Task)
		return task, nil
	}
	return nil, nil
}
