FROM golang:1.22.5 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/auto_proforma .

FROM alpine:latest
WORKDIR /root/
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/auto_proforma .
EXPOSE 8080
CMD ["./auto_proforma"]