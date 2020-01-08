# the clean architecture by go with gin, gorm

## directory
```
├── README.md
├── domain                          // Enterprise Business Rules (Entitie)
│   └── model
│       └── Users.go
├── infrastructure                  // Framwork, Drivers (External Interfaces)
│   ├── Config.go
│   ├── DB.go
│   └── Routing.go
├── interface                       // Interface Adapters (Controller, Gateway, Presenter)
│   ├── controller
│   │   ├── App.go
│   │   ├── Context.go
│   │   └── UsersController.go
│   └── database
│       ├── DB.go
│       └── UsersRepository.go
├── main.go
├── registry
│   └── registry.go
└── usecase                          // Application Business Rules (Use Case)
    ├── UserInteractor.go
    └── UserRepository.go
```

## environment
- macOS Mojave 10.14.6
- go: 13.4
- mysql: 5.7
- postgres: 12.1
- date: 08 January, 2020

