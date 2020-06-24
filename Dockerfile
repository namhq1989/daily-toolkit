# Start from the latest golang base image
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="Nam Huynh <namhq.1989@gmail.com>"

# Set the Current Working Directory inside the container
RUN mkdir -p /go/src/github.com/namhq1989/daily-toolkit
WORKDIR /go/src/github.com/namhq1989/daily-toolkit

# Copy data to working dir
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./main.go

######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add tzdata zip ca-certificates

WORKDIR /src/github.com/namhq1989/daily-toolkit

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/src/github.com/namhq1989/daily-toolkit .

# Command to run the executable
CMD ["./main"]