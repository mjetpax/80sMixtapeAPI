FROM golang:alpine AS buildenv
COPY . /go/src/github.com/mjetpax/80sMixtapeAPI
WORKDIR /go/src/github.com/mjetpax/80sMixtapeAPI
RUN go build -o ./80sMixtapeAPI .

FROM alpine
LABEL maintainer="Michael Babcock <mjetpax@gmail.com>"
LABEL description="80's Mixtape API container."
WORKDIR /app
COPY --from=buildenv /go/src/github.com/mjetpax/80sMixtapeAPI/80sMixtapeAPI /app/
ENTRYPOINT ./80sMixtapeAPI
EXPOSE 8080
CMD "80sMixtapeAPI"