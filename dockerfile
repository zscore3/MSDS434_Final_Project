FROM golang:1.14

LABEL base.name "golangtest"

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 8080

ENTRYPOINT ["./main"]
