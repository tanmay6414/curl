# Base image
FROM golang:1.17.2-alpine as builder
# Create directory
RUN mkdir /app
# Make it as working directory
WORKDIR /app 
# Copy application data into this directory
COPY ./app ./
EXPOSE 587
RUN CGO_ENABLED=0 GOOS=linux go build main.go

FROM busybox  
COPY --from=builder /app ./go-app 
WORKDIR /go-app
ENTRYPOINT ["/bin/sh", "-c", "./main"]  