# Deploy EKS com Docker File

Modulo demonstrando como criar um Docker file da aplicação e como subir no EKS.

## Docker

Toolkit que roda runtime containers que rodam processos utilizando o Hardware para rodar aplicações.

### Dockerfile

Dockerfile é um arquivo que gera a imagem da aplicação e sua estrutura para rodar o container

### docker-compose.yml

docker-compose.yml é um arquivo que gera o container a partir de uma imagem da aplicação com determinadas configurações da imagem.

- services -> aplicação que irá subir
- build -> Rodar o Dockerfile no diretório especificado.
- port -> Porta compartilhada entre o computador e o container
- volumes -> Compartilha o diretório entre o computador e o container

Para gerar um container docker:

- `docker compose up -d`

Para acessar o CMD do container:

- `docker-compose exec -it APP_NAME bash`

Para construir uma imagem Docker:

- `docker build -t NAME:latest -f Dockerfile.prod .`

- "`-t`" -> Tag com nome que preferir:versão
- "`-f`" -> Especifica o Dockerfile a ser utilizado

Para buscar imagens docker que contem um nome (Regular expression):

- `docker images | grep REGXP`

Rodar uma imagem docker:

- `docker run --rm -p 8080:8080 NAME:latest`

- "`--rm`" -> remove o container após executado
- "`-p`" -> porta onde será executado
- "`-d`" -> detached, libera o CMD após executar o container

## Kubernetes

Kubernetes é uma ferramenta que gerencia suas imagens de apliações, complementando Docker, garantindo que o container esteja disponível após uma falha assim como outros recursos.

[kind](https://kind.sigs.k8s.io/) é um Kubernetes que pode ser rodado localmente.

Cluster são nós(containers) que rodam máquinas virtuais de aplicações que são gerenciadas pelo Kubernetes.

### Service

Service no Kubernetes gerencia quais pods o usuário deve acessar em caso de multiplos pods iguais (Load balancing).

Para aplicar services no Kubernetes: 

- `kubectl apply -f path/to/service.yaml`.

Para listar services: 

- `kubectl get svc`

Para rodar service: 

- `kubectl port-forward svc/METADATA_NAME PC_PORT:EKS_PORT`

### Probes

Probes confere se o serviço EKS já subiu, se já esta pronto para ser utilizado. Se o pod nào estiver pronto o service não pode ser utilizado.

## Como instalar Kind

Precisa ter Go +1.16 e Docker para rodar o seguinte comando:

- `go install sigs.k8s.io/kind@v0.22.0`

## Comandos Kind

Para criar um cluster é necessário rodar o seguinte comando:

- `kind create cluster --name=goexpert`

kubectl é um comando para gerenciar Kubernetes clusters

- `kubectl cluster-info --context kind-goexpert`

Pod é a menor unidade do Kubernetes gerenciada pelo deployment que contém um container a ser usado.

Para rodar um pod do EKS, é preciso de um deployment configurado e depois executar o comando:

- `kubectl apply -f k8s/deployment.yaml`

Para ver os pods ativos:

- `kubectl get pods`

## Como executar um container no Kubernetes via Kind

1. Criar Dockerfile
2. Build Docker: `docker build -t NAME:VERSION -f DOCKERFILE .`
3. Criar cluster Kind: `kind create cluster --name CLUSTER_NAME`
4. Criar deployment.yaml (de preferência numa pasta) e usar a extensão vs code (Kubernetes) para configurar o deployment.
4.1. Importante! Se a versão do container for LATEST então deve aplicar a configuração `imagePullPolicy: Never` dentro de spec.containers
5. Carregar imagem docker dentro do Kind: `kind load docker-image NAME:VERSION --name CLUSTER_NAME`
6. Aplicar o deployment.yaml no Kubernetes: `kubectl apply -f path/to/deployment.yaml`
7. Verificar se não deu nenhum erro: `kubectl describe pods` e `kubectl get pods`