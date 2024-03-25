FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build ./app/server 

RUN go install github.com/pressly/goose/cmd/goose@latest

EXPOSE 50051

CMD ["sh", "-c", "goose -dir sql/migrations up && ./server"]