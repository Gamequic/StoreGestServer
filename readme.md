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

* [ ] pkg
  * [ ] auth
  * [x] database
  * [ ] features
    * [ ] food
      * [ ] CRUD routes
      * [ ] CRUD service
      * [ ] validation
    * [ ] users
      * [X] CRUD routes
      * [ ] CRUD service
      * [ ] validation
    * [ ] money
      * [ ] CRUD routes
      * [ ] CRUD service
      * [ ] validation
    * [ ] orders
      * [ ] CRUD routes
      * [ ] CRUD service
      * [ ] validation
* [ ] Run proyect file
