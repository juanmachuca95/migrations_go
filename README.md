# migrations_go

Arquitectura hexagonal golang. Documentación mysql driver 
http://go-database-sql.org/index.html
Documentación: Obtener los valores de un body request from JSON.
https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body

# Guide golang api rest basic

    usefull link: https://blog.friendsofgo.tech/posts/como_crear_una_api_rest_en_golang/

# Router Mux 
    go get -u github.com/gorilla/mux

# Driver MySQL 

    go get -u github.com/go-sql-driver/mysql

# Variables de entorno .env

    go get -u github.com/joho/godotenv

# JWT JSON Web Token

    go get -u github.com/dgrijalva/jwt-go


# Password Bcrypt 

    go get golang.org/x/crypto/bcrypt


# Routes: Functions of migrations 

| Route           | Handlers           | Actions            | Return                 |
| --------------- |:------------------:| ------------------:|-----------------------:|
| /users          | GetUsersHandler    | Migra usuarios.    |  true || error         | 

