# Build environment
# -----------------
FROM golang:1.23-alpine AS build-env
WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/api ./cmd/api \
  && go build -ldflags '-w -s' -a -o ./bin/migrate ./cmd/migrate


# Deployment environment
# ----------------------
FROM alpine

COPY --from=build-env /app/bin/api /app/
COPY --from=build-env /app/bin/migrate /app/
COPY --from=build-env /app/migrations /app/migrations

EXPOSE 8080
CMD ["/app/api"]