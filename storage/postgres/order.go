package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/book_shop_produc_service/genproto/product_service"
	"github.com/saidamir98/udevs_pkg/logger"
)

type orderRepo struct {
	db  *pgx.Conn
	log logger.LoggerI
}

func NewOrderRepo(db *pgx.Conn, log logger.LoggerI) OrderRepoI {

	return &orderRepo{db: db, log: log}
}

func (o *orderRepo) CreateOrder(ctx context.Context, req *product_service.OrderCreateReq) (*product_service.Order, error) {
	id := uuid.New()

	query := `
		INSERT INTO
			orders(
				order_id,
				user_id,
				total_amount,
				order_status 
			)VALUES(
				$1,$2,$3,$4
			)
			`

	_, err := o.db.Exec(
		ctx,
		query,
		id,
		req.UserId,
		req.TotalAmount,
		req.OrderStatus,
	)
	if err != nil {

		o.log.Error("err on db CreateOrder", logger.Error(err))
		return nil, err
	}

	resp, err := o.GetOrderById(context.Background(), &product_service.GetByIdReq{Id: id.String()})

	if err != nil {

		o.log.Error("err on db GetOrderById", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (o *orderRepo) GetOrderById(ctx context.Context, req *product_service.GetByIdReq) (*product_service.Order, error) {

	var resp product_service.Order
	qury := `
		SELECT 
			order_id,
			user_id,
			total_amount,
			order_status
		FROM 
			orders
		WHERE
			order_id = $1
	`

	err := o.db.QueryRow(
		ctx,
		qury,
		req.Id,
	).Scan(
		&resp.OrderId,
		&resp.UserId,
		&resp.TotalAmount,
		&resp.OrderStatus,
	
	)

	if err != nil {

		o.log.Error("err on db GetOrderById", logger.Error(err))
		return nil, err
	}

	return &resp, nil
}


func (o *orderRepo) GetOrders(ctx context.Context, req *product_service.GetListReq) (*product_service.OrderGetListResp, error) {

	offset := (req.Page - 1) * req.Limit


	var resp product_service.Order

	var res product_service.OrderGetListResp
	qury := `
		SELECT 
			order_id,
			user_id,
			total_amount,
			order_status
		FROM 
			orders
		WHERE 
    		deleted_at IS NULL
		LIMIT $1 OFFSET $2;
	`

	row, err := o.db.Query(
		ctx,
		qury,
		req.Limit,
		offset,
	)

	fmt.Println("ssssssssssssssssssssssss")
	if err != nil {

		o.log.Error("err on db GetOrders", logger.Error(err))
		return nil, err
	}

	for row.Next() {

		row.Scan(
			&resp.OrderId,
			&resp.UserId,
			&resp.TotalAmount,
			&resp.OrderStatus,
		)
		if err != nil {

			o.log.Error("err on db GetOrders", logger.Error(err))
			return nil, err
		}
		res.Count++
		res.Order = append(res.Order, &resp)

	}

	return &res, nil
}


func (o *orderRepo) UpdateOrder(ctx context.Context, req *product_service.OrderUpdateReq) (*product_service.Order, error) {


	req.UpdatedAt = time.Now().String()

	query := `
			UPDATE
				orders
			SET
				total_amount = $1,
				order_status = $2,
				updated_at = $3
			WHERE 
				order_id = $4
			AND  
				deleted_at is null
			`

	_, err := o.db.Exec(
		ctx,
		query,
		req.TotalAmount,
		req.OrderStatus,
		req.UpdatedAt,
		req.OrderId,
	)
	if err != nil {

		o.log.Error("err on db UpdateOrder", logger.Error(err))
		return nil, err
	}

	resp, err := o.GetOrderById(context.Background(), &product_service.GetByIdReq{Id: req.OrderId})

	if err != nil {

		o.log.Error("err on db GetOrderById", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (o *orderRepo) DeleteOrder(ctx context.Context, req *product_service.DeleteReq) (*product_service.Empty, error) {


	time:=time.Now()

	query := `
			UPDATE
				orders
			SET
				deleted_at = $1
			WHERE 
				order_id = $2 
			AND  
				deleted_at is null
			`

	_, err := o.db.Exec(
		ctx,
		query,
		time,
		req.Id,
	)
	if err != nil {

		o.log.Error("err on db UpdateOrder", logger.Error(err))
		return nil, err
	}

	return &product_service.Empty{}, nil
}