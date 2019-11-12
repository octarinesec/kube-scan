# Kube-Scan
Octarine k8s cluster risk assesment tool  
https://www.octarinesec.com

## Quickstart
```bash
kubectl apply -f https://raw.githubusercontent.com/octarinesec/kube-scan/master/kube-scan.yaml
kubectl port-forward --namespace kube-scan svc/kube-scan-ui 8080:80
```

Then set your browser to `http://localhost:8080`.

## Using a load-balancer service
```bash
kubectl apply -f https://raw.githubusercontent.com/octarinesec/kube-scan/master/kube-scan-lb.yaml
kubectl -n octaudit get service kube-scan-ui -o jsonpath={..hostname}
```

Then set your browser to that hostname.
