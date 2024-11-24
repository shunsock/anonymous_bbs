## System Architecture

## Local Development

### Overview

- container: `Docker Compose` to run the application locally
- database: `PostgreSQL` to store the data (volume is used to persist the data)

### Deployment

we use `.env.local` file to store the environment variables. 

```shell
source .env.local
docker compose up -d
```

## Remote (Production)

I used following stack. Of course, you can change the container runtime and database.

- container: `cloud run` to run the application
- database: `supabase` to store the data

### Deployment

we use `.env` file to store the environment variables. 

```shell
source .env
bash commands/deploy
```

