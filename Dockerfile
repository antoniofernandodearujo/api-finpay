# Use uma imagem base do Go
FROM golang:1.22.2-alpine AS builder

# Defina o diretório de trabalho
WORKDIR /app

# Copie go.mod e go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copie o restante do código
COPY . .

# Compile a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o api-finpay .

# Use uma imagem menor para executar
FROM alpine:latest
RUN apk add --no-cache libc6-compat
WORKDIR /root/
COPY --from=builder /app/api-finpay .

# Exponha a porta que sua API usará
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./api-finpay"]
