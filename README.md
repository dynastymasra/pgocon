# PGoCon

[![Go Report Card](https://goreportcard.com/badge/github.com/dynastymasra/pgocon)](https://goreportcard.com/report/github.com/dynastymasra/pgocon)
[![GoDoc](https://godoc.org/github.com/dynastymasra/pgocon?status.svg)](https://godoc.org/github.com/dynastymasra/pgocon)
[![License: MIT](https://img.shields.io/badge/license-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Go helper for Postgres database connection, base on [GORM](https://gorm.io/index.html) and [migrate](https://github.com/golang-migrate/migrate)
for database migration

## How to Use

import the library go get `github.com/dynastymasra/pgocon`

### Create a New Database Client

```go
config := pgocon.Config{		
	Database:     "database",
	Host:         "localhost",
	Port:         5432,
	Username:     "postgres",
	Password:     "root",
	Params:       "sslmode=disable",
	MaxIdleConn:  10,
	MaxOpenConn:  10,
	LogMode:      1,
	DebugEnabled: true,
}
		
// Return gorm client
conn, err := config.Client()	
if err != nil {		
	// Handle error
}

var model Model
if err := conn.Model("tables").Find("id = ?", id).Error; err != nil {
// Handle error
}
```

### Ping The Database Connection

```go
if err := config.Ping(); err != nil {
	// Handle error 
}
```

### Close the Database Connection

```go
if err := config.Close(); err != nil {
	// Handle error
}
```

### Set Database with Existing Connection

```go
config.SetDB(gorm.DB)
```

## How to Migrate

Import the library go get `github.com/dynastymasra/pgocon`, create a new folder in root and set the name to `migration`

### Create a New Migration File

```go
if err := pgocon.CreateMigrationFiles("filename_name"); err != nil {
    // Handler error
}
```

### Set The Database Connection

```go
migration, err := pgocon.Migration(gorm.DB)
if err != nil {
    // Handler error
}

if err := pgocon.RunMigration(migration); err != nil {
    // Handler error
}
```

### Rollback the Database Migration

```go
if err : pgocon.RollbackMigration(migration); err != nil { 
	// Handler error
}
```

## Development and Contributing

Open to anyone who wants to contribute to this library, `fork` and create a `Pull Request`