FROM golang:latest
RUN mkdir -p /app
COPY . /app
WORKDIR /app
EXPOSE 8080
RUN go mod tidy
RUN go build -o /main
CMD [ "/main" ]