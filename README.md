# Quotes API

API REST desenvolvida em Go utilizando o framework Gin para gerenciamento de citações. O projeto utiliza PostgreSQL como banco de dados, Google Wire para injeção de dependencias e segue uma arquitetura em camadas (handler, service, repository) para manter o codigo organizado e testavel.

## Comandos

| Comando | Descricao |
|---|---|
| `make dev` | Gera o Wire e inicia o servidor em modo desenvolvimento |
| `make build` | Gera o Wire e compila o binario em `bin/quotes-api` |
| `make wire` | Regenera o arquivo de injecao de dependencias (`di/wire_gen.go`) |
| `make test` | Executa todos os testes do projeto |
| `make clean` | Remove o diretorio `bin/` |
