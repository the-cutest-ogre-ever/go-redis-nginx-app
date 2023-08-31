FROM golang:latest

RUN apt-get update && apt-get install -y golang

WORKDIR /app

COPY . /app

CMD ["go", "run", "./main"]