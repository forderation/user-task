# Welcome to User Task!
Go project demo of user task api, config for golang are located at config.toml

# How to start
## Install go-migrate tools
This tools is only for running migration, the migration sql are located at db/migration directory \
Reference go migrate: [Release Downloads](https://github.com/golang-migrate/migrate/releases)

```bash
$ curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$os-$arch.tar.gz | tar xvz
```
### MacOS
```bash
$ brew install golang-migrate
```
### Windows
Using [scoop](https://scoop.sh/)
```bash
$ scoop install migrate
```
### Linux (*.deb package)
```bash
$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
$ apt-get update
$ apt-get install -y migrate
```
## Running through docker compose
### spawn required docker container services
```bash
docker compose -f docker-compose.yaml up -d --build
```
### running migration
db_url parameter is dsn related after docker container is spawned
```bash
make migrateup db_url="mysql://root:root@tcp(127.0.0.1:52000)/user_task?x-tls-insecure-skip-verify=true"
```
### docker status
type docker ps to see status
```bash
docker ps

CONTAINER ID   IMAGE              COMMAND                  CREATED         STATUS         PORTS                                           NAMES
aa9890ae4fa8   api:user_task      "./main"                 3 minutes ago   Up 3 minutes   0.0.0.0:8081->8081/tcp                          api-user_task
32c1f6da98dd   mysql:8.0          "docker-entrypoint.s…"   12 hours ago    Up 12 hours    3306/tcp, 33060/tcp, 0.0.0.0:52000->52000/tcp   user-task-mysql_db_user_task-1
```
## API Documentation
There is postman documentation you can look up at `user_task_postman_collection.json` file