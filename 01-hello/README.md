Install docker-desktop: https://docs.docker.com/desktop/install/ubuntu/

```
docker compose version
docker --version
```

## Build Container

When you build an image, by default it just gets a hexadecimal ID, which you can
use to refer it later. You can add a tag name with the `-t` flag.

```
docker image build -t 01-hello .
```

Run with tag name

```
docker container run -p 9999:8888 01-hello
```

To tell a docker to forward port, you can use `-p` switch.

```
doccker container run -p HOST_PORT:CONTAINER_PORT ...
```

## Container Registries

Create account on https://hub.docker.com

```
docker login
username: DOCKER_ID
password: DOCKER_PASSWORD
```

## Naming and pushing your image

In order to able to push a local image to registry, you need to name it using
this format: `DOCKER_ID/tagname`

```
docker image tag 01-hello DOCKER_ID/01-hello
docker image push DOCKER_ID/01-hello
```

Run your image:

```
docker container run -p 9999:8888 DOCKER_ID/01-hello
```

## Hello Kubernetes

Install minikube: https://minikube.sigs.k8s.io/docs/start/

```
alias kubectl="minikube kubectl --"
```

Start, pause or stop minikube cluster

```
minikube start
minikube pause
minikube stop
```

### Run the demo app

```
kubectl run demo --image=DOCKER_ID/01-hello --port=9999 --labels app=demo
```

It's basically the kubernetes equivalent of the `docker container run` command.

Forward port 9999 on your local machine to the container's port 8888.

```
kubectl port-forward pod/demo 9999:8888
```

List running pods:

```
kubectl get pods
```

Clean up your demo container:

```
kubectl delete pod demo
```
