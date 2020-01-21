#### Step 1: Build the executable
FROM ubuntu:18.04 as build

# for the build process
RUN apt-get update && apt-get install -y wget

# Download & Install GO
ENV GOLANG_FILE_NAME=go1.13.5.linux-amd64.tar.gz
RUN wget --quiet https://dl.google.com/go/${GOLANG_FILE_NAME}
RUN tar -C /usr/local -xzf ${GOLANG_FILE_NAME}
RUN ln -sf /usr/local/go/bin/go /usr/local/bin/go
RUN ln -sf /usr/local/go/bin/godoc /usr/local/bin/godoc
RUN ln -sf /usr/local/go/bin/gofmt /usr/local/bin/gofmt

# Copy source code and complie
ADD src /
RUN go build


#### Step 2: Copy & Run the executable
FROM ubuntu:18.04
COPY --from=build /kube-scan /
CMD ["/kube-scan"]
