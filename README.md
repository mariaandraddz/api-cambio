# API de Conversão de Moedas

API REST simples em Go para conversão de moedas em tempo real, utilizando cotações da AwesomeAPI.

----

### Funcionalidades
- Conversão entre duas moedas usando cotação atualizada
- Suporte a qualquer par de moedas disponível na AwesomeAPI
- Retorno da taxa de câmbio utilizada no cálculo
- Tratamento de erros para pares inválidos, JSON malformado e falhas no provedor externo

### Tecnologias
- Go
- Gin — framework web
- AwesomeAPI — provedor de cotações

### Pré-requisitos
- Go 1.20+ instalado
- Conexão com a internet (a API consulta cotações externas em tempo real)

### Instalação
```
# Clone o repositório
git clone <url-do-seu-repositorio>
cd <nome-do-projeto>

# Baixe as dependências
go mod download
```
### Executando o projeto
```
go run main.go
```
O servidor sobe por padrão na porta 8080:
```
http://localhost:8080
```
### Testando a API
O projeto inclui um arquivo test.http com exemplos de requisições prontos para usar com a extensão REST Client (VS Code) ou diretamente no GoLand/IntelliJ.

### Estrutura do projeto
```
├── main.go        # Lógica principal da API
├── go.mod
├── go.sum
├── test.http       # Requisições de exemplo para testes
└── README.md
```

Feito com carinho por <a href="https://www.linkedin.com/in/mariaeduardaandradee/">Maria Eduarda</a> 🐈
