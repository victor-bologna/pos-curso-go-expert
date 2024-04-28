# Extras

Modulos extras relacionados ao conteúdo Go

## Graceful Shutdown

Impede novas requisições e finaliza as atuais antes de fechar o servidor.

## Panic Recover

Quando acontece um erro na aplicação (panic) tomamos alguma ação (catch/finally do Java)

## FSNotify

[FSNotify](https://github.com/fsnotify/fsnotify) fica lendo variáveis de ambiente e se for alterado ele recarrega a variável (Ex. senha do banco de dados em PRD mudando a cada x tempo).

## Fast JSON

[FastJson](https://github.com/valyala/fastjson) é uma API complementar e mais rápida que o enconding/json do Go.

## Webserver Go

Webserver Go é uma melhoria na versão 1.22 que possibilita passar parâmetros na URL, fazer URL exatas, passar diretórios na URL, passar qualquer valor no endpoint atraves de {} e fazer outras requisições HTTP como Post Put Patch e Delete.