# Magfile Server

Backend for [Magfile](https://github.com/saltchang/magfile).

## Getting Started

### Requirements

- [Docker](https://www.docker.com/)

### Setup

Copy the `.env.example` file for making a `.env` file, and then configure your own environment variables.

```bash
cp .env.example .env
```

### Install [golang-migrate](https://github.com/golang-migrate/migrate)

With [Homebrew](https://brew.sh/):

```bash
brew install golang-migrate
```

### Pull docker image

```bash
docker pull postgres:13.2-alpine
```

### Launch postgres and create database

#### Launch postgres from docker

```bash
make postgres
```

The command will remove the existing docker container of postgres and run from the image again.
If the current container is running, it will automatically remove it.

#### Create the database

```bash
make createdb
```

If you've re-run the cotainer, please create database again.

#### Migrate up the database

Migrate up the database to current version:

```bash
make migrateup
```

### Run the service

```bash
go run main.go
```

### Access to the postgres

If you wish to access the postgres directly, run:

```bash
make accessdb
```
