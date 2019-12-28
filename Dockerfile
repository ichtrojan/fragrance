FROM golang:latest

WORKDIR /

COPY . .

CMD ["go", "run", "main.go"]