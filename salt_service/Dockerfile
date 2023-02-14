FROM golang:1.19-alpine as builder
WORKDIR /salt_service
COPY . .
RUN go build -o salt_svc ./cmd/main.go

FROM alpine:3.16
WORKDIR /salt_service
COPY --from=builder /salt_service/salt_svc .
EXPOSE 9090
CMD ["/salt_service/salt_svc"]