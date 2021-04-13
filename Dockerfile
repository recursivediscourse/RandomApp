FROM golang:1.16-alpine

COPY . /app
WORKDIR /app

CMD ["/app/scripts/demo.sh"]
