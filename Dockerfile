# Use a imagem base do Go
FROM golang:latest

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie o código fonte para o diretório de trabalho no contêiner
COPY . /app

# Compile o código Go
RUN go build -o vps-event

# Exponha a porta 8080 para acessar o servidor HTTP
EXPOSE 8080

# Comando padrão a ser executado quando o contêiner for iniciado
CMD ["./vps-event"]