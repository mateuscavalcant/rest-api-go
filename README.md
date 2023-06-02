# REST API usando Go 

<img align="right" width="159px" src="https://i0.wp.com/cdn-images-1.medium.com/max/1200/1*lSUb1T4YW1td0UskwsGZ1w.gif?w=1920&ssl=1">


Este projeto é um exemplo de API em Go que lida com operações CRUD (Create, Read, Update, Delete) em uma coleção de carros. A API é implementada usando o pacote github.com/gorilla/mux para lidar com as rotas HTTP.

## Pré-requisitos
Certifique-se de ter o Go instalado em sua máquina antes de executar este projeto e instale o seguinte pacote:

```sh
$ go get -u github.com/gorilla/mux
```

## Utilização
Clone este repositório em sua máquina local.
Navegue até o diretório do projeto.
Execute o seguinte comando no seu terminal para iniciar o servidor:
```shell
go run main.go
```
O servidor será iniciado e estará ouvindo na porta 'http://localhost:8081/cars'

## Endpoints
A API possui os seguintes endpoints:
- GET /cars
Retorna todos os carros armazenados.

- GET /cars/brand/{brand}
Retorna carros filtrados pela marca especificada.

- GET /cars/model/{model}
Retorna carros filtrados pelo modelo especificado.

- GET /cars/volume/{volume}
Retorna carros filtrados pelo volume especificado.

- GET /cars/weight/{weight}
Retorna carros filtrados pelo peso especificado.

- GET /cars/co2/{co2}
Retorna carros filtrados pela emissão de CO2 especificada.

- GET /cars/{id}
Retorna um carro específico com base no ID fornecido.

- POST /cars
Cria um novo carro com base nos dados fornecidos.

- PUT /cars/{id}
Atualiza um carro existente com base no ID fornecido.

- DELETE /cars/{id}
Remove um carro existente com base no ID fornecido.

## Contribuição
Sinta-se à vontade para contribuir com este projeto. Você pode abrir uma issue para relatar problemas ou sugerir melhorias. Além disso, pull requests são bem-vindos.
