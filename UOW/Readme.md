# UOW

UoW é Unit of Work, uma forma de garantir que transações sejam feitas de forma atômicas (ou tudo ou nada) de uma só vez. A forma usada em Go foi criando uma função Do na qual recebe uma função como parâmetro capaz de executar o instrução SQL na base de dados.

o UoW também é capaz gerenciar repositórios e impedir que mais de uma transação seja feita ao mesmo tempo no mesmo repositório.