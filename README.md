Go Standart Project Layout
https://github.com/golang-standards/project-layout


Чистая архитектура / DDD

// project/
   |   bin/
// ├── cmd/
// │   └── api/
// │       └── main.go
// ├── internal/
// │   ├── domain/               // бизнес-сущности
// │   │   └── models.go 
// │   ├── repository/           // работа с БД
// │   │   ├── interfaces.go
// │   │   └── memory.go
// │   ├── service/              // бизнес-логика
// │   │   ├── interfaces.go
// │   │   └── user_service.go
// │   └── transport/            // коммуникация с клиентами
// │       └── http/
// │           ├── handlers.go
// │           ├── middleware.go
// │           └── server.go
// |   
// └── go.mod
// |   pkg