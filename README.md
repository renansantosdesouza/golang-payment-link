# payment-link
CRUD para link de pagamento com consumo de Api Rest e base de dados PostGres

Necessario criar uma base de dados no Postgres.
Apos criar a base criar a table com o script sql abaixo:

create table links(
	Id  varchar, 
	Name varchar, 
	Description  varchar, 
	Price decimal, 
	ShortUrl  varchar, 
	BodyJson  varchar
)

Configurar a string de conexao ao banco no arquivo config/config.yaml.
No mesmo arquivo tambem é necessario informar a configuracao para consumo da ApiRest geradora do link de pagamento.
