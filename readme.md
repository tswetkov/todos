# Todos Backend

## Instruction for starting up the project

```bash
docker pull postgres && \n
docker run --name=todos-db -e POSTGRES_PASSWORD='<your-password>' -p 5432:5432 -d --rm postgres && \n
migrate -path ./schema -database 'postgres://postgres:<your-password>@localhost:5432/postgres?sslmode=disable' up
```
