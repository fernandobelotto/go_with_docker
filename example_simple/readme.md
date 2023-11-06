# Using docker with go!

## Build
```bash
docker build -t go-docker .
```

## Run
```bash
docker run -it --rm go-docker
```

## Run with compose
```bash
docker-compose up --build --force-recreate --no-deps
```

