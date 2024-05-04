# Ecom

Ecom(an API service) is an online marketplace that provides buyers and sellers with an avenue to meet and exchange goods and services. No chat system, Phone call is the method of communication between a buyer and a seller. if you are a seller, the phone number you provided when you registered on the platform will be showing on all of your product and if you are a buyer, you will see the seller phone number on any product you want to buy.

### configuration

- Make sure to set the appropriate value of the config.yaml file and the .env file

### Run on local

- Make sur to have Go installed on your local machine
- Make sure you have postgresql(with auth) and redis(with auth) running on your local machine.
- Make sure you have the database.sql at models/database.sql schema executed against the running postgresql
- Make sure to set the appropriate value of the config.yaml file

### Prerequisite

- Golang
- Postgresql
- Redis

### Clone the repo

```
git clone https://github.com/mukailasam/ecom

```

### Move into the program directory

```

cd ecom

```

### Start the program

```
$ go run . -env local

```

</br>

### Run on container

- Make sure to set the appropriate value of the .env file

### Prerequisite

- Docker

### Clone the repo

```
git clone https://github.com/mukailasam/ecom

```

### Move into the program directory

```

cd ecom

```

### Start all services

```
$ docker compose up

```
