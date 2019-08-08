FROM golang:1.11-alpine as builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a --installsuffix cgo -o helloserver *.go

FROM scratch
COPY --from=builder /build/helloserver .
EXPOSE 80
CMD ["./helloserver"]
