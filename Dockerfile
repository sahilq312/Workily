FROM golang:1.23.2-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o workly main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/workly .
EXPOSE 8080
ENV PORT=8080
CMD ["./workly"]
