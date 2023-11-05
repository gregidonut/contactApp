FROM golang:latest AS build

RUN apt update
RUN apt install -y ca-certificates
RUN update-ca-certificates

WORKDIR /src

COPY ./cmd /src/cmd
COPY go.* /src

RUN CGO_ENABLED=0  go build -o /bin/contactApp ./cmd/web/main.go

FROM scratch

WORKDIR /bin

COPY --from=build /etc/ssl /etc/ssl
COPY --from=build /bin/contactApp /bin/contactApp
COPY ./ui /bin/ui

EXPOSE 8080

ENTRYPOINT ["/bin/contactApp"]
