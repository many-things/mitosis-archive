FROM golang:1.20-alpine AS deps

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

FROM deps AS build

COPY . .

RUN go build -o build/mitosisd ./cmd/mitosisd

FROM alpine:latest AS runner

COPY --from=build /app/build/mitosisd /usr/bin/

ENTRYPOINT ["/usr/bin/mitosisd"]
