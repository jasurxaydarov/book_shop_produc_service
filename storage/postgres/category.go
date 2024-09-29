package postgres

import (
	"context"
	"time"

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
			description,
			created_at,
			updated_at
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
		&resp.CreatedAt,
		&resp.UpdatedAt,
	)

	if err != nil {

		c.log.Error("err on db GetCategoryById", logger.Error(err))
		return nil, err
	}

	return &resp, nil
}

func (c *CategoryRepo) GetCategories(ctx context.Context, req *product_service.GetListReq) (*product_service.CategoryGetListResp, error) {

	offset := (req.Page - 1) * req.Limit

	var resp product_service.Category
	var res product_service.CategoryGetListResp
	qury := `
		SELECT 
			*
		FROM 
			categories
		WHERE
			deleted_at IS NULL
		LIMIT $1 OFFSET $2;
			
	`

	row, err := c.db.Query(
		ctx,
		qury,
		req.Limit,
		offset,
	)

	if err != nil {

		c.log.Error("err on db GetCategories", logger.Error(err))
		return nil, err
	}

	for row.Next() {

		row.Scan(
			&resp.CategoryId,
			&resp.CategoryName,
			&resp.Description,
			&resp.CreatedAt,
			&resp.UpdatedAt,
			&resp.DeletedAt,
		)

		res.Count++
		res.Category = append(res.Category, &resp)

	}

	return &res, nil
}


func (c *CategoryRepo) UpdateCategory(ctx context.Context, req *product_service.CategoryUpdateReq) (*product_service.Category, error) {

	time:=time.Now()

	query := `
			UPDATE
				categories
			SET
				name = $1,
				description = $2,
				updated_at = $3
			WHERE
				category_id = $4
			AND  
				deleted_at is null
			`

	_, err := c.db.Exec(
		ctx,
		query,
		req.CategoryName,
		req.Description,
		time,
		req.CategoryId,
	)
	if err != nil {

		c.log.Error("err on db UpdateCategory", logger.Error(err))
		return nil, err
	}

	resp, err := c.GetCategoryById(context.Background(), &product_service.GetByIdReq{Id: req.CategoryId})

	if err != nil {

		c.log.Error("err on db GetCategoryById", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *CategoryRepo) DeleteCategory(ctx context.Context, req *product_service.DeleteReq) (string, error) {

	time:=time.Now()

	query := `
			UPDATE
				categories
			SET
				updated_at = $1
			WHERE
				category_id = $2
			`

	_, err := c.db.Exec(
		ctx,
		query,
		time,
		req.Id,
	)
	if err != nil {

		c.log.Error("err on db DeleteCategory", logger.Error(err))
		return "", err
	}

	

	return "successfully deleted", nil
}