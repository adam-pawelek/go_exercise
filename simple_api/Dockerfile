FROM golang:latest


WORKDIR /simple_api
COPY . .

RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init

EXPOSE 8080
RUN go build
RUN chmod 777 simple_api
CMD ["./simple_api"]


