FROM golang:1.15 AS build

WORKDIR /home/srilanka
COPY . .
ARG GOPROXY="https://goproxy.cn"
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o srilanka .

EXPOSE 8081
CMD ["/bin/sh", "-c", "cd /home/srilanka/ && ./srilanka"]

