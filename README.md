## Requirements and Bonuses

1. [X] Implement a Rest API with CRUD functionality.
2. [X] Database: MySQL or PostgreSQL.
3. [X] Unit test as much as you can. ( only storage )
4. [X] Set up service with docker compse.

----

1. [X] Cleanly structured
2. [X] Authen with jwt
3. [X] generate query with sqlc instead of raw sql
4. [X] Private and public routes
5. [X] CI integration


### **API Endpoints**

import postman collection (minibank.json)


## Database UML

![Database UML](/bank.png "Database UML")

## Technology Stack

- **Go 1.17**: *Leverage the standard libraries as much as possible*
- **SQLc**: *Generates efficient native SQL CRUD code*
- **PostgreSQL**: *RDBMS of choice because of faster read due to its indexing model and safer transaction with better isolation levels handling*
- **Fiber**: *Fast and have respect for native net/http API*
- **JWT Token**: *Also implemented to demonstrate the decoupility*
- **Golang-Migrate**: *Efficient schema generating, up/down migrating*
- **Docker** + **Docker-Compose**: *Containerization, what else to say ...*
- **Viper**: *Add robustness to configurations*
- **Github Actions CI**: *Make sure we don't push trash code into the codebase*

## Booting Up

```bash
docker-compose build

docker-compose up

# docker-compose down
```

## Todos
- add manifest K8s
- deploy with Github Action and ArgoCD
- add replica postgres
- add cache DB ( Redis )
- add more testable for repository, usecase, ...
- add pagination and meta data
