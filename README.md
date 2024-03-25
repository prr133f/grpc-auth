# gRPC authorization service
## Configure
Create .env file in root and set up required variables:
```env
POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_DB=
DATABASE_HOST=

APP_PORT=50051
APP_STATUS=DEBUG

JWT_SECRET=jwtsecretkey12345

DEFAULTUSER_EMAIL=
DEFAULTUSER_PWD=
DEFAULTUSER_ROLE=admin
```
## How to run
First of all, clone this repo on your local machine
Make `cd auth` and run
```sh
docker-compose up --build
```
