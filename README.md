
<p align="center">
  <img src="./images/octarine_logo.png">
</p>

# Kube-Scan
Try our free Kubernetes risk assessment tool today. No data leaves your cluster. We do not collect any information. Run it on any cluster at any time. See everything you need to get started below. For more information on Octarine head to https://www.octarinesec.com. 

# Get the risk score of your workloads

kube-scan gives a risk score, from 0 (no risk) to 10 (high risk) for each workload. The risk is based on the runtime configuration of each workloads, currently 20+ settings. The exact rules and scoring formula are part of the open-source framework [KCCSS](https://github.com/octarinesec/kccss), the Kubernetes Common Configuration Scoring System. 

KCCSS is similar to the Common Vulnerability Scoring System (CVSS), the industry-standard for rating vulnerabilities, but instead focuses on the configurations and security settings themselves. Vulnerabilities are always detrimental, but configuration settings can be insecure, neutral, or critical for protection or remediation. KCCSS scores both risks and remediations as separate rules, and allows users to calculate a risk from 0 to 10, for every runtime setting of workload, and calculates the global risk of the workloads.

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

## Building from source code
Build the server image
```bash
cd server
docker build -t SERVER_TAG_NAME .
docker push SERVER_TAG_NAME
```

Build the client image
```bash
cd ../client
docker build -t CLIENT_TAG_NAME .
docker push CLIENT_TAG_NAME
```

go to root folder:
```bash
cd ../
```

Set kube-scan container image on kube-scan.yaml:
```bash
image: SERVER_TAG_NAME
```

Set kube-scan-ui container image on kube-scan.yaml:
```bash
image: CLIENT_TAG_NAME
```

Apply kube-scan.yaml:
```bash
kubectl apply -f kube-scan.yaml
kubectl port-forward --namespace kube-scan svc/kube-scan-ui 8080:80
```

# Screenshots

![Risk score](https://info.octarinesec.com/hubfs/home-1.png)

![Risk details](https://info.octarinesec.com/hubfs/risk-expanded.png)
