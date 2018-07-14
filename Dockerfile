FROM golang:1.10.3 as builder

ARG APP_PATH=$GOPATH/src/github.com/rogr-king/go-ecommerce

ADD . $APP_PATH

WORKDIR $APP_PATH

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN dep ensure

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/app .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/github.com/rogr-king/go-ecommerce/build/app .

CMD ["./app"]