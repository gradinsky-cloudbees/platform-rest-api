FROM golang:1.21.7-alpine3.19 AS build

WORKDIR /work

COPY go.mod* go.sum* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /usr/local/bin/rest main.go

FROM alpine:3.20

COPY --from=build /usr/local/bin/rest-api /usr/local/bin/rest-api

WORKDIR /cloudbees/home

ENTRYPOINT ["rest-api"]
