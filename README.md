# Tekton-CICD : Cloud Native CICD

## ***Installing Tekton Pipelines on Kubernetes***
### *To install Tekton Pipelines on a Kubernetes cluster:*

    Run the following command to install Tekton Pipelines and its dependencies:

    kubectl apply --filename https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml

##  ***Install Tekton dashboard***

## Install from here

    kubectl apply --filename https://storage.googleapis.com/tekton-releases/dashboard/latest/tekton-dashboard-release.yaml

## *Enter the Dashboard via:*
 
    kubectl proxy

    http://localhost:8001/api/v1/namespaces/tekton-pipelines/services/tekton-dashboard:http/proxy/

## **Create Namespace tekton-ci**
    ***Run the following command to create namespace tekton-ci***
    kubectl create namespace tekton-ci

## **Create Secret tekton-ci**
    ***Run the following command to create secret docker registry***
    kubectl create secret docker-registry docker-config --docker-server=index.docker.io --docker-username=<DOCKER_USERNAME> --docker-password=<DOCKER_PASSWORD> --docker-email=<DOCKER_EMAIL> -n tekton-

## **Create Resource tekton-ci**
    ***Run the following command to create resource (PVC)***
   kubectl apply -f https://github.com/babulalt/tekton-ci/blob/main/manifest/resources.yaml -n tekon-ci

## **Create Task CRD tekton-ci**
    ***Run the following command to create task such as git-clone and build-push image***
   kubectl apply -f https://github.com/babulalt/tekton-ci/blob/main/manifest/git-clone.yaml -n tekon-ci

   kubectl apply -f https://github.com/babulalt/tekton-ci/blob/main/manifest/g0-lint.yaml -n tekon-ci
   
   kubectl apply -f https://github.com/babulalt/tekton-ci/blob/main/manifest/build-push.yaml -n tekon-ci

## **Create Pipeline and Pipeline Run CRD tekton-ci**
    ***Run the following command to create pipeline and pipelinerun***
    kubectl apply -f  https://github.com/babulalt/tekton-ci/blob/main/manifest/run.yaml -n tekon-ci

### See the output
    Kubectl get pod -n tekton-ci