FROM golang:alpine
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /mal
CMD ["/mal"]