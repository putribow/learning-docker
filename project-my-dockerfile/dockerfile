FROM golang:1.21

WORKDIR /myapp

RUN go version

COPY AUTHORS.md ./

COPY LINKS.md ./

COPY go.mod go.sum ./

COPY . .

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /my-go-app

EXPOSE 77

CMD ["/my-go-app"]