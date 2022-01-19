#!/bin/bash

docker-compose -f ./test/integration/docker-compose.yml up -d --build

while ! docker exec -it prophet_mysql_test mysql -uuser -ppassword -e "SELECT 1" &> /dev/null ; do
    echo "Waiting for database connection..."
    sleep 1
done

docker exec -it prophet_mysql_test mysql -uuser -ppassword -e "$(cat ./test/integration/init_test_database.sql)"

echo Done!