FROM golang:1.14 as builder

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /executable


FROM scratch

COPY --from=builder /executable /
CMD ["/executable"]
EXPOSE 8080