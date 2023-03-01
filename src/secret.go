package src

import (
	"context"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Clients) CreateSecret(namespace string) (*v1.Secret, error) {
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
	csecret, err := c.KubeClient.CoreV1().Secrets(namespace).Create(context.Background(), secret, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return csecret, nil
}

func (c *Clients) CreateSecrets() (*v1.Secret, error) {
	conf, err := ioutil.ReadFile("./manifest/secret.yaml")
	if err != nil {
		return nil, err
	}
	var secret *v1.Secret
	fmt.Println(string(conf))
	secrets := &v1.Secret{
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
	cytes, err := yaml.Marshal(secrets)

	fmt.Println(string(cytes), err)
	err = yaml.Unmarshal(cytes, &secret)

	fmt.Println(secret, err)

	// csecret, err := c.KubeClient.CoreV1().Secrets(namespace).Create(context.Background(), secret, metav1.CreateOptions{})
	// if err != nil {
	// 	return nil, err
	// }
	return secret, nil
}
