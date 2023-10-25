FROM golang:latest AS build

WORKDIR /src

COPY ./cmd /src/cmd
COPY go.* /src

RUN CGO_ENABLED=0  go build -o /bin/contactApp ./cmd/web/main.go

FROM scratch

WORKDIR /bin

COPY --from=build /bin/contactApp /bin/contactApp
COPY ./testingAssets /bin/testingAssets
COPY ./ui /bin/ui

EXPOSE 8080

ENTRYPOINT ["/bin/contactApp"]
