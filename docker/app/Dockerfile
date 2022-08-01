FROM golang:1.18-alpine

WORKDIR /go/src
COPY . .

RUN apk upgrade --update && \
    apk --no-cache add git

CMD ["go", "run", "main.go"]