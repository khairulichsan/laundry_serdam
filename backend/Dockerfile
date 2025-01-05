# Gunakan Golang image
FROM golang:1.23

# Set working directory
WORKDIR /app

# Copy semua file ke container
COPY . .

# Install dependencies
RUN go mod tidy

# Build aplikasi
RUN go build -o main .

EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]