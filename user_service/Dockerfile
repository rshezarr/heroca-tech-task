FROM golang:1.19-alpine as builder
WORKDIR /user_service
COPY . .
RUN go build -o user_svc ./cmd/main.go

FROM alpine:3.16
WORKDIR /user_service
COPY --from=builder /user_service/user_svc .
EXPOSE 9091
CMD ["/user_service/user_svc"]