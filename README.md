# API Based Authentication using JWT in Golang (Sample Application) 
#### A framework-ish approach

Hello, World!

I decided to make Go my first language and since most applications I build required some form of client authentication, especially stateless authentication, I decided to do JSON Web Tokens (JWT) authentication with Go. 

## Project Description
I try as much to do separation of concerns. 
```bash
├── config
│   ├── config.go
├── database
│   ├── database.go
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

## Setup
1. Clone the repository

    ```bash
    git clone <repo-url>
    ```
 2. Update config file
    - Copy **example.config.toml** to **config.toml** in the project directory
    - Enter the application configuration in the newly created file <config.toml>
    - The following are required: DB_USERNAME (database username), DB_PASSWORD (database password), DB_DATABASE (name of database)
 
 3. Download dependencies
    
    I used go modules to manage all dependencies
 4. Start application