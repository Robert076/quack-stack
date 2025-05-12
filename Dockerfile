FROM golang:1.24-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o duck-app ./cmd/quack-stack

FROM alpine:edge

WORKDIR /src

COPY --from=build /src/duck-app .

RUN chmod +x /src/duck-app

RUN apk --no-cache add ca-certificates

EXPOSE 8080

ENTRYPOINT [ "/src/duck-app" ]