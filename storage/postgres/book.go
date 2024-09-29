package postgres

import (
	"context"
	"time"

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
				description,
				published_date
			)VALUES(
				$1,$2,$3,$4,$5,$6,$7,$8
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
		req.PublishedDate,
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
			description,
			published_date
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
		&resp.PublishedDate,
	)

	if err != nil {

		b.log.Error("err on db GetBookById", logger.Error(err))
		return nil, err
	}

	return &resp, nil
}

func (b *BookRepo) GetBooks(ctx context.Context, req *product_service.GetListReq) (*product_service.BookGetListResp, error) {

	offset := (req.Page - 1) * req.Limit

	var resp product_service.Book
	var res product_service.BookGetListResp

	qury := `
		SELECT 
			*
		FROM 
			books
		WHERE
			deleted_at IS NULL
		LIMIT $1 OFFSET $2;

			
	`

	row, err := b.db.Query(
		ctx,
		qury,
		req.Limit,
		offset,
	)

	if err != nil {

		b.log.Error("err on db GetBookById", logger.Error(err))
		return nil, err
	}

	for row.Next() {

		row.Scan(
			&resp.Title,
			&resp.AuthorId,
			&resp.CategoryId,
			&resp.Price,
			&resp.Stock,
			&resp.Description,
			&resp.PublishedDate,
			&resp.CreatedAt,
			&resp.UpdatedAt,
			&resp.DeletedAt,
		)

		res.Count++
		res.Book = append(res.Book, &resp)

	}
	return &res, nil
}

func (b *BookRepo) UpdateBook(ctx context.Context, req *product_service.BookUpdateReq) (*product_service.Book, error) {

	time := time.Now()

	query := `
			UPDATE
				books
			SET
				title = $1,
				author_id = $2,
				category_id =$3,
				price = $4,
				stock = $5,
				description = $6,
				published_date = $7,
				updated_at = $8
			WHERE 
				book_id = $9 
			AND  
				deleted_at is null
			`

	_, err := b.db.Exec(
		ctx,
		query,
		req.Title,
		req.AuthorId,
		req.CategoryId,
		req.Price,
		req.Stock,
		req.Description,
		req.PublishedDate,
		time,
		req.BookId,
	)
	if err != nil {

		b.log.Error("err on db UpdateBook", logger.Error(err))
		return nil, err
	}

	resp, err := b.GetBookById(context.Background(), &product_service.GetByIdReq{Id: req.BookId})

	if err != nil {

		b.log.Error("err on db GetBookyById", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (b *BookRepo) DeleteBook(ctx context.Context, req *product_service.DeleteReq) (string, error) {

	time := time.Now()

	query := `
			UPDATE
				books
			SET
				deleted_at = $1
			WHERE book_id = $2
			`

	_, err := b.db.Exec(
		ctx,
		query,
		time,
		req.Id,
	)
	if err != nil {

		b.log.Error("err on db DeleteBook", logger.Error(err))
		return "", err
	}

	return "successfuly deleted", nil
}
