# Octaudit
Octarine k8s cluster risk assesment tool  
https://www.octarinesec.com/

## Quickstart
```bash
kubectl apply -f https://raw.githubusercontent.com/octarinesec/octaudit/master/octaudit.yaml
kubectl port-forward --namespace octaudit svc/octaudit-ui 8080:80
```

Then set your browser to `http://localhost:8080`.

## Using a load-balancer service
```bash
kubectl apply -f https://raw.githubusercontent.com/octarinesec/octaudit/master/octaudit_lb.yaml
kubectl -n octaudit get service octaudit-ui -o jsonpath={..hostname}
```

Then set your browser to that hostname.
