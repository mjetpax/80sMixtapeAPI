# 80's Mixtape API

 The 80's mixtape API is a fun application that creates mixtapes of top pop hits from the 1980s. The project is built in GO and utilizes a Postgresql database. The project is dockerized and demonstrates some basic micro-services application strategies.

 Author: Michael Babcock <mjetpax@gmail.com>

## Getting started

Getting started locally is simple. Run the following commands in your terminal.

### Start the app

```bash
go build .
```

```bash
./80sMixtapeAPI
```

### Use cUrl to ping endpoints

Hit the health check endpoint with cUrl.

```bash
curl localhost:8080/health
```

Response: `80's Mixtape API is listening to some rad jamz!`
