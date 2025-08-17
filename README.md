# Order Service - gRPC + Go

Este projeto implementa um microsserviço de **ordens de compra** utilizando **Go** e **gRPC**.  
Ele faz parte de uma arquitetura de microsserviços, onde o serviço de **Order** se comunica com o serviço de **Payment** via gRPC.

---

## 🚀 Funcionalidades

- Criar pedidos com múltiplos itens  
- Validação de quantidade máxima permitida (até 50 itens por pedido)  
- Integração com serviço de pagamento via gRPC  
- Cancelamento automático de pedidos em caso de falha no pagamento  
- Mecanismo de **retry automático** para chamadas gRPC com backoff linear  

---

## 🛠️ Tecnologias

- Go  
- gRPC  
- Protocol Buffers  
- GORM  
- go-grpc-middleware/retry  

---

## 📂 Estrutura do Projeto

application/     # Implementação dos casos de uso (ex: PlaceOrder)  
domain/          # Entidades e regras de negócio  
infrastructure/  # Conexões externas (DB, Payment service)  
proto/           # Definições .proto do gRPC  
server/          # Inicialização do servidor gRPC  
commands.txt     # Lista de comandos úteis para rodar/testar  
main.go          # Ponto de entrada da aplicação  

---

## ⚙️ Como rodar os serviços

### 1. Clonar o repositório
git clone https://github.com/LucasJaud/microservices.git

cd microservices

### 2. Instalar dependências
go mod tidy

### 3. Criar o Banco de Dados
Na pasta raiz executar o comando: `docker run -p 3306:3306 -e MYSQL_ROOT_PASSWORD=minhasenha -v "$(pwd)/init.sql:/docker-entrypoint-initdb.d/init.sql" mysql`

Esse comando deve criar um container mysql escutando a porta 3306

caso queira verificar o container executar o comando: `docker exec -it [container] mysql -uroot -p` e digirar a senha `minhasenha`

### 3. Rodar o serviço de Payment
O Order depende do Payment para processar os pagamentos.  
Entre na pasta do serviço de pagamento e execute:

`DB_DRIVER=mysql DATA_SOURCE_URL=root:minhasenha@tcp(127.0.0.1:3306)/payment APPLICATION_PORT=3001 ENV=development go run cmd/main.go`


### 4. Rodar o serviço de Order
Em outro terminal, va para a pasta do service de order e execute:

`DATA_SOURCE_URL=root:minhasenha@tcp(127.0.0.1:3306)/order APPLICATION_PORT=3000 ENV=development PAYMENT_SERVICE_URL=localhost:3001 go run cmd/main.go`


### 5. Testar os serviços
No arquivo commands.txt você encontra exemplos prontos de comandos (grpcurl) para criar pedidos e simular falhas de pagamento.

---

## 🔄 Retry do gRPC

O cliente de pagamento foi configurado para realizar até 5 tentativas em caso de erro codes.Unavailable ou codes.ResourceExhausted, utilizando backoff linear de 1 segundo entre tentativas.

---

## 👨‍💻 Autor

Projeto desenvolvido por Lucas Jaud.  
