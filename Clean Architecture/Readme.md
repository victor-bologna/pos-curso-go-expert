# Clean Architecture

- Objetivo princípal: Separar domínio da aplicação
- Criado por Uncle Bob
- Baixo acoplamento entre camadas

- Gerar arquivos proto na pasta grpc: protoc --go_out=. --go-grpc_out=. proto/course_category.proto
- Para regenerar arquivos GraphQL: go run github.com/99designs/gqlgen generate
- Rodar o servidor na pasta server: go run main.go wire_gen.go
- Rodar o Evans (GRPC): evans -r repl

## Pontos importantes sobre Arquitetura

- Formatar o software (Formato do software)
- Dividir claramente os componentes (Não tem metade de um andar em um prédio). Componentes tem lugares específicos para ficar
- Comunicação entre componentes
- Arquitetura ajuda a desenvolver, facilitando o deploy, operação e manutenção. (Organizar a estrutura do código)
- “The strategy behind that facilitation is to leave as many options open as possible, for as long as possible.”
- Flexibilidade de código

## Objetivos de uma boa arquitetura

O objetivo principal da arquitetura é dar suporte ao ciclo de vida do sistema. Uma boa arquitetura torna o sistema fácil de entender, fácil de desenvolver, fácil de manter e fácil de implantar. O objetivo final é minimizar o custo de vida útil do sistema e maximizar a produtividade do programador.

### Beneficios de uma boa arquitetura

- A vida do software fica mais longa
- Upgrade de tecnologia
- Custa mais para implementar coisas em um software mal arquiteturado.
- Keep options open (Mantenha opções abertas)

### Regras vs Detalhes

- Sempre começar a trabalhar com as regras de software primeiro
- Regras de negócio trazem valor real para o software
- Pior coisa é perder tempo com detalhe (Configuração do RabbitMQ por exemplo)
- Detalhes ajudam a suportar as regras
- Detalhes não devem impactar nas regras de negócio. (Se o detalhe começa a impactar nas regras de negócio, então não há delimitação de camada entre regra de negócio e detalhes.)
- Frameworks, banco de dados, tipos de APIs não devem impactar na regra de negócio.
- Detalhes podem ser substituidos (RabbitMQ por SQS; REST por gRPC; Postgres por MongoDB)
- DDD -> Atacar a complexicidade no coração do software (regras de negócios)

## Use Cases

- Intenção do software (Cada ação é uma intenção, cada intenção é um caso de uso)
- Clareza de cada comportamento do software tem
- Detalhes (Frameworks, bancos, APIs) não devem impactar a regra de negócio
- DRY (Don't repeat yourself)
- Use Cases contam uma história
- Software é uma forma de automatizar tarefas do dia a dia. Use Cases é uma concretização de um Software.
- Use Case é o orquestrador do fluxo da regra de negócio. Que gera o fluxo/automação do software.
- Entidades -> Camada onde guarda as regras de negócio.
- Regras =/= Fluxo do caso de uso é outra.
- UseCase em DDD seria uma camada de aplicação, pois ela diz como a aplicação deve trabalhar.
Exemplo: 
    - Pegar informação para gerar um empréstimo:
     1. Receber informações do cliente (nome, data de nascimento, endereço...);
     2. Validar nome e endereço, data de nascimento, etc...;
     3. Pegar histórico de crédito;
     4. Se crédito < 500 então nega, se nào crie um cliente e ative um orçamento de empréstimo;
     5. Retorna as informações do cliente + o orçamento.

### Use Cases e SRP (Single Responsability Principle)

- Use Cases não devem ser reaproveitados (Inserir e Alterar) não podem ser reutilizados
- Cada use case tem que ser independente de um do outro
- SRP -> Mudar por razões diferentes
- Duplicação acidental -> Repetir um trecho de código várias vezes (exemplo: uma validação)
- Duplicação real -> Código pode parecer repetido, porém tem ações diferentes no final (Ex: Inserir e alterar código)

## Limites Arquiteturais

- Tudo que não impacta diretamente nas regras de negócio deve estar em um limite arquitetural diferente. Ex: não será Frontend ou banco de dados que mudarão regras de negócio da aplicação.
- Dependa de abstrações e não de implementações concretas
- Cada camada tem que ser responsável pela sua função somente, evitando acoplamento de camadas.

## Input vs Output

- Input que retorna um Output.
- Simplificar raciocínio ao criar um software pensando em Input e Ouput.
- Input chama o controller que chama o Use case e por fim as entidades.
- Transformar output num formato ideal (para CMD, REST, Txt...) -> Presenter

## DTOs (Data transfer Object)

- Trafega dados entre limites arquiteturais (Como se fosse um envelope)
- Objeto anêmico, sem comportamento e só possui dados.
- Contém Dados de Input e Output (InputDTO e OutputDTO).
- Normalmente cada Use Case possui Input e Output separado.

## Presenters

- Objetos de transformação (Serializar os dados)
- Adequa o DTO do Output num formato apresentável e que esteja sendo exigido.
- Content-Type do REST é uma das opções de Presenters (application/json).
- Geralmente usado após o UseCase 
Ex.: 
    - jsonResult = CategoryPresenter(outputDTO).toJson();
    - xmlResult = CategoryPresenter(outputDTO).toXML();

## Entities vs DDD (Domain Driven Design)

- Entities da Clean Architecture =/= Entity do Domain Driven Design
- Entity da Clean Architecture -> Conceito de camada de regra de negócio.
- Entity do Domain Driven Design -> Representação de algo único na aplicação.
- Entity do Clean Architecture usa os conceitos do DDD (Aggregations, Value Objects, Contracts).
- Entities = Agregados + Domain Services (Separados por camadas diferentes no domínio)
- Agregados
- Contratos -> Ex: Repositórios
- Serviços -> Definições de eventos
- Entity CA contém regra crítica e inváriavel da aplicação.
- Use Case pode variar de acordo com o fluxo, mas a Entity não, pois possui de regra solidificada.