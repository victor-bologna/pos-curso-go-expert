# Cobra CLI

Cobra CLI é uma ferramenta que permite que o desenvolvedor crie argumentos para rodar na sua aplicação durante o go run.

## Como começar o Cobra CLI

- Primeiro instalar o Cobra CLI em [Cobra CLI](https://github.com/spf13/cobra)
- Usar 'cobra-cli init'

## Como adicionar argumentos

- No Root: cobra-cli add nome
- No Argumento: cobra-cli add nome -p 'argumento-pai'

## Como executar comandos

- go run main.go argumento1 argumento2... --flag1=value1... -f1=value1...

## Comandos
- Flags persistentes são flags globais (aplicam aos argumentos filhos)
- Flags não persistentes são flags locais