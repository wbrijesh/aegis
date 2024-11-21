FROM golang:1.23-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/main.go

FROM alpine:3.20.1 AS prod
WORKDIR /app

# Copy the .env file and set environment variables
COPY .env .env
ENV APP_ENV=${APP_ENV}
ENV PORT=${PORT}
ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}
ENV DB_DATABASE=${DB_DATABASE}
ENV DB_USERNAME=${DB_USERNAME}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_SCHEMA=${DB_SCHEMA}

COPY --from=build /app/main /app/main
EXPOSE ${PORT}
CMD ["./main"]
