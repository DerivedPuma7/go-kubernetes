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
