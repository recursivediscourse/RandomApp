FROM alpine

COPY . /app
CMD ["/app/scripts/demo.sh"]
