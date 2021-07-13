FROM golang:alpine as build

WORKDIR /go/src
COPY /src /go/src

RUN go get
RUN go install
RUN go build

FROM alpine:latest
COPY --from=build /go/src/kube-scan /
CMD ["/kube-scan"]


# # kube-scan/state_reader
# state_reader/cluster_state_reader.go:85:65: not enough arguments in call to reader.kubeClient.CoreV1().Namespaces().List
#         have ("k8s.io/apimachinery/pkg/apis/meta/v1".ListOptions)
#         want (context.Context, "k8s.io/apimachinery/pkg/apis/meta/v1".ListOptions)
# state_reader/cluster_state_reader.go:102:83: not enough arguments in call to reader.kubeClient.RbacV1().ClusterRoleBindings().List
#         have ("k8s.io/apimachinery/pkg/apis/meta/v1".ListOptions)
#         want (context.Context, "k8s.io/apimachinery/pkg/apis/meta/v1".ListOptions)
# state_reader/cluster_state_reader.go:117:81: not enough arguments in call to reader.kubeClient.AppsV1().Deployments(namespace.ObjectMeta.Name).List
#         have ("k8s.io/apimachinery/pkg/apis/meta/v1".ListOptions)
#         want (context.Context, "k8s.io/apimachinery/pkg/apis/meta/v1".ListOptions)
# state_reader/cluster_state_reader.go:126:81: not enough arguments in call to reader.kubeClient.AppsV1().ReplicaSets(namespace.ObjectMeta.Name).List
#         have ("k8s.io/apimachinery/pkg/apis/meta/v1".ListOptions)
