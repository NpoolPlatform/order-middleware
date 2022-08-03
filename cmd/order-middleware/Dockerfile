FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

WORKDIR /app

COPY order-middleware /app

ENTRYPOINT ["/app/order-middleware"]

EXPOSE 50440 50441 50442 50443

CMD ["run"]
