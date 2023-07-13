package k8shelper

import (
	"io/ioutil"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetDockerSecret() (*v1.Secret, error) {
	conf, err := ioutil.ReadFile("./data/docker-config.json")
	if err != nil {
		return nil, err
	}
	secret := &v1.Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "docker-config",
		},
		Data: map[string][]byte{
			".dockerconfigjson": conf,
		},
		Type: "kubernetes.io/dockerconfigjson",
	}
	return secret, nil
}
