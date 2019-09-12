FROM golang:1.12-stretch AS build
COPY main.go .
RUN go get github.com/prometheus/client_golang/prometheus
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /main main.go

FROM scratch
COPY --from=build /main /main

CMD ["/main"]
