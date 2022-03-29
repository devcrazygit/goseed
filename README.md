# Goseed - Golang + JWT + MongoDB Rest Api Boilerplate
## Description
This project is for integration of golang and jwt authentication with mogo.

## Prerequisites
- Install Go latest version from https://golang.org/dl
- Latest [MongoDB Community Edition](https://docs.mongodb.com/manual/administration/install-community/)

You can check golang version from follwing command

```
> go version
```

## Golang project initialize
Golang project init is relatively easy. Please open terminal in the new project directory you want and execute following command.

```
> go mod init goseed
```

Here, “goseed” is the project name. Then “go.mod” file created under the directory. All dependencies are registered here.

In previous Go versions, dependency management is slightly annoying. You have to down packages by “go get -u …”, But from this version(1.14), above command resolve all things.

### Folder Structure
```
+ controllers
|--- authcontroller.go
+ middlewares
|--- middlewares.go
+ models
   + db
   |-- mongodb.go
   + entity
   |-- user.go
   + service
   |-- userservice.go
+ routers
|--- index.go
+ utils
|--- index.go
main.go
```

Like any other frameworks, controllers wraps application logic whether api project or not. Middlewares embeds pre or post controller hooks. And models are responsible for database connection and DB logic.

And maybe the most important part in a project is route system. In our project, routers package is responsible for it. And trivial but useful functions are located in utils package.

## MongoDB with Mogo
Mogo is a wrapper for [mgo](https://github.com/globalsign/mgo) that adds ODM, hooks, validation and population process, to its raw Mongo functions. 

Mogo started as a fork of the bongo project and aims to be a re-thinking of the already developed concepts, nearest to the backend mgo driver. It also adds advanced features such as pagination, population of referenced document which belongs to other collections, and index creation on document fields.

It’s easy to use. Simply import package from github repository.
```
import "github.com/goonode/mogo"
```

Then, create `.env` file under the root path of the project and add parameters like this.
```
DB_CONNECTION_STRING=localhost
DB_NAME=goseed
```

It is convenient to store project configuration parameters. To handle .env file, I use “EnvVar” function implemented in “utils/index.go”

That’s it. We can get connection from anywhere by following.

## Middlewares
There are two functions in middlewares package. One is “ErrorHandler” function and another is “Authentication”.
```
//ErrorHandler is for global error
func ErrorHandler(c *gin.Context) {
c.Next()
if len(c.Errors) > 0 {
c.JSON(http.StatusBadRequest, gin.H{
            "errors": c.Errors,
        })
    }
}
```
“ErrorHandler” function is to handle application errors uncaught. It is useful to prevent application from broken by unexpected exceptions.

“Authentication” function in middlewares package deals with JWT. It checks token and get user information. If not, it rejects request.

Further information can be found [here](https://medium.com/@devcrazy/golang-gin-jwt-mogo-mongodb-orm-golang-authentication-example-52c3c1189488?sk=9169b794339b8aab56de9b99ec45b3ff)
## License

The Goseed project is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).