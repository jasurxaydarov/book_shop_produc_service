package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/book_shop_produc_service/genproto/product_service"
	"github.com/saidamir98/udevs_pkg/logger"
)

type AuthRepo struct {
	db  *pgx.Conn
	log logger.LoggerI
}

func NewAuthRepo(db *pgx.Conn, log logger.LoggerI) AuthRepoI {

	return &AuthRepo{db: db, log: log}
}

func (u *AuthRepo) CreateAuth(ctx context.Context, req *product_service.AuthorCreateReq) (*product_service.Author, error) {
	id := uuid.New()
	query := `
		INSERT INTO
			authors (
				author_id,
				name,
				bio 
			)VALUES(
				$1,$2,$3
			)
			`

	_, err := u.db.Exec(
		ctx,
		query,
		id,
		req.AuthorName,
		req.Bio,
	)
	if err != nil {

		u.log.Error("err on db CreateAuth", logger.Error(err))
		return nil, err
	}

	resp, err := u.GetAuthById(context.Background(), &product_service.GetByIdReq{Id: id.String()})

	if err != nil {

		u.log.Error("err on db GetAuthById", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (u *AuthRepo) GetAuthById(ctx context.Context, req *product_service.GetByIdReq) (*product_service.Author, error) {

	var resp product_service.Author
	qury := `
		SELECT 
			*
		FROM 
			authors 
		WHERE
			author_id = $1
	`

	err := u.db.QueryRow(
		ctx,
		qury,
		req.Id,
	).Scan(
		&resp.AuthorId,
		&resp.AuthorName,
		&resp.Bio,
		&resp.CreatedAt,
		&resp.UpdatedAt,
		&resp.DeletedAt,
	)

	if err != nil {

		u.log.Error("err on db GetAuthById", logger.Error(err))
		return nil, err
	}

	return &resp, nil
}

func (u *AuthRepo) GetAuths(ctx context.Context, req *product_service.GetListReq) (*product_service.AuthorGetListResp, error) {

	offset := (req.Page - 1) * req.Limit


	var resp product_service.Author
	var res product_service.AuthorGetListResp
	qury := `
		SELECT 
			*
		FROM 
			authors 
		WHERE 
    		deleted_at IS NULL
		LIMIT $1 OFFSET $2;
			
	`

	row, err := u.db.Query(
		ctx,
		qury,
		req.Limit,
		offset,
	)

	if err != nil {

		u.log.Error("err on db GetAuths", logger.Error(err))
		return nil, err
	}

	for row.Next() {

		row.Scan(
			&resp.AuthorId,
			&resp.AuthorName,
			&resp.Bio,
			&resp.CreatedAt,
			&resp.UpdatedAt,
			&resp.DeletedAt,
		)

		res.Count++

		res.Author = append(res.Author, &resp)

	}

	return &res, nil
}

func (u *AuthRepo) UpdateAuth(ctx context.Context, req *product_service.AuthorUpdateReq) (*product_service.Author, error) {
	
	time := time.Now()
	
	query := `
		UPDATE
			authors
		SET	
				name = $1,
				bio = $2,
				updated_at = $3 
		WHERE 
				author_id = $4 
		AND  
				deleted_at is null
			`

	_, err := u.db.Exec(
		ctx,
		query,
		req.AuthorName,
		req.Bio,
		time,
		req.AuthorId,
	)
	if err != nil {

		u.log.Error("err on db UpdateAuth", logger.Error(err))
		return nil, err
	}

	resp, err := u.GetAuthById(context.Background(), &product_service.GetByIdReq{Id: req.AuthorId})

	if err != nil {

		u.log.Error("err on db GetAuthById", logger.Error(err))
		return nil, err
	}

	return resp, nil
}


func (u *AuthRepo) DeleteAuth(ctx context.Context, req *product_service.DeleteReq) (*product_service.Empty, error) {
	
	time := time.Now()
	
	query := `
		UPDATE
			authors
		SET	
				deleted_at = $1
		WHERE 
				author_id = $2
		AND  
				deleted_at is null
			`

	_, err := u.db.Exec(
		ctx,
		query,
		time,
		req.Id,
	)
	if err != nil {

		u.log.Error("err on db DeleteAuth", logger.Error(err))
		return nil, err
	}


	return &product_service.Empty{}, nil
}
