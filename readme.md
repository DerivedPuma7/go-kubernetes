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

## watch hpa
```console
watch -n1 kubectl get hpa
```

## inspect pod just like docker container
kubectl exec -it podname -- bash

# fortio
## run stress test 
```console
kubectl run -it fortio --rm --image=fortio/fortio -- load -qps 900 -t 120s -c 80 "http://goserver:8080/healthz"
```