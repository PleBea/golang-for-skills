FROM golang:alpine

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o main .

EXPOSE 8080

CMD ["./main"]
