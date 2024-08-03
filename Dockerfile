FROM golang:1.22.3-alpine3.20 AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY pkg pkg
COPY main.go ./
RUN go build


FROM alpine:3.20

COPY --from=builder /usr/src/app/terraform-lock-checker /usr/local/bin/

WORKDIR /workspace
ENTRYPOINT ["/usr/local/bin/terraform-lock-checker"]
