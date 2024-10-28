# syntax=docker/dockerfile:1

# Tahap pembangunan
FROM golang:1.23.1 AS build

WORKDIR /app

# Salin dependensi dan unduh
COPY go.mod go.sum ./
RUN go mod download

# Salin kode sumber
COPY . .

# Bangun aplikasi
RUN go build -o main .

# Tahap produksi
FROM ubuntu:latest

WORKDIR /app

# Salin aplikasi dari tahap pembangunan
COPY --from=build /app/main .

# Eksekusi aplikasi
CMD ["./main"]