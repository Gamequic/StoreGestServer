# Folder directories

```
my_project/
├── main.go            // Main entry point
├── pkg/               // Reusable packages by features
│   ├── auth/          // Authentication package
│   │   ├── auth.go    // Authentication functionalities
│   │   └── middleware.go // Authentication middleware
│   ├── crud/          // Generic CRUD operations package
│   │   ├── crud.go    // Generic CRUD functions
│   │   └── middleware.go // Generic middleware
│   └── database/      // Database functions and configurations
│       └── db.go      // Database configuration and functions
├── features/          // Project features specific functionalities
|   ├── photos         // Crud for server public fotos
│   │   ├── handlers/  // HTTP handlers for photos
│   │   │   └── photos_handler.go
│   │   ├── models/    // photos-specific models
│   │   │   └── photos.go
│   │   └── repository/ // photos-specific repository
│   │       └── photos_repository.go
│   ├── user/          // User-related functionalities
│   │   ├── handlers/  // HTTP handlers for users
│   │   │   └── user_handler.go
│   │   ├── models/    // User-specific models
│   │   │   └── user.go
│   │   └── repository/ // User-specific repository
│   │       └── user_repository.go
│   ├── food/          // Food-related functionalities
│   │   ├── handlers/  // HTTP handlers for food
│   │   │   └── food_handler.go
│   │   ├── models/    // Food-specific models
│   │   │   └── food.go
│   │   └── repository/ // Food-specific repository
│   │       └── food_repository.go
│   ├── money/         // Money-related functionalities
│   │   ├── handlers/  // HTTP handlers for money
│   │   │   └── money_handler.go
│   │   ├── models/    // Money-specific models
│   │   │   └── money.go
│   │   └── repository/ // Money-specific repository
│   │       └── money_repository.go
│   └── order/         // Order-related functionalities
│       ├── handlers/  // HTTP handlers for orders
│       │   └── order_handler.go
│       ├── models/    // Order-specific models
│       │   └── order.go
│       └── repository/ // Order-specific repository
│           └── order_repository.go
└── utils/             // Shared utilities
└── middleware.go // Other middleware, utilities, etc.
```

## files on root

./runDev.sh // Command to run proyect on dev

./go.mod && ./go.sum // Dependencies for proyect

# To-do

* [ ] pkg 🔴
  * [x] database
  * [ ] features
    * [ ] Dashboard
      * [ ] SalesOnCurrentDate
      * [ ] MostSale
      * [ ] Orders
        * [ ] orders attended today
        * [ ] average per orders
    * [X] Photos 🟢
      * [X] CRUD routes
      * [X] CRUD service
      * [X] validation
    * [X] food
      * [X] CRUD routes
      * [X] CRUD service
      * [X] validation
    * [X] users
      * [X] CRUD routes
      * [X] CRUD service
      * [X] validation
      * [X] Encrypt
    * [X] money
      * [X] CRUD routes
      * [X] CRUD service
      * [X] validation
    * [X] orders
      * [X] CRUD routes
      * [X] CRUD service
      * [X] validation
    * [X] auth
      Not working due a missing smtp server
      (Recovery passoword)
      * [X] Routes
      * [X] Service
  * [ ] middlewares
    * [X] auth
    * [X] rootAuth
    * [X] errorHandler
    * [X] ValidatorHandler
* [ ] relationships between tables 🔴
* [ ] Run proyect file 🟢
* [ ] Explain erros captures 🟢
* [ ] Fix readme.md 🟢
* [ ] Unit testing
* [ ] Load the timezone from .env

## Flags
- 🔴 Urgent
- 🟢 Later

# We use SemVer to versioning

https://semver.org/spec/v2.0.0.html