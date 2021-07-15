FROM golang:1.16.5-buster AS build

# All these steps will be cached
RUN go version
ENV GOPATH=/
WORKDIR /src/

# Get dependancies - will also be cached if we won't change mod/sum
COPY ./go.mod /src/
COPY ./go.sum /src/
RUN go mod download

COPY ./ /src/
# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -a -v -o /app ./cmd/main.go

FROM alpine:latest

# copy go app, config and wait-for-postgres.sh
COPY --from=build /app /app
COPY ./configs/ /configs/
COPY ./.env /.env
COPY ./wait-for-postgres.sh ./

# install psql and make wait-for-postgres.sh executable
RUN apk add --no-cache libc6-compat postgresql-client && chmod +x wait-for-postgres.sh app

CMD ["./app"]