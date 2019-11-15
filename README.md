# Kube-Scan
Try our free Kubernetes risk assessment tool today. No data leaves your cluster. We do not collect any information. Run it on any cluster at any time. See everything you need to get started below. For more information on Octarine head to https://www.octarinesec.com. 

## Quickstart
```bash
kubectl apply -f https://raw.githubusercontent.com/octarinesec/kube-scan/master/kube-scan.yaml
kubectl port-forward --namespace kube-scan svc/kube-scan-ui 8080:80
```

Then set your browser to `http://localhost:8080`.

## Using a load-balancer service
```bash
kubectl apply -f https://raw.githubusercontent.com/octarinesec/kube-scan/master/kube-scan-lb.yaml
kubectl -n kube-scan get service kube-scan-ui -o jsonpath={..hostname}
```

Then set your browser to that hostname.
