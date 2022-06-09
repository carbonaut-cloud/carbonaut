##
## Build
##
FROM golang:1.18-buster AS build

WORKDIR /app

COPY . ./
RUN go mod download
RUN make install

RUN go build -o /carbonaut cmd/main.go

##
## Deploy
##
FROM gcr.io/distroless/base

WORKDIR /

COPY --from=build /carbonaut /carbonaut

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["/carbonaut"]
