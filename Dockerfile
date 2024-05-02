FROM golang:1.22.2-alpine AS builder

WORKDIR /builder

COPY . .

RUN go mod download
RUN go build -o backend cmd/api/main.go

FROM alpine

WORKDIR /backend

COPY --from=builder /builder/backend ./

EXPOSE 8080

CMD [ "./backend" ]
