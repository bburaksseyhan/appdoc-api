FROM golang:1.16-alpine as build-env
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN  go build -o /appdoc-api github.com/bburaksseyhan/appdoc-api/src/cmd/appdoc-server   

FROM alpine:3.14

RUN apk update \
    && apk upgrade\
    && apk add --no-cache tzdata curl

#RUN apk --no-cache add bash
ENV TZ Europe/Istanbul

WORKDIR /app
COPY --from=build-env /appdoc-api .
COPY --from=build-env /app/src/cmd/appdoc-server /app/

EXPOSE 80
CMD [ "./appdoc-api" ]