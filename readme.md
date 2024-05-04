# Ecom

Ecom(an API service) is an online marketplace that provides buyers and sellers with an avenue to meet and exchange goods and services. No chat system, Phone call is the method of communication between a buyer and a seller. if you are a seller, the phone number you provided when you registered on the platform will be showing on all of your product and if you are a buyer, you will see the seller phone number on any product you want to buy.

### Prerequisite

- Golang
- Postgresql
- Redis
- Docker

### configuration

- Make sure to set the appropriate value of the yaml file and the .env file

### Run on local

clone the repo

```
git clone https://github.com/mukailasam/ecom

```

move into the program directory

```

cd ecom

```

run the program

```
$ go run . -env local

```

### Run on container

clone the repo

```
git clone https://github.com/mukailasam/ecom

```

move into the program directory

```

cd ecom

```

start all services

```
$ docker compose up

```
