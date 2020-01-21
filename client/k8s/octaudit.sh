#! /bin/sh

YAML=octaudit.yaml
NAMESPACE=octaudit
PORT=8080
DEPLOYMENT=deployment/octaudit
SERVICE=service/octaudit-ui

cleanup()
{
#  kubectl delete -f ${YAML}
#  kubectl -n ${NAMESPACE} wait --for=condition=deleted ${DEPLOYMENT}
  exit 2
}

trap "cleanup" 2

kubectl apply -f ${YAML}
kubectl -n ${NAMESPACE} wait --for=condition=Ready ${DEPLOYMENT}
kubectl -n ${NAMESPACE} port-forward ${SERVICE} ${PORT}:80