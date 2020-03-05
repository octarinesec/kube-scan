FROM golang:1.13.5-alpine as build

WORKDIR /go/src
COPY /src /go/src

RUN go get
RUN go install
RUN go build

FROM alpine
COPY --from=build /go/src/kube-scan /
CMD ["/kube-scan"]
