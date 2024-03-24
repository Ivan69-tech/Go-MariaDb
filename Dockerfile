FROM golang:latest

WORKDIR /app


COPY . .


RUN go build -o godata .


CMD ["./godata"]
