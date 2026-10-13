# Etap 1: budowanie
FROM golang:1.22-alpine AS build
WORKDIR /src
COPY app/go.mod ./
COPY app/ ./
RUN go test ./... && CGO_ENABLED=0 go build -ldflags="-s -w" -o /out/app .

# Etap 2: minimalny obraz wykonawczy
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/app /app
ARG APP_VERSION=dev
ENV APP_VERSION=${APP_VERSION} PORT=8080
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app"]
