# Ecom

Ecom is an online marketplace that provides buyers and sellers with an avenue to meet and exchange goods and services, it's like jiji.ng web platform but unlike jiji.ng, it's only backend Rest API no frontend and it only support phone call as a means of communication between sellers and and buyers, if you are a seller, the phone number you provided when registered on the platform will be showing on all of your product and if you are a buyer, you will see the seller phone number on any product you want to buy.

## How TO Run

### Required Environment

- Postgre
- Redis

### Env configuration
- you should open the .env file and modify it

### Run 

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
$ go run .

```

NOTE: make sure you have the postgres server, redis server and the .env all setup before running the program

## API Techstack
- Go
- Postgres
- Redis

### API Endpoints
- /api/index/home
- /api/index/about
- /api/index/contact
- /api/auth/register
- /api/auth/login
- /api/auth/logout
- /api/auth/verify_email/{username}/{token}
- /api/auth/password/update
- /api/auth/register/password/reset
- /api/auth/password/reset/for/{username}/{token}
- /api/product/create
- /api/product/read
- /api/product/update
- /api/product/delete
- /api/product/image/create
- /api/account/profile/{username}/read
- /api/account/profile/update
- /api/account/profile/delete

## Usage Example

- **Register Functon API Callback(POST REQUEST)**

127.0.0.1:8080/api/auth/register

```json

{
    "username": "cnerd",
    "email": "cnerd@example.com",
    "firstName": "Computer",
    "lastName": "Nerd",
    "password": "example",
    "phone": "07000000000"
}

```
Response:
```json
{
    "Status": 200,
    "Message": "Account Creation Successfull",
    "Detail": "check your email for verification",
    "Path": "/api/auth/register",
    "Redirect": "/api/auth/login"
}
```


- **Email Verification Functon API Callback(GET REQUEST)**

http://127.0.0.1:8080/api/auth/verify_email/cnerd/a196b952ef1d33a6244db8fa275a1f2cc97ad4943cd12a6342db04a4a197c30a

Response:
```json
{"Status":200,"Message":"Successfully verified your email","Detail":null,"Path":null,"Redirect":null}
```


- **Login Functon API Callback (POST REQUEST)**

127.0.0.1:8080/api/auth/login

```json

{
    "email":"example@example.com",
    "password":"pass"
}

```

Response:

```json
{
    "Status": 200,
    "Message": "Successfully Login",
    "Detail": null,
    "Path": "/api/auth/login",
    "Redirect": "/api/index/home"
}
```

- **Logout Functon API Callback (GET REQUEST)**

127.0.0.1:8080/api/auth/logout


Response:

```json
{
    "Status": 200,
    "Message": "successfully logout",
    "Detail": "you successfully logout your account",
    "Path": "/api/auth/logout",
    "Redirect": "/api/auth/login"
}
```



- **Create Product Functon API Callback (POST REQUEST)**

127.0.0.1:8080/api/product/create

```json
{
    "name": " NEW Iphone ",
    "category": "phones",
    "price": 60000.00,
    "description": "Tecno povouir 4 is a quality phone that has a strong battery life",
    "other": {
        "color": "black",
        "operating system": "android",
        "battery": "6000 mAh",
        "screen size": "7.0 inches",
        "RAM": "3 GB",
        "selfie camera": "8 MP",
        "internal storage": "32 GB",
        "resolution": "720 x 1640"
    }
}
```

Response:

```json
{
    "Status": 200,
    "Message": "Successfully Created",
    "Detail": null,
    "Path": "/api/product/create",
    "Redirect": null
}
```

NOTE: the product is not fully created, image/images is needed to be associate with it, the frontend handles sending http message to the two endpoints /api/product/image/create, /api/product/create.


- **Create Product Image Functon API Callback (POST REQUEST)**

127.0.0.1:8080/api/product/image/create

explanation on how it should be handle at the frontend:
it associate image/images to a created product(metadata) by creating the product(metadata) first then it returned the post(product) id to the client. the client then send the images/images with the post(product) id, and the server then associate the image/images and the product(metadata).

Screenshot

create product(metadata) first
![alt text](https://github.com/ftsog/ecom/blob/main/static/images/product.jpg?raw=true)

post(product) id stored on the client side after product(metadata) created
![alt text](https://github.com/ftsog/ecom/blob/main/static/images/postid.jpg?raw=true)

associate the image/images to product(metadat) created
![alt text](https://github.com/ftsog/ecom/blob/main/static/images/image.jpg?raw=true)



Response:

```json
{
    "Status": 200,
    "Message": "Successfully upload image",
    "Detail": null,
    "Path": "/api/product/image/create",
    "Redirect": null
}
```


*With the above giving example on how to use the API and by looking into the program you should be able to use the rest of the API*






