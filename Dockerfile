# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

# Dependencies
COPY src/ ./
RUN go mod download

# Build
RUN go build -o /selene

# Run
CMD [ "/selene" ]
