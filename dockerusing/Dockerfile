FROM golang:alpine

RUN mkdir /files
COPY . /files
WORKDIR /files
RUN go build -o /files/hw hw.go
ENTRYPOINT [ "ash", "/files/hw" ]
