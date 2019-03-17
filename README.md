# kaniko-example

This example was suitable for beginners who want to start using kaniko with local environment. It requires only general kubernetes and docker usage, and NO google cloud or aws or any third part cloud service support. It aims to help user establish a quick start successfully and deep dive after that.

## Table of Content
1. [Prerequisities](#Prerequisities)
2. [Getting Started](#Getting Started)
3. [Prepare config files for kaniko](#Prepare-config-files-for-kaniko)
4. [Prepare the mounted directory in minikube](#Prepare-the-mounted-directory-in-minikube)
5. [Create a Secret that holds your authorization token](#Create-a-Secret-that-holds-your-authorization-token)
6. [Create resources in kubernetes](#Create-resources-in-kubernetes)
7. [Pull the image and test](#Pull-the-image-and-test)



## Prerequisities

- Kubernetes basic knowledge about [container](https://kubernetes.io/docs/concepts/overview/what-is-kubernetes/), [pod](https://kubernetes.io/docs/concepts/workloads/pods/pod/), [volume](https://kubernetes.io/docs/concepts/storage/volumes/) and [secret](https://kubernetes.io/docs/concepts/configuration/secret/).

- Docker basic usage and a [dockerhub](https://hub.docker.com/) account to push built image public.

## Getting Started

* Install Docker - https://www.docker.com/get-started 
* Install kubectl - https://kubernetes.io/docs/tasks/tools/install-kubectl/ 
* Install Minikube - https://kubernetes.io/docs/tasks/tools/install-minikube/ 

## Prepare config files for kaniko

After installing minikube, we should prepare several config files to create resources in kubernetes, which are:
- [pod.yml] is for starting a kaniko container to build the example image. 

## Prepare workspace using configMap

```
# Create configmap for example1 workspace
kubectl create configmap dockerwsfile --from-file=example1/
configmap/dockerwsfile created
```

## Create a Secret that holds your authorization token
A Kubernetes cluster uses the Secret of docker-registry type to authenticate with a docker registry to push an image.

Create this Secret, naming it regcred:

```
kubectl create secret docker-registry regcred --docker-server=https://index.docker.io/v1/ --docker-username=<your-name> --docker-password=<your-pword> --docker-email=<your-email>
```

It will be used in pod.yml config.

## Create resources in kubernetes

```
# create pod
$ kubectl create -f pod.yml
pod/kaniko created
$ kubectl get pods
NAME     READY   STATUS              RESTARTS   AGE
kaniko   0/1     ContainerCreating   0          7s

# check whether the build complete and show the build logs
$ kubectl get pods
NAME     READY   STATUS      RESTARTS   AGE
kaniko   0/1     Completed   0          34s
$ kubectl logs kaniko
...
INFO[0075] Skipping paths under /workspace, as it is a whitelisted directory 
INFO[0075] WORKDIR /app                                 
INFO[0075] cmd: workdir                                 
INFO[0075] Changed working directory to /app            
INFO[0075] Creating directory /app                      
INFO[0075] Taking snapshot of files...                  
INFO[0075] COPY --from=build-env /app/hello /app/hello  
INFO[0075] Taking snapshot of files...                  
INFO[0076] RUN chown nobody:nogroup /app                
INFO[0076] cmd: /bin/sh                                 
INFO[0076] args: [-c chown nobody:nogroup /app]         
INFO[0076] Taking snapshot of full filesystem...        
INFO[0076] Skipping paths under /kaniko, as it is a whitelisted directory 
INFO[0076] Skipping paths under /var/run, as it is a whitelisted directory 
INFO[0076] Skipping paths under /root, as it is a whitelisted directory 
INFO[0076] Skipping paths under /dev, as it is a whitelisted directory 
INFO[0076] Skipping paths under /sys, as it is a whitelisted directory 
INFO[0076] Skipping paths under /proc, as it is a whitelisted directory 
INFO[0076] Skipping paths under /workspace, as it is a whitelisted directory 
INFO[0076] USER nobody                                  
INFO[0076] cmd: USER                                    
INFO[0076] ENTRYPOINT /app/hello                        
2019/03/17 03:29:09 existing blob: sha256:8e402f1a9c577ded051c1ef10e9fe4492890459522089959988a4852dee8ab2c
2019/03/17 03:29:10 pushed blob sha256:17f2f052ee5b8567a895aa6360cfa08e853ad62fa3c3b1b2a98d50ab60daabc3
2019/03/17 03:29:10 pushed blob sha256:1796299b6c4c91574d557bdc476a19a346d4cdd863c3608dfb4e7cc3048fb9eb
2019/03/17 03:29:10 pushed blob sha256:26902e01f9b230f25ce2da6cbede03c0a9055f18e6b42a7c83dcd17eef0b3d10
2019/03/17 03:29:11 pushed blob sha256:c8aa7014be74f93e51fbccd4b3a99e266f0a71b79bf76dc44d09f2c76eb6de37
2019/03/17 03:29:11 index.docker.io/siddhardhan/kaniko-example:latest: digest: sha256:9b71ff9e213dd4ae4883fd6e872fb4c96f2c19b2c528aa6e9e5d542209937cbb size: 911
```

## Pull the image and test

If as expected, the kaniko will build image and push to dockerhub successfully. Pull the image to local and run it to test:

```
$ sudo docker run -it <user-name>/<repo-name>
Unable to find image '<user-name>/<repo-name>:latest' locally
latest: Pulling from <user-name>/<repo-name>
32802c0cfa4d: Already exists
da1315cffa03: Already exists
fa83472a3562: Already exists
f85999a86bef: Already exists
Digest: sha256:5706190951cbf69cfd880f010da015703a1fdca0f2f25b992c3b9193116bbd88
Status: Downloaded newer image for <user-name>/<repo-name>:latest
hello
```
