## Pós Graduação Go Expert

Repositório referente a todo o desenvolvimento e desafios feitos durante a realização do curso. Todos os modulos estão separados e classificados por categorias.

## GraphQL

Serve para definir dados necessários para comunicação entre servidores RPC (REST, gRPC).
[gqlgen](https://gqlgen.com/getting-started/)

## gRPC

gRPC é uma comunicação entre cliente servidor usando HTTP/2 permitindo se comunicar via stream (recebendo/enviando) bytes de dados de forma mais rápida e leve.

## Go migrate

SQL migrate é um processo de realizar mudanças no schema do banco por versões com finalidade de dar rollback.
[go-migrate](https://github.com/golang-migrate/migrate)

### Comandos Go Migrate

- Subir Dados -> migrate -path=sql/migrations/ -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up
- Excluir Dados -> migrate -path=sql/migrations/ -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down