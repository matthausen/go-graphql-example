package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/matthausen/gql-example/pkg/user"
)

// UserService - The db string connection
type UserService struct {
	DbUserName string
	DbPassword string
	DbURL      string
	DbName     string
}

// Initialise - initialise postgres db
func (u *UserService) Initialise() error {
	targetSchemaVersion := 2
	dbConnectionString := u.getDBConnectionString()
	db, err := sql.Open("pgx", dbConnectionString)
	if err != nil {
		return err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	version, dirty, err := driver.Version()
	if dirty {
		log.Fatalf("ERROR: The current database schema is reported as being dirty. A manual resolution is needed")
	}
	log.Printf("Target database schema version is: %d and current database schema version is: %d", targetSchemaVersion, version)
	if version != targetSchemaVersion {
		log.Printf("Migrating database schema from version: %d to version %d", version, targetSchemaVersion)
		m, err := migrate.NewWithDatabaseInstance("file://../pkg/postgres/migrations", u.DbName, driver)
		if err != nil {
			return err
		}
		err = m.Steps(targetSchemaVersion)
		if err != nil {
			return err
		}
		return nil
	} else {
		log.Println("No database schema migrations need to be performed.")
	}
	if err != nil {
		log.Fatalf("ERROR: Could not determine the current database schema version")
	}
	return nil
}

// Create - add a user to the table
func (u *UserService) Create(name string, isPremium bool) (*string, error) {
	insertSQL := "insert into users(id, name, isPremium) values ($1, $2, $3)"
	ctx := context.Background()
	dbPool := u.getConnection()
	defer dbPool.Close()
	tx, err := dbPool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	id := uuid.New()
	idStr := id.String()
	_, err = tx.Exec(ctx, insertSQL, idStr, name, isPremium)
	if err != nil {
		log.Println("ERROR: Could not save the User item due to the error:", err)
		rollbackErr := tx.Rollback(ctx)
		log.Fatal("ERROR: Transaction rollback failed due to the error: ", rollbackErr)
		return nil, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}
	return &idStr, nil
}

// Update - update a record in user table
func (u *UserService) Update(id string, name string, isPremium bool) error {
	updateSQL := "update users set name = $1, isPremium = $2 where id = $3"
	ctx := context.Background()
	dbPool := u.getConnection()
	defer dbPool.Close()
	tx, err := dbPool.Begin(ctx)
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, updateSQL, name, isPremium, id)
	if err != nil {
		log.Println("ERROR: Could not save the User item due to the error:", err)
		rollbackErr := tx.Rollback(ctx)
		log.Fatal("ERROR: Transaction rollback failed due to the error: ", rollbackErr)
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Get - retrieve a single record of a user by id
func (u *UserService) Get(id string) (*user.UserItem, error) {
	selectSQL := "select id, name, isPremium, creation_timestamp, update_timestamp from users where id = $1"
	dbPool := u.getConnection()
	defer dbPool.Close()
	var userItem user.UserItem
	err := dbPool.QueryRow(context.Background(), selectSQL, id).Scan(&userItem.Id, &userItem.Name, &userItem.IsPremium, &userItem.CreatedOn, &userItem.UpdatedOn)
	if err != nil {
		return nil, err
	}
	return &userItem, nil
}

// List - retrieve all users from the table
func (u *UserService) List() ([]user.UserItem, error) {
	selectSQL := "select id, name, isPremium, creation_timestamp, update_timestamp from users"
	dbPool := u.getConnection()
	defer dbPool.Close()
	var userItems []user.UserItem
	rows, err := dbPool.Query(context.Background(), selectSQL)
	if rows != nil {
		defer rows.Close()
	}
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var userItem user.UserItem
		err = rows.Scan(&userItem.Id, &userItem.Name, &userItem.IsPremium, &userItem.CreatedOn, &userItem.UpdatedOn)
		if err != nil {
			return nil, err
		}
		userItems = append(userItems, userItem)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return userItems, nil
}

func (u *UserService) getConnection() *pgxpool.Pool {
	dbPool, err := pgxpool.Connect(context.Background(), u.getDBConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	return dbPool
}

func (u *UserService) getDBConnectionString() string {
	fmt.Println(u.DbUserName, u.DbPassword, u.DbURL, u.DbName)
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", u.DbUserName, u.DbPassword, u.DbURL, u.DbName)
}
