FROM --platform=$BUILDPLATFORM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go .

ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o main .

#FROM gcr.io/distroless/static:nonroot
FROM alpine:latest
WORKDIR /
COPY --from=builder /app/main .
EXPOSE 8080

CMD ["./main"]
