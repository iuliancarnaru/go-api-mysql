FROM golang:1.22 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /api

FROM build-stage AS run-test-stage

RUN go test -v ./..

FROM scratch AS run-release-stage

WORKDIR /app

COPY --from=build-stage /api /api

EXPOSE 4000

CMD ["/api"]