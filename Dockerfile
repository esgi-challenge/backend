FROM golang:1.22.2-alpine AS builder

WORKDIR /builder

COPY . .

RUN go mod download
RUN go build -o backend cmd/api/main.go

FROM alpine

# curl not installed in alpine by default so wget is used
HEALTHCHECK --interval=2m --timeout=3s --retries=3 CMD wget --no-verbose --tries=1 --spider http://127.0.0.1:8080/api/healthz || exit 1

WORKDIR /backend

COPY templates templates
COPY --from=builder /builder/backend ./

EXPOSE 8080

CMD [ "./backend" ]
