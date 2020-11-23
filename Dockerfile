FROM golang as builder 
## build docker form scratch

COPY main.go .

RUN go build -o /main .

FROM alpine:3.7

CMD ["./main"]

COPY --from=builder /main .