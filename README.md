# house-match-server

## DB
sudo docker run --name db-postgres -p 5432:5432 -e POSTGRES_PASSWORD=jei9sagheixohHoolohbaVae -e POSTGRES_USER=admin -d postgres:14
## API
docker run -d --name house-match-api -p 8080:8080 -e DB_HOST=172.17.0.1 -e DB_PORT=5432 -e DB_USER=admin -e DB_PASS=admin -e DB_NAME=housematch -e DB_SSLMODE=disable -e DB_MAXLIMIT=100 -e STORAGE_PATH=/home/storage -v /home/arthur/files:/home/storage arthur8davis/housematch-server
docker run -d --name house-match-api -p 8080:8080 -e DB_HOST=172.17.0.1 -e DB_PORT=5432 -e DB_USER=admin -e DB_PASS=jei9sagheixohHoolohbaVae -e DB_NAME=housematch -e DB_SSLMODE=disable -e DB_MAXLIMIT=100 -e STORAGE_PATH=/home/storage -v /root/storage/multimedia:/home/storage arthur8davis/housematch-server

ooPhahwan5ahveighu6fei9T
167.71.53.173

DB_HOST=localhost
DB_PORT=5432
DB_USER=admin
DB_PASS=admin
DB_NAME=housematch
DB_SSLMODE=disable
DB_MAXLIMIT=100

STORAGE_PATH=/home/arthur/files
