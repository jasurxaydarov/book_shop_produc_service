package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/book_shop_produc_service/genproto/product_service"
	"github.com/saidamir98/udevs_pkg/logger"
)

type BookRepo struct {
	db  *pgx.Conn
	log logger.LoggerI
}

func NewBookRepo(db *pgx.Conn, log logger.LoggerI) BookRepoI {

	return &BookRepo{db: db, log: log}
}

func (b *BookRepo) CreateBook(ctx context.Context, req *product_service.BookCreateReq) (*product_service.Book, error) {
	id := uuid.New()

	query := `
		INSERT INTO
			books(
			 	book_id,
				title,
				author_id,
				category_id,
				price,
				stock,
				description
			)VALUES(
				$1,$2,$3,$4,$5,$6,$7
			)
			`

	_, err := b.db.Exec(
		ctx,
		query,
		id,
		req.Title,
		req.AuthorId,
		req.CategoryId,
		req.Price,
		req.Stock,
		req.Description,
	)
	if err != nil {

		b.log.Error("err on db CreateBook", logger.Error(err))
		return nil, err
	}

	resp, err := b.GetBookById(context.Background(), &product_service.GetByIdReq{Id: id.String()})

	if err != nil {

		b.log.Error("err on db GetBookyById", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (b *BookRepo) GetBookById(ctx context.Context, req *product_service.GetByIdReq) (*product_service.Book, error) {

	var resp product_service.Book
	qury := `
		SELECT 
			title,
			author_id,
			category_id,
			price,
			stock,
			description
		FROM 
			books
		WHERE
			book_id = $1
	`

	err := b.db.QueryRow(
		ctx,
		qury,
		req.Id,
	).Scan(
		&resp.Title,
		&resp.AuthorId,
		&resp.CategoryId,
		&resp.Price,
		&resp.Stock,
		&resp.Description,
	)

	if err != nil {

		b.log.Error("err on db GetBookById", logger.Error(err))
		return nil, err
	}

	return &resp, nil
}
