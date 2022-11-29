## Build
FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go .
COPY ./user/*.go ./user/
COPY ./failure/*.go ./failure/

RUN go build -o /blobling

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /blobling /blobling

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/blobling"]