FROM golang:1.21.5-alpine3.19

WORKDIR /app

COPY . /app

# build the go app with the name main
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]