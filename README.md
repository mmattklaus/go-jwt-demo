# API Based Authentication using JWT in Golang (Sample Application) 
#### A framework-ish approach

Hello, world!

I decided to make Go my first language and since most applications I build required some form of client authentication, especially stateless authentication, I decided to do JSON Web Tokens (JWT) authentication with Go. 

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
    
    Build with so much :heart: