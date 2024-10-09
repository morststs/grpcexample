FROM golang:1.23.1-alpine3.20 AS builder
RUN wget https://github.com/protocolbuffers/protobuf/releases/latest/download/protoc-28.2-linux-x86_64.zip \
    && unzip -q protoc-28.2-linux-x86_64.zip bin/protoc 'include/*' -d /usr/local \
    && chmod a+x /usr/local/bin/protoc \
    && rm -rf protoc-28.2-linux-x86_64.zip \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

FROM builder AS modules
COPY ./project/go.mod /go/src/go.mod
COPY ./project/proto /go/src/proto
WORKDIR /go/src
RUN cd proto \
    && protoc  --go_out=. --go-grpc_out=require_unimplemented_servers=false:. sample.proto \
    && cd .. 
RUN go mod tidy
COPY ./project/client /go/src/client
RUN cd client && go build -o client main.go && cd ..
COPY ./project/server /go/src/server
RUN cd server && go build -o server main.go && cd ..

FROM scratch AS client_stage
COPY --from=modules /go/src/client/client /

FROM scratch AS server_stage
COPY --from=modules /go/src/server/server /
