# syntax=docker/dockerfile:1

FROM golang:1.18beta1-bullseye

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /k8sjobs

EXPOSE 8080

CMD [ "/k8sjobs" ]