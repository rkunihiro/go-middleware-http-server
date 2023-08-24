FROM golang:1.21.0-bullseye as builder
ENV CGO_ENABLED=0
COPY . /app
WORKDIR /app
RUN go build -ldflags "-s -w" -o main main.go

FROM gcr.io/distroless/static-debian11:nonroot
COPY --from=builder /app/main /main
CMD ["/main"]
