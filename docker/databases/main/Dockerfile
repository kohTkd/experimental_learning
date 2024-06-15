FROM postgres:alpine

COPY ./docker-entrypoint-initdb.d /docker-entrypoint-initdb.d/
ENV POSTGRES_INITDB_ARGS --locale-provider=icu --icu-locale=ja_JP.UTF-8
