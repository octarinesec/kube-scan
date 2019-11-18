
<p align="center">
  <img src="./images/octarine_logo.png">
</p>

# Kube-Scan
Try our free Kubernetes risk assessment tool today. No data leaves your cluster. We do not collect any information. Run it on any cluster at any time. See everything you need to get started below. For more information on Octarine head to https://www.octarinesec.com. 

## Quickstart
```bash
kubectl apply -f https://raw.githubusercontent.com/octarinesec/kube-scan/master/kube-scan.yaml
kubectl port-forward --namespace kube-scan svc/kube-scan-ui 8080:80
```

Then set your browser to `http://localhost:8080`.

## Using a load-balancer service
* This method assumes you are using a cloud provider that provides load balancers.
```bash
kubectl apply -f https://raw.githubusercontent.com/octarinesec/kube-scan/master/kube-scan-lb.yaml
```
Then get the load-balancer address by
```bash
kubectl -n kube-scan get service kube-scan-ui -o jsonpath={..ip}
```
or
```bash
kubectl -n kube-scan get service kube-scan-ui -o jsonpath={..hostname}
```
depending on the load-balancer type.

Then set your browser to that address.

# Screenshots

![Risk score](https://info.octarinesec.com/hubfs/kube-scan.png)

![Risk details](https://info.octarinesec.com/hubfs/kube-scan-risk.png)
