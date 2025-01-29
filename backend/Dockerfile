# Usar a imagem oficial do Golang
FROM golang:latest

# Definir o diretório de trabalho dentro do container
WORKDIR /app

# Definir o proxy Go para garantir a instalação de dependências
#ENV GOPROXY=https://proxy.golang.org,direct

# Copiar arquivos de dependência
COPY go.mod go.sum ./
RUN go mod download

# Instalando Delve
#RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Copiar arquivos de configuração (como .env)
COPY .env ./

# Copiar o código-fonte para o contêiner
COPY . .

# Compilar o binário da aplicação
RUN go build -o /app/main ./cmd/main.go

# Listar o conteúdo para verificar a existência do binário
RUN ls -al /app

# Expor a porta da aplicação
EXPOSE 8080

# Executar o binário
CMD ["/app/main"]
