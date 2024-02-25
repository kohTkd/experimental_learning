FROM mysql:8

COPY ./docker-entrypoint-initdb.d /docker-entrypoint-initdb.d/
COPY ./my.cnf /etc/mysql/conf.d/my.cnf

CMD ["--character-set-server=utf8mb4", "--collation-server=utf8mb4_unicode_ci"]
