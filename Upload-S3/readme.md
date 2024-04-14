## Como fazer upload no S3

1. Criar um usuário sem acesso ao painel de controle com acesso full ao bucket.
2. Colocar no uploader/main.go o ID KEY e SECRET KEY do usuário gerado na linha 24 e 25.
3. Criar um bucket simples com o mesmo nome no uploader da linha s3Bucket.

Struct no channel é o menor tamanho possível.