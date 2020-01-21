#### Step 1: Build the executable
FROM ubuntu:18.04 as build

ENV LANG C.UTF-8
ENV LC_ALL C.UTF-8
ENV DEBIAN_FRONTEND noninteractive

# for the build process
RUN apt-get update && apt-get install -y build-essential \
    wget

# Download & Install GO
ENV GOLANG_FILE_NAME=go1.13.5.linux-amd64.tar.gz
RUN wget --quiet https://dl.google.com/go/${GOLANG_FILE_NAME}
RUN tar -C /usr/local -xzf ${GOLANG_FILE_NAME}
RUN ln -sf /usr/local/go/bin/go /usr/local/bin/go
RUN ln -sf /usr/local/go/bin/godoc /usr/local/bin/godoc
RUN ln -sf /usr/local/go/bin/gofmt /usr/local/bin/gofmt
RUN rm ${GOLANG_FILE_NAME}

# Copy source code and complie
ADD src /
RUN pwd
RUN cd /
RUN pwd
RUN ls
RUN go build
RUN ls


#### Step 2: Copy & Run the executable
FROM ubuntu:18.04
COPY --from=build /kube-scan /
RUN ls
CMD ["/kube-scan"]
