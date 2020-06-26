FROM golang:1.14-alpine3.12 as builder

WORKDIR /src

COPY go.mod .
RUN go mod download

COPY . .

RUN go build -o internal-namespace-expoter

FROM alpine:3.12 as release

COPY --from=builder /src/internal-namespace-expoter /usr/local/bin/internal-namespace-expoter

EXPOSE 1192

ENTRYPOINT /usr/local/bin/internal-namespace-expoter
