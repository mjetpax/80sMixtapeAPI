# 80's Mixtape API

![cassette tapes](/assets/images/cassettes.png?raw=true)

The 80's mixtape API is a fun application that creates mixtapes of top pop hits from the 1980s. The project is built in GO and utilizes a Postgresql database. The project is dockerized and demonstrates some micro-services application strategies.

Author: Michael Babcock <mjetpax@gmail.com>

[![Coverage Status](https://coveralls.io/repos/github/mjetpax/80sMixtapeAPI/badge.svg?branch=master)](https://coveralls.io/github/mjetpax/80sMixtapeAPI?branch=master)

## Getting started

Getting started locally is simple. Make sure you have all the requirements on your local machine. Then run the Docker Compose commands in your terminal. Note, that Go is not required to run the dockerized application. It is only required for development.

### Requirements

* [Docker](https://docs.docker.com/get-started/)
* [Docker Compose](https://docs.docker.com/compose/install/)
* [Go programming language](https://golang.org/doc/install) (not required to run the app in Docker)
* [cURL](https://duckduckgo.com/?q=curl+getting+started) (optional)

### Running with Docker Compose

Note, that while developing & running with docker-compose, 80's Mixtape API will auto-compile and redeploy on script update... yay!

```bash
docker-compose up
```

This will launch the application and a development database. When the app is ready, you'll see this output `80's Mixtape API is listening on port 8080!` in your terminal.

(Seriously, that's it.)

### Use cURL to ping endpoints

Hit the health check endpoint with cUrl.

```bash
curl http://0.0.0.0:8080/health
```

For pretty formatted returns use jq. Example:

```bash
curl http://0.0.0.0:8080/health | jq .
```

Response:

```json
{
  "application_name": "80's Mixtape API",
  "message": "80's Mixtape API is running smooth!",
  "start_date": "2019-01-04 23:18:11.826575403 +0000 UTC",
  "up_time": "4.300778317s"
}
```

## API Endpoints

### GET /health

Check the app's start date up time and other bits of info.

Examples:

`curl http://0.0.0.0:8080/health`

Example result:

```json
{
  "application_name": "80's Mixtape API",
  "message": "80's Mixtape API is running smooth!",
  "start_date": "2019-01-04 23:18:11.826575403 +0000 UTC",
  "up_time": "4.300778317s"
}
```

### GET /cassette/:cassette_type

_cassette type is optional_

The endpoint `/cassette/:cassette_type` is the main endpoint for this API. It creates a play list JSON response structured as if it were a cassette tape, with A and B sides. Optionally, pass along one of three cassette types, **C60, C90, C120**. Cassette types are used to indicate cassette tape storage capacity. For example, C90 indicates 90 minutes of playback time, 45 minutes per side. If excluded C60 is selected by default.

Examples:

`curl http://0.0.0.0:8080/cassette`

`curl http://0.0.0.0:8080/cassette/C90`

Example result:

```json
{
  "tape_type": "C60",
  "side_a": {
    "songs": [
      {
        "id": 85,
        "title": "Karma Chameleon",
        "artist": "Culture Club",
        "year": 1984,
        "video": "https://music.youtube.com/watch?v=JvBInlDDQaU",
        "play_time": "4:21"
      },
      ...
    ],
    "play_time": "29:13"
  },
  "side_b": {
    "songs": [
      {
        "id": 56,
        "title": "I Can't Go For That",
        "artist": "Daryl Hall & John Oates",
        "year": 1982,
        "video": "https://music.youtube.com/watch?v=EHaJLvQgEOM",
        "play_time": "6:45"
      },
      ...
    ],
    "play_time": "27:39"
  }
}
```


### GET /songs/:last_value/:limit

Paginate through the API's song catalog with calls to the `/songs` endpoint. Pagination uses the seek method. Include the last id value from the previous page and the page size _limit_ with each GET request.

Example, walk through the song catalog with a page size of 10:

* Page 1: `curl http://0.0.0.0:8080/songs/0/10`
* Page 2: `curl http://0.0.0.0:8080/songs/10/10`
* Page 3: `curl http://0.0.0.0:8080/songs/20/10`

A list of song objects will be returned as a JSON response.

```json
[{
    "id": 1,
    "title": "Call Me",
    "artist": "Blondie",
    "year": 1980,
    "video": "https://music.youtube.com/watch?v=StKVS0eI85I",
    "duration_label": "2:15"
  },
  {
    "id": 2,
    "title": "Another Brick In The Wall",
    "artist": "Pink Floyd",
    "year": 1980,
    "video": "https://music.youtube.com/watch?v=HrxX9TBj2zY",
    "duration_label": "3:19"
  }]
```

## Testing

Use this command for running unit tests.

```bash
go test -v -cover -covermode=atomic ./...
```
