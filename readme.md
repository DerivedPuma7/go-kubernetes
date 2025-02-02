# docker

## build image
```console
docker build -t derivedpuma7/kubernetes-go:version .
```

## push to docker hub
```console
docker push derivedpuma7/kubernetes-go:version
```

# kubernetes
## restart resource
### restart deployment
```console
kubectl rollout restart deployment/goserver
```

## apply deplyment and watch
```console
kubectl apply -f k8s/deployment.yml && watch -n1 kubectl get pods
```

## pod resource use
```console
kubectl top pod podname
```