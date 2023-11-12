FROM golang:1.21.3-alpine AS build

# install make
RUN apk update
RUN apk add --no-cache make

WORKDIR /app

# restore
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# build
RUN make all

FROM alpine:latest

COPY --from=build /app/bin /app

WORKDIR /app

CMD ["./init.sh"]
