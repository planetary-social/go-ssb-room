# Creates a slim image that only holds what's absolutely necessary to run it.

FROM golang:1.17-alpine AS build
RUN apk add --no-cache \
      build-base \
      git \
      sqlite \
      sqlite-dev

RUN mkdir /app
WORKDIR /app
COPY . /app

RUN cd /app/cmd/server && go build && \
    cd /app/cmd/insert-user && go build

FROM alpine AS deploy
RUN apk add --no-cache \
      sqlite \
      sqlite-dev

RUN mkdir /app
COPY --from=build /app/cmd /app/cmd
COPY --from=build /app/start.sh /app/start.sh

EXPOSE 8008
EXPOSE 3000

RUN chmod +x /app/start.sh

CMD /app/start.sh
