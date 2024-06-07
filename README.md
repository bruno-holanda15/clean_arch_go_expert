# Clean Architecture

Projeto desenvolvido em Go para entrega do desafio do Goexpert Clean Architecture. 
O mesmo possui um objetivo básico de criação e listagem de Orders por meio de 3 API's/servers diferentes, sendo eles web, gRPC e Graphql expostos para comunicação nas seguintes portas.
- Webserver port 8000
- gRPC port 50051
- Graphql: port 8080

## Iniciando Servers

Para iniciar os servidores já considerando a criação do banco de dados para armazenamento das orders.

```
make up
```

Para verificar os logs dos servidores.
```
make server-logs
```

Para o "desligamento" dos containers.
```
make down
```

## Testando

- webserver: Temos como base o documento [orders.http](api/orders.http)
- gRPC: Sugerimos instalar o [evans](https://github.com/ktr0731/evans) e usá-lo como o client para esse caso.
- Graphql: Interaja com o endpoint http://localhost:8080 e no mesmo estará rodando o GraphQL playground.

## Documentações libs core utilizadas

* [webserver](https://github.com/go-chi/chi)
* [gRPC](https://grpc.io/docs/languages/go/quickstart/)
* [GraphQL](https://gqlgen.com/getting-started/)
* [wire](https://github.com/google/wire)
* [evans](https://github.com/ktr0731/evans)