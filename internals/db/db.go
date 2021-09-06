package db

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-book-api/internals/entities/book"
)

type PGRepo struct {
	db *sqlx.DB
}

func CreateRepo(db *sqlx.DB) *PGRepo {
	return &PGRepo{db: db}
}

// AddEntities(entities []book.Book) error
func (repo *PGRepo) AddEntities(entities []book.Book) error {
	insertQuery := sq.Insert("books").Columns(
		"user_id",
		"title",
		"author",
		"isbn10",
		"isbn13",
	)

	for _, book := range entities {
		insertQuery = insertQuery.Values(
			book.UserId,
			book.Title,
			book.Author,
			book.ISBN10,
			book.ISBN13,
		)
	}

	insertQuery = insertQuery.RunWith(repo.db).PlaceholderFormat(sq.Dollar)
	_, err := insertQuery.Exec()
	if err != nil {
		return err
	}

	return nil
}

// ListEntities(limit, offset uint64) ([]book.Book, error)
func (repo *PGRepo) ListEntities(limit, offset uint64) ([]book.Book, error) {
	books := make([]book.Book, 0, limit)
	rowQuery := sq.Select(
		"user_id",
		"title",
		"author",
		"isbn10",
		"isbn13",
	).From(
		"books",
	).Offset(
		offset,
	).Limit(
		limit,
	)
	rowQuery = rowQuery.RunWith(repo.db)

	rows, err := rowQuery.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		book := book.Book{}
		err := rows.Scan(&book.UserId, &book.Title, &book.Author, &book.ISBN10, &book.ISBN13)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

// DescribeEntity(entityId uint64) (*book.Book, error)
func (repo *PGRepo) DescribeEntity(isbn10 string) (*book.Book, error) {
	descQuery := sq.Select(
		"user_id",
		"title",
		"author",
		"isbn10",
		"isbn13",
	).From(
		"books",
	).Where(
		sq.Eq{"isbn10": isbn10},
	)
	descQuery = descQuery.RunWith(repo.db).PlaceholderFormat(sq.Dollar)
	bookFromDB := book.Book{}
	err := descQuery.QueryRow().Scan(
		&bookFromDB.UserId,
		&bookFromDB.Title,
		&bookFromDB.Author,
		&bookFromDB.ISBN10,
		&bookFromDB.ISBN13,
	)
	if err != nil {
		return nil, err
	}
	return &bookFromDB, nil
}

// RemoveEntity(entityId uint64) error
func (repo *PGRepo) RemoveEntity(isbn10 string) error {
	deleteQuery := sq.Delete("books").Where(sq.Eq{"isbn10": isbn10})
	deleteQuery = deleteQuery.RunWith(repo.db).PlaceholderFormat(sq.Dollar)
	_, err := deleteQuery.Exec()
	return err
}

var dbKey = "db"

// Возвращает подключение к БД
func Connect(dsn string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Ошибка при подключении к БД: %v", err)
	}
	return db
}

// Добавляет подключение к БД в context
func AttachToContext(ctx context.Context, db *sqlx.DB) context.Context {
	return context.WithValue(ctx, &dbKey, db)
}

// Возвращает подключение к БД из context
func FromContext(ctx context.Context) *sqlx.DB {
	db, ok := ctx.Value(&dbKey).(*sqlx.DB)
	if !ok {
		log.Fatal("Ошибка при получении подключения к БД из context")
	}
	return db
}
