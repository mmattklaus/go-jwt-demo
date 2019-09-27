# API Based Authentication using JWT in Golang (Sample Application) 
#### A framework-ish approach

Hello, world!

I decided to make Go my first language and since most applications I build required some form of client authentication, especially stateless authentication, I decided to do JSON Web Tokens (JWT) authentication with Go. 

## Table of Content
1. [Project Description](#project-description)
2. [Setup](#setup) (get it running)

## Project Description
I try as much to do separation of concerns. 
```bash
├── config
│   ├── config.go
├── database
│   ├── database.go
│   ├── migrations.go
├── helpers
│   ├── helpers.go
├── middleware
│   ├── auth.go
├── models
│   ├── User.go
│   ├── UserController.go
├── output
│   ├── build-linux-amd64
│   ├── build-linux-arm
│   ├── build-windows-amd64.exe
├── router
│   ├── api.go
├── .gitignore
├── example.config.toml
├── go.mod
├── go.sum
├── main.go
└── README.md
```
> Go ahead and explore this files. I believe the code should seem straight forward. However, just a few important tips:
- *_router/api.go_*: This os where you define all your routes (endpoints). A simple consist of two important parts: - a path (e.g. `/user/update`) and a Controller function (e.g. `UserController.Update(...)`). I recommend that all controller functions are wrapped in a logger function so that the access to the routes are logged. If a route is protected (i.e. only members with a valid token are allowed to access it), then the route should be wrapped around the `IsAuthorized` middleware func (e.g. `..."user/update", middleware.IsAuthorized(UserController.Update)`). I recommend this file stays clean. It should only map routes Avoid writing application logic within the route file. (except otherwise)
- *_models/UserController.go_*: This is where you should place your application (route) logic. If a function needs to interact with a model, it's better to define the function within the model (e.g. `User.go`) and simply make a call to that function (e.g. calling the function to get a user with a given ID: `user, err := User.Find(ID)`).

- *_models/User.go_*: This field contains your model schema (table design) defined as a struct. Each model most extend the `gorm.Model`. You can read the [gorm doc](https://gorm.io/docs/models.html) for more. This file also contains the functions which are defined on the model. It's recommend that these functions should be specific and straight forward and interact directly with the model only (every other logic should be placed in the controller calling this function).

- *_database/migrations.go_*: This file contains one simple working function `InitMigrations` which gets called each time you start your app. If you've used database migrations before in other languages/frameworks then you probably have a clue about what is happening here. What this function does simply is creating your database tables the first time you run it (described by your models). You do not have to worry about manually creating your tables using a database tool (e.g PHPMyAdmin). Something to note however is that even though this function gets called each time, [gorm](https://gorm.io/docs/migration.html) **would not** create these tables each time it runs thereby result to errors. Instead the library only updates your database table if there are changes added to your model definition. Again, read their [docs](https://gorm.io/docs/migration.html) if you're not familiar with this.

- *_output/\*_*: This directory contains the final compiled version of the app. See setup step 5 for more details. 

## Setup
###### Note: Make sure you've got golang installed on your system. For me, it was Go v1.12.9
1. Clone the repository

    ```shell script
    $ git clone <repo-url>
    ```
 2. Update config file
    - Copy **example.config.toml** to **config.toml** in the project directory
    - Enter the application configuration in the newly created file <config.toml>
    - The following are required: DB_USERNAME (database username), DB_PASSWORD (database password), DB_DATABASE (name of database) & DB_PORT (port on which application will be serve e.g. 8001)
     
     NOTE: 
     > For this implementation  I used a  MySQL database and I used the [gorm](https://gorm.io) package for my Object Relational Mapping (ORM).  However, you can go ahead to use any other database  [supported gorm dialect](https://gorm.io/docs/dialects.html).
 
 3. Download dependencies
    
    I used [go modules](https://blog.golang.org/using-go-modules) for dependencies management.
    ```shell script
    $ go mod download
    ```
 4. Start application
    ```shell script
    $ go run main
    ```
 5. Build your app
    One interest thing about Golang is that it compiles to Byte code (zeros & ones). That means that an entire application can compile down to a single file (binary) which you can carry around (instead of the entire application consisting of folders). And most interestingly, you can build your app on one platform for another (e.g. develop your app on Windows & build it for a linux system - on Windows)
    The following command would build the entire app:
    ```shell script
    $ go build .
    ```
    By default, the build command would build for your system specifications. To build for another system, you need to specify the 2 main session variables (**GOOS** - Go-OS and **GOARCH** - Go architecture (arch)). <br/><br/>
    e.g. 
    1. Build for Linux OS, 64bits
        ```shell script
        $ GOOS=linux GOARCH=amd64 go build .
        ```
    2. Build for an arm processor (like a Raspberry Pi)
        ```shell script
        $ GOOS=linux GOARCH=arm go build .
        ```
    3. Build for windows, 64bits
        ```shell script
        $ GOOS=windows GOARCH=amd64 go build .
        ```
     - To provide the output name & location, provide the **-o** option. (e.g. **$ go build . -o output/build-linux-amd64** would build a linux binary into **output/build-linux-amd64**)
     
    <br/>
    
    I provided a build for Linux (amd64), Windows (amd64) and Raspberry Pi (arm) in the project. They can be found in the output folder. You can simply download then and use without setting up the app. To used them you'll need to setup your database credentials to match the default as defined in the config file (config.toml):
    
    DB_HOST="127.0.0.1" <br/>
    DB_USERNAME="root" <br/>
    DB_PASSWORD="" <br/>
    DB_DATABASE="gophers" <br/>
 > Tip: You can extend the application to receive commandline arguments like DB_USERNAME, DB_PASSWORD & DB_NAME which will not require these defaults.
    
Build with so much :heart:                                       
