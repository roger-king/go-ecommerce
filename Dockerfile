FROM golang:1.10-1.10-alpine

ARG APP_PATH=$GOPATH/src/github.com/rogr-king/go-ecommerce

ADD . ${APP_PATH}

WORKDIR ${APP_PATH}

RUN dep ensure

RUN go build -o build/app ${APP_PATH}/src

CMD [""]