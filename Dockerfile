# Base image
FROM golang:1.17.2-alpine as builder
# Create directory
RUN mkdir /app
# Make it as working directory
WORKDIR /app 
# Copy application data into this directory
COPY ./app ./
# Add default environment variable
# This can be override
ENV SENDER_USERNAME="some email"
ENV SENDER_PASSWORD="some password"
ENV RECEIVER_MAIL="some password"
# Build go app
RUN CGO_ENABLED=0 GOOS=linux go build main.go

# multi steg Dockerfile
FROM scratch
# Copy from above build
COPY --from=builder /app ./go-app 
WORKDIR /go-app
# Starting go app
ENTRYPOINT [ "./main"]  