FROM golang:1.16.5 as builder

WORKDIR /app
COPY . /app
RUN make build-static

FROM scratch

COPY --from=builder /app/bin/foods-static ./foods
COPY --from=builder /app/var ./var

EXPOSE 8080

CMD ["/foods"]
