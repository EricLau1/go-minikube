FROM golang

ENV GO111MODULE=on

WORKDIR /api

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build main.go

EXPOSE 9000

ENTRYPOINT [ "/api/main" ]