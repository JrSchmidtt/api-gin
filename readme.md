# api-postgresql Golang Boilerplate

## 1 - Criando abiente para Desenvolvimento local

1.1 - Criando container docker com postgres:
```bash
 docker run -d \
	--name some-postgres \
	-e POSTGRES_PASSWORD=mysecretpassword \
	-e PGDATA=/var/lib/postgresql/data/pgdata \
	-v /custom/mount:/var/lib/postgresql/data \
	postgres
```

1.9 - Rodando Aplicação (Bash)
```bash
go run main.go
```
