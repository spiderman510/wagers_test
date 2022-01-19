#!/bin/bash

docker-compose -f docker-compose.yml up -d --build

while ! docker exec -it prophet_mysql mysql -uuser -ppassword -e "SELECT 1" &> /dev/null ; do
    echo "Waiting for database ready..."
    sleep 1
done

echo "Start to initialize database..."
docker exec -it prophet_mysql mysql -uuser -ppassword -e "$(cat ./build/init_database.sql)"

echo "HTTP server is ready to serve..."