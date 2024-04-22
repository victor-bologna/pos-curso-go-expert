# Dependency Injection

Dependency Injection é uma forma de abstrair configurações e criação de várias instâncias de dados a fim de simplificar o processo de criação. Esse processo usa um container de injeção de dependência para facilitar a criação de uma instância de dado. Exemplo:

## Sem Depedency Injection:

- Criar uma conexão com DB
- Criar um Repositório com essa conexão DB;
- Criar um Use Case com o Repositório com a conexão DB;

## Com Dependency Injection:

- Criar um container responsável por gerar um use Case e abstrair toda a implementação com Banco de Dados

## Inversão de Controle:

Inversão de controle é implementar uma interface em um metódo a fim de desacoplar conexões concretas. Ou seja, um metódo depender de um objeto abstrato (interface) ao invés do objeto em si. 

## Dependency Injection em Go:

### Uber FX

[Uber FX](https://github.com/uber-go/fx) é uma biblioteca de injeção de dependência usando [Reflections](https://go.dev/blog/laws-of-reflection). Reflections tem um problema no qual a tipagem da variável é identificada em tempo de execução, e por isso pode gerar um erro na execução do programa.

### Google Wire

[Google Wire](https://github.com/google/wire) é uma biblioteca que gera código do tipo Go sem uso de Reflections, ou seja, todos os tipos são definidos e gerados no usso da biblioteca e não na execução.

#### Como usar Google Wire

Criar uma classe wire.go com pacote main e colocar as dependências necessárias do projeto com as seguintes anotações:

//go:build wireinject
// +build wireinject

Para rodar o projeto com Google Wire precisa dar: "go run main.go wire_gen.go"