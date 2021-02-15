# Magfile Server

Backend for [Magfile](https://github.com/saltchang/magfile).

## Getting Started

- install [migrate](https://github.com/golang-migrate/migrate)

```bash
brew install golang-migrate
```

- pull docker image

```bash
docker pull postgres:13.2-alpine
```

- launch postgres and create database

```bash
make postgres
make createdb
```

- access the db

```bash
make accessdb
```
