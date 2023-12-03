FROM golang:1.21

WORKDIR /main_terminal
COPY ["./go.mod", "./go.sum", "./"]
RUN go mod download

COPY ./ ./
RUN go build -o ./bin/app cmd/app/main.go

EXPOSE 8000

CMD ["./bin/app"]