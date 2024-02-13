---
runme:
  id: 01HPHEFFGZD78E0TCWAMGWVXS2
  version: v2.2
---

# Go Unit of Work

This package is an implementation of the unit of work pattern in Golang. Unit of work is a design pattern that is widely used to help manage transactions and maintain consistency in an application. When combined with a clean architecture, this can provide a solid foundation for building scalable and maintainable applications.

## Installation

The following command is used for install this package into your golang project

```sh {"id":"01HPHEV9G90PX9TPJ213BXQXCJ"}
go get github.com/Muruyung/Go-UoW@latest
```

## How to Use

### Example using NewSession

This implementation is used if you want to customize the placement of new sessions, commits, and rollbacks

```go {"id":"01HPHFAYKN8ADEGSD38EPPNGVX"}
package main

import (
	"context"
	"database/sql"

	"github.com/Muruyung/Go-UoW/gouow" // import gouow package
)

var (
	db *gorm.DB // Example db engine
)

func main() {
	var (
		uow = gouow.Init() // Initialize unit of work
		ctx = context.TODO()
		err error
	)

	err = uow.NewSession(&ctx) // Create new session for start transaction
	if err != nil {
		return
	}

	err = RepositoryFunction(ctx) // Call first function
	if err != nil {
		_ = gouow.Rollback(err) // It will rollback RepositoryFunction if there is an error
		return
	}

	err = AnotherFunction(ctx) // Call another function
	if err != nil {
		_ = gouow.Rollback(err) // It will rollback RepositoryFunction and AnotherFunction if there is an error
		return
	}

	err = uow.Commit() // Commit all transaction process
	if err != nil {
		return
	}
}

// Example function for repository
func RepositoryFunction(ctx context.Context) error {
	var (
		err		error
		sqlDB	= db
		tx		= ctx.Value(gouow.TX_KEY)
	)

	if tx != nil {
		if dbTx := tx.(*gouow.TX); dbTx.UseTx {
			sqlDB = dbTx.GormDB()
		}
	}

	// Implement your repository logic here ...

	return err
}

// Example function for another logic outside of repository logic
func AnotherFunction(ctx context.Context) error {
	var err error
	// Implement your another logic here ...
	return err
}
```

### Example using BeginTx

This implementation is used if you want to use a simple method

```go {"id":"01HPHGYZPQ2EXZ99C03VSQT0TT"}
package main

import (
	"context"
	"database/sql"

	"github.com/Muruyung/Go-UoW/gouow" // import gouow package
)

var (
	db *gorm.DB // Example db engine
)

func main() {
	var (
		uow = gouow.Init() // Initialize unit of work
		ctx = context.TODO()
	)

	err := uow.BeginTx(ctx, func(ctxTx context.Context) error {
		err := RepositoryFunction(ctx) // Call first function
		if err != nil {
			return err // It will return error and rollback RepositoryFunction
		}

		err = AnotherFunction(ctx) // Call another function
		if err != nil {
			_ = gouow.Rollback(err) // It will return error and rollback RepositoryFunction and AnotherFunction
			return err
		}
	})
	if err != nil {
		return
	}
}

// Example function for repository
func RepositoryFunction(ctx context.Context) error {
	var (
		err		error
		sqlDB	= db
		tx		= ctx.Value(gouow.TX_KEY)
	)

	if tx != nil {
		if dbTx := tx.(*gouow.TX); dbTx.UseTx {
			sqlDB = dbTx.GormDB()
		}
	}

	// Implement your repository logic here ...

	return err
}

// Example function for another logic outside of repository logic
func AnotherFunction(ctx context.Context) error {
	var err error
	// Implement your another logic here ...
	return err
}
```
