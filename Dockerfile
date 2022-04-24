FROM golang:1.16

WORKDIR /var/www/html

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

ENV MYSQL_DSN="user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local" \
    REDIS_ADDR="redis:6379" \
    REDIS_PW="" \
    REDIS_DB="" \
    GIN_MODE="debug" \
    PORT=8888

COPY . /var/www/html

RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go build -o api_server

RUN chmod +x ./api_server

CMD ["./api_server"]