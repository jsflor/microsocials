FROM golang:1.20-alpine

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

ARG PORT=8080
ARG DATABASE_PORT=5432
ARG DATABASE_USER=postgres
ARG DATABASE_PASSWORD=postgres
ARG DATABASE_NAME=postgres

ENV PORT=${PORT}
ENV DATABASE_PORT=${DATABASE_PORT}
ENV DATABASE_USER=${DATABASE_USER}
ENV DATABASE_PASSWORD=${DATABASE_PASSWORD}
ENV DATABASE_NAME=${DATABASE_NAME}

RUN go build -o /messaging-service

CMD [ "/messaging-service" ]