FROM golang:1.22

WORKDIR /app

COPY . .

RUN go mod init github.com/brunoronte/fullcycle-ci

RUN go build -o math

CMD ["./math"]