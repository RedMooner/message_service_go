# Используем официальный образ Go
FROM golang:1.21

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем исходный код
COPY message-service /app

# Собираем приложение
RUN go build -o message-service .

# Открываем порт
EXPOSE 8081

# Запускаем приложение
CMD ["./message-service"]
