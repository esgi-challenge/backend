FROM golang:alpine AS builder

WORKDIR /builder

COPY . .

RUN go build -o backend

FROM alpine

WORKDIR /backend

COPY --from=builder /builder/backend ./

CMD [ "./backend" ]
