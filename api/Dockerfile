FROM golang:1.18.2-alpine AS build
ADD . /src
WORKDIR /src
RUN GOOS=linux go build -tags netgo -installsuffix netgo -o server cmd/server/main.go

FROM scratch
COPY --from=build /src/server /server
ENTRYPOINT ["/server"]
