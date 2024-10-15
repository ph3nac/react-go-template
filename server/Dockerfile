FROM golang:1.23.1
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
RUN go install github.com/air-verse/air@latest
CMD ["air"]
