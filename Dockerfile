# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.17-alpine AS build

WORKDIR /app

COPY src/ ./
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /selene

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /selene /selene

EXPOSE 8000

USER nonroot:nonroot

ENTRYPOINT ["/selene"]
