FROM golang:latest as build

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GO111MODULE=on

WORKDIR $GOPATH/src/github.com/grandcolline/ecs_demo_batch
COPY . .
RUN go install


FROM gcr.io/distroless/base

COPY --from=build /go/bin/ecs_demo_batch /ecs_demo_batch
ENV PORT=8080
EXPOSE 8080

CMD ["/ecs_demo_batch"]
