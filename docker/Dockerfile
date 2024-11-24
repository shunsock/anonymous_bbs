FROM golang:1.20

# Set working directory
WORKDIR /server

# Copy Go modules and install dependencies
COPY ./server/go.mod ./server/go.sum ./
RUN go mod download

# Copy application code
COPY ./server ./

# Expose port and set default command
EXPOSE 8080
CMD ["go", "run", "main.go"]

