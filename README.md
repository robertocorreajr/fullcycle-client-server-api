# fullcycle-client-server-api

## Desafio: 
Olá dev, tudo bem?
 
Neste desafio vamos aplicar o que aprendemos sobre webserver http, contextos,
banco de dados e manipulação de arquivos com Go.
 
Você precisará nos entregar dois sistemas em Go:
- client.go
- server.go
 
Os requisitos para cumprir este desafio são:
 
O client.go deverá realizar uma requisição HTTP no server.go solicitando a cotação do dólar.
 
O server.go deverá consumir a API contendo o câmbio de Dólar e Real no endereço: https://economia.awesomeapi.com.br/json/last/USD-BRL e em seguida deverá retornar no formato JSON o resultado para o cliente.
 
Usando o package "context", o server.go deverá registrar no banco de dados SQLite cada cotação recebida, sendo que o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms e o timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.
 
O client.go precisará receber do server.go apenas o valor atual do câmbio (campo "bid" do JSON). Utilizando o package "context", o client.go terá um timeout máximo de 300ms para receber o resultado do server.go.
 
Os 3 contextos deverão retornar erro nos logs caso o tempo de execução seja insuficiente.
 
O client.go terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}
 
O endpoint necessário gerado pelo server.go para este desafio será: /cotacao e a porta a ser utilizada pelo servidor HTTP será a 8080.
 
Ao finalizar, envie o link do repositório para correção.


## Instruções:
### FullCycle Client-Server API

Este projeto demonstra a interação entre um servidor Go e um cliente Go. O servidor obtém a cotação atual do dólar em relação ao real a partir de uma API externa, armazena no banco de dados SQLite, e o cliente requisita essa cotação ao servidor e a salva em um arquivo local.

**Tecnologias Utilizadas**
	•	Go (Golang) 1.23.4
	•	SQLite (via github.com/mattn/go-sqlite3)
	•	Docker e Docker Compose

**Requisitos**
	•	Docker instalado (recomendado: última versão)
	•	Docker Compose instalado (recomendado: última versão)
	•	Acesso à internet (para consultar a API externa de cotação)

**Como Executar**
1.	Clonar o Repositório:
```bash
git clone https://github.com/robertocorreajr/fullcycle-client-server-api.git
cd fullcycle-client-server-api
```

2.	Construir as Imagens e Subir os Contêineres:
```bash
docker-compose build
docker-compose up
```

Ao executar o comando acima:
	•	O server será construído e iniciado na porta 8080.
	•	O client será construído, aguardará o servidor iniciar e fará uma requisição para obter a cotação.

3.	Verificando a Cotação no Cliente:
Se tudo ocorrer bem, você verá no terminal algo como:
```plaintext
Cotação salva com sucesso: 6.0449
```

Esse valor é apenas um exemplo. A cotação real irá variar conforme a API externa.

4.	Arquivo cotacao.txt:
O client salva a cotação no arquivo cotacao.txt. Esse arquivo fica dentro do contêiner do cliente.
Para inspecionar o arquivo, basta copiá-lo do contêiner para sua máquina local:
```bash
docker cp client:/app/cotacao.txt .
cat cotacao.txt
```

O conteúdo será no formato:
```plaintext
Dólar: 6.0449
```

5.	Consultando o Servidor Manualmente:
O servidor expõe um endpoint /cotacao na porta 8080. Você pode fazer uma requisição HTTP manual (por exemplo, via curl):
```bash
curl http://localhost:8080/cotacao
```

Você receberá uma resposta JSON com a cotação atual:
```json
{"bid": "6.0449"}
```

6.	Parando os Contêineres:
Pressione CTRL + C no terminal onde está rodando o docker-compose up para parar os contêineres.
Caso queira removê-los:
```bash
docker-compose down
```


**Observações**
- Caso esteja utilizando um Mac com chip M1 (ARM64), as configurações do Dockerfile e do docker-compose foram preparadas para rodar nativamente.
- O projeto cria ou utiliza o banco cotacoes.db no contêiner do servidor. Para manter persistência local, um volume Docker é utilizado.