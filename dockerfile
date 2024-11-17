FROM golang:latest AS pipeline
WORKDIR /go/src/pipeline
RUN go env -w GO111MODULE=auto
COPY . /go/src/pipeline
RUN go build -o pipeline .

FROM alpine:latest  
COPY --from=pipeline /go/src/pipeline/pipeline /build_pipe/
RUN chmod +x /build_pipe/pipeline
WORKDIR /build_pipe
ENTRYPOINT ["/build_pipe/pipeline"]






