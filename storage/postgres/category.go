package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/book_shop_produc_service/genproto/product_service"
	"github.com/saidamir98/udevs_pkg/logger"
)

type CategoryRepo struct {
	db  *pgx.Conn
	log logger.LoggerI
}

func NewCategoryRepo(db *pgx.Conn, log logger.LoggerI) CategoryRepoI {

	return &CategoryRepo{db: db, log: log}
}

func (c *CategoryRepo) CreateCategory(ctx context.Context, req *product_service.CategoryCreateReq) (*product_service.Category, error) {
	id := uuid.New()
	query := `
		INSERT INTO
			categories(
				category_id,
				name,
				description 
			)VALUES(
				$1,$2,$3
			)
			`

	_, err := c.db.Exec(
		ctx,
		query,
		id,
		req.CategoryName,
		req.Description,
	)
	if err != nil {

		c.log.Error("err on db CreateCategory", logger.Error(err))
		return nil, err
	}

	resp, err := c.GetCategoryById(context.Background(), &product_service.GetByIdReq{Id: id.String()})

	if err != nil {

		c.log.Error("err on db GetCategoryById", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *CategoryRepo) GetCategoryById(ctx context.Context, req *product_service.GetByIdReq) (*product_service.Category, error) {

	var resp product_service.Category
	qury := `
		SELECT 
			category_id,
			name,
			description 
		FROM 
			categories
		WHERE
			category_id = $1
	`

	err := c.db.QueryRow(
		ctx,
		qury,
		req.Id,
	).Scan(
		&resp.CategoryId,
		&resp.CategoryName,
		&resp.Description,
	)

	if err != nil {

		c.log.Error("err on db GetCategoryById", logger.Error(err))
		return nil, err
	}

	return &resp, nil
}
