# Assig5BookApi


Structure

├── Models
│   ├── Book.go // Book models

|	├── Scheme.go // Book struct and tabel

├── Config
│   └── Database.go // Global DB

├── Controllers
│   └── Book.go // Book Controller

├── ApiHelpers
│   └── Response.go // response function

├── Routers
|   └── Routers.go // Routers

└── main.go


API

/book
GET : Get all book

POST : Create a new book
/book/:id

GET : Get a book

PUT : Update a book

DELETE : Delete a book

#Post Params

{
	"author": "Op Super John Doe Bilw",
	"name": "Implementation Golang",
	"category": "Knowledge"
}
