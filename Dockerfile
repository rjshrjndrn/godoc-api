FROM golang:1.22-alpine AS builder
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /tmp/godoc .

FROM alpine
COPY --from=builder /tmp/godoc /usr/bin/godoc
EXPOSE 8000
CMD ["godoc"]

