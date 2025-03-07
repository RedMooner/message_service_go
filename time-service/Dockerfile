# Используем официальный образ Go
FROM golang:1.21

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем исходный код
COPY time-service /app

# Собираем приложение
RUN go build -o time-service .

# Открываем порт
EXPOSE 8082

# Запускаем приложение
CMD ["./time-service"]
