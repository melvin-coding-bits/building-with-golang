FROM golang:1.13-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /user-service

EXPOSE 3000

CMD [ "/user-service" ]