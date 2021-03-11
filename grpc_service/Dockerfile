FROM andrewfromgilead/dependency_go_image AS build_base

WORKDIR /go/src
COPY grpc /go/src/github.com/Andrew161644/full_app/grpc_service/grpc
WORKDIR /go/src/github.com/Andrew161644/full_app/grpc_service/grpc/main
RUN go get -d -v
RUN go build -o out .

FROM alpine:3.9
RUN apk add ca-certificates
COPY --from=build_base /go/src/github.com/Andrew161644/full_app/grpc_service/grpc/main/out /app/go-sample-app
EXPOSE 8081
CMD ["/app/go-sample-app"]