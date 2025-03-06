# Используем официальный образ Go
FROM golang:1.21

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем исходный код
COPY api-gateway /app

# Собираем приложение
RUN go build -o api-gateway .

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./api-gateway"]
