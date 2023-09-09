# since the image size is 550Mb, we need to move to multi-stage builds

# Build stage
FROM golang:1.20.5-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go    
 
# Run stage 
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main . 
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY auth-service/db/migration ./db/migration

EXPOSE 8100
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]