FROM golang:1.23.1 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./web-app 

FROM scratch

COPY --from=builder /app/web-app /web-app
COPY /templates /templates
COPY /posts /posts

EXPOSE 8080

ENTRYPOINT [ "/web-app" ]