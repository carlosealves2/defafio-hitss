FROM golang:1.20 as builder
LABEL authors="carlos"

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o application src/cmd/main.go

FROM gcr.io/distroless/static-debian11

COPY --from=builder /go/src/app/application /usr/local/bin/

EXPOSE 8000

USER nonroot:nonroot

CMD ["/usr/local/bin/application"]