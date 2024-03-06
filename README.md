# Ristek Go Backend

## Description

This is the repository made for RISTEK Open Recruitment Web Development project. This repository specifies on the back-end.

## Getting Started

### Dependencies

* Go 1.16 or later
* Gin framework: `go get -u github.com/gin-gonic/gin`
* GORM SQLite driver: `go get -u gorm.io/driver/sqlite`
* GORM framework: `go get -u gorm.io/gorm`
* CompileDaemon: `go install github.com/githubnemo/CompileDaemon`

### Executing program

To run the project, you need to have Go 1.16 or later installed on your system. Once you have Go installed, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/zerothy/ristek-go-be.git
```

2. Navigate to the project directory:

```bash
cd ristek-go-be
```

3. Install all dependencies:

```bash
go get -u github.com/gin-gonic/gin
go get -u gorm.io/driver/sqlite
go get -u gorm.io/gorm
go install github.com/githubnemo/CompileDaemon
```

4. Start the server:

```bash
CompileDaemon -command="./server"
```


## Built With

* [Go](https://golang.org/) - Backend framework
* [Gin](https://gin-gonic.com/) - Web framework
* [GORM](https://gorm.io/) - Object-relational mapper

## Authors

* **Joe Mathew Rusli** - [zerothy](https://github.com/zerothy)