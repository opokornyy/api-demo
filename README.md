# api-demo

## HOW TO RUN

### Run containerized app

```bash
docker compose up -d --build
```

### Run tests

```bash
# Run database
docker compose up -d db

# Set up ENV
export DATABASE_URL="postgresql://postgres:postgres@localhost:5432/userDb"
export SERVER_URL="0.0.0.0:8080"

# Run integration tests
make integration-test

# Run unit tests
make test
```

## Testing CURLs

### Create User Request

```bash
curl -X "POST" "http://localhost:8080/user" \
 -H 'Content-Type: text/plain; charset=utf-8' \
 -d $'{
    "id": "26908e04-868c-4d8e-85f8-6b1284dcf750",
    "name": "Mike Oxlong",
    "email": "mike@oxlong.cz",
    "date_of_birth": "2020-01-01T12:12:34+00:00"
}'
```

### Get User Request

```bash
curl "http://localhost:8080/user/26908e04-868c-4d8e-85f8-6b1284dcf750"
```

### Future work

- add more unit tests
- add more integration tests
- finish CRUD endpoints
- implement prometheus metrics
- create deployment manifests for k8s
- implement health checks
