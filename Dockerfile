FROM golang:1.21.5 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /intertask /app/cmd/blog

FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY --from=build-stage /intertask /intertask
EXPOSE 8080
ENTRYPOINT ["/intertask"]