FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o ecom .

EXPOSE 8080

CMD [ "./ecom", "-env", "container" ]