# Start from the latest golang base image
FROM golang:1.21 as builder

RUN apt-get update && apt-get install -y ca-certificates git-core openssh-client

# Add Maintainer Info
LABEL maintainer="User Task"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# # Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# # Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod vendor

# # Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

######## Start a new stage from scratch #######
FROM alpine:3 AS runner

RUN apk --no-cache add ca-certificates
RUN apk add --no-cache tzdata
ENV TZ=Asia/Jakarta

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy files setting project to the Workdir
COPY --from=builder /app/config.toml .

# Expose port 3000 to the outside world
EXPOSE 8081

# Command to run the executable
CMD ["./main"]