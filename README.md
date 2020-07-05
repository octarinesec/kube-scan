
<p align="center">
  <img src="./images/octarine_logo.png">
</p>

# Kube-Scan
Try our free Kubernetes risk assessment tool today.  
Run it on any cluster at any time. No data leaves your cluster. We do not collect any information.  
For more information on Octarine see https://www.octarinesec.com. 

# Get the risk score of your workloads

Kube-Scan gives a risk score, from 0 (no risk) to 10 (high risk) for each workload. The risk is based on the runtime configuration of each workload (currently 20+ settings). The exact rules and scoring formula are part of the open-source framework [KCCSS](https://github.com/octarinesec/kccss), the Kubernetes Common Configuration Scoring System. 

KCCSS is similar to the Common Vulnerability Scoring System (CVSS), the industry-standard for rating vulnerabilities, but instead focuses on the configurations and security settings themselves. Vulnerabilities are always detrimental, but configuration settings can be insecure, neutral, or critical for protection or remediation. KCCSS scores both risks and remediations as separate rules, and allows users to calculate a risk for every runtime setting of a workload and then to calculate the total risk of the workload.

**Please notice** that kube-scan currently scans the cluster when starting and will re-scan it every 24 hours. Thus, if you want to get an up-to-date risk score (e.g. after installing a new app), you should restart the kube-scan pod.

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

## Using the API


Getting all of the risks in your cluster:
```
GET http://HOST/api/risks
```

Requesting the kube-scan service to calculate again the risks:
```
POST http://HOST/api/refresh
```

This might be a long operation - depending on the cluster size, so you can pull the refresh operation status:
```
GET http://HOST/api/refreshing_status
```

## Building from source code
Build the server image (from root folder)
```bash
cd server
docker build -t SERVER_TAG_NAME .
docker push SERVER_TAG_NAME
```

Build the client image (from root folder)
```bash
cd client
docker build -t CLIENT_TAG_NAME .
docker push CLIENT_TAG_NAME
```

Set kube-scan containers images on the desired yaml (from root folder)
kube-scan container with SERVER_TAG_NAME
kube-scan-ui container with CLIENT_TAG_NAME

Apply the desired yaml and use "quick start" or "using load-balancer" instructions 

## Uninstall
```bash
kubectl delete -f https://raw.githubusercontent.com/octarinesec/kube-scan/master/kube-scan.yaml
```

In case of using a load-balancer:
```bash
kubectl delete -f https://raw.githubusercontent.com/octarinesec/kube-scan/master/kube-scan-lb.yaml
```

# Screenshots

![Risk score](https://info.octarinesec.com/hubfs/home-1.png)

![Risk details](https://info.octarinesec.com/hubfs/risk-expanded.png)
