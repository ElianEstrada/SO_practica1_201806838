FROM golang:1.17.6
WORKDIR /home/server
COPY . .
RUN go download
CMD ["go", "run", "main.go"]