FROM golang:1.12-alpine AS build
COPY main.go .
RUN CGO_ENABLED=0 go build -o /main main.go

FROM scratch
COPY --from=build /main /main
CMD ["/main"]
