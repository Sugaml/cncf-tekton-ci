package models

import (
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	"k8s.io/client-go/kubernetes"
)

type Pipeline struct {
	ClientSet    *kubernetes.Clientset
	Input        *CIRequest
	TektonClient *versioned.Clientset
}

type CIRequest struct {
	EnvironmentId     int64                   `json:"environment_id"`
	Namespace         string                  `json:"namespace"`
	ConfigPath        string                  `json:"config_path"`
	Name              string                  `json:"name"`
	ImageRepoUsername string                  `json:"image_repo_username"`
	ImageRepoPassword string                  `json:"image_repo_password"`
	ImageRepoService  string                  `json:"image_repo_service"`
	ImageRepoProject  string                  `json:"image_repo_project"`
	ImageRepoProvider string                  `json:"image_repo_provider"`
	GitUrl            string                  `json:"git_url"`
	GitBranch         string                  `json:"git_branch"`
	GitUserName       string                  `json:"git_username"`
	RepositoryImage   RepositoryImage         `json:"repository_image"`
	GitAccessToken    string                  `json:"git_access_token"`
	SubDirectory      string                  `json:"sub_directory"`
	PluginUrl         string                  `json:"plugin_url"`
	Author            string                  `json:"author"`
	CommitMessage     string                  `json:"commit_message"`
	BaseImage         string                  `json:"base_image"`
	BaseTag           string                  `json:"base_tag"`
	BuildScript       string                  `json:"base_script"`
	RunScript         string                  `json:"run_script"`
	CISteps           interface{}             `json:"ci_steps"`
	CiFiles           *map[string]interface{} `json:"ci_files"`
}

type RepositoryImage struct {
	Name          string  `json:"name"`
	Repository    string  `json:"repository"`
	Tag           string  `json:"tag"`
	CommitMessage *Commit `json:"commit_message"`
}

type Commit struct {
	SHA     string `json:"sha"`
	Message string `json:"message"`
	Author  string `json:"author"`
	Time    string `json:"time"`
}
