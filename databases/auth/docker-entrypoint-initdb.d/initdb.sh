#!/bin/bash
set -e

echo 'Initialization start.'

mysql -u "$MYSQL_ROOT_USER" -p${MYSQL_ROOT_PASSWORD} <<-EOSQL
    CREATE USER '$API_DB_USER' IDENTIFIED BY '$API_DB_PASSWORD';
    CREATE DATABASE $API_DB_NAME;
    GRANT ALL PRIVILEGES ON $API_DB_NAME.* TO $API_DB_USER;

    CREATE DATABASE $TEST_DB_NAME;
    GRANT ALL PRIVILEGES ON $TEST_DB_NAME.* TO $API_DB_USER;
EOSQL

echo 'Initialization finished'
