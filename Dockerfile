FROM golang:1.18.2-alpine3.16
RUN mkdir /app
ADD . /app 
WORKDIR /app
RUN go mod init main 
RUN go build -o main .
CMD ["/app/main"]

