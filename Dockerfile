FROM alpine:latest

WORKDIR /root

COPY main /root/main

CMD ["./main"]