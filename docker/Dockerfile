FROM golang:1.23

# Set working directory
WORKDIR /server

# Copy Go modules and install dependencies
COPY ./server/go.mod ./server/go.sum ./
RUN go mod download

# Copy application code
COPY ./server ./

# Expose port
EXPOSE 8080

# Command to run the application
CMD ["go", "run", "."]

