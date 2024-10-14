# Dockerfile

FROM golang:1.23.2-alpine AS builder

WORKDIR /go/src/app
COPY . .

# Sao chép wait-for-it.sh vào thư mục làm việc
COPY wait-for-it.sh ./

RUN chmod +x wait-for-it.sh

# Copy the Go modules files first để tận dụng bộ nhớ đệm Docker layer
COPY go.mod go.sum ./

# Download và verify dependencies trước khi copy source code (tận dụng cache của Docker)
RUN go mod download
RUN go mod verify


# Biên dịch ứng dụng
RUN go build -o main .


# Bước tiếp theo cho image chính
FROM alpine:latest

WORKDIR /app

# Sao chép file thực thi từ bước trước
COPY --from=builder /go/src/app/main .
COPY --from=builder /go/src/app/app.env .
COPY --from=builder /go/src/app/wait-for-it.sh .


EXPOSE 8080
CMD ["./wait-for-it.sh", "postgres:5432", "--", "./main"]
