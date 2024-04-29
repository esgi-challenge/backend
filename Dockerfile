FROM golang:1.22.2-alpine AS builder

WORKDIR /builder

COPY . .

RUN go build cmd/api/main.go

FROM alpine

WORKDIR /backend

COPY --from=builder /builder/main ./

EXPOSE 8080

CMD [ "./main" ]
