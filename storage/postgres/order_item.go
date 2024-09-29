package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/book_shop_produc_service/genproto/product_service"
	"github.com/saidamir98/udevs_pkg/logger"
)

type orderedItemRepo struct {
	db  *pgx.Conn
	log logger.LoggerI
}

func NewOrderedItemRepo(db *pgx.Conn, log logger.LoggerI) OrderedItemRepoI {

	return &orderedItemRepo{db: db, log: log}
}

func (o *orderedItemRepo) CreateOrderedItem(ctx context.Context, req *product_service.OrderItemCreateReq) (*product_service.OrderItem, error) {
	id := uuid.New()

	query := `
		INSERT INTO
			order_items(
				order_item_id,
				order_id,
				user_id,
				book_id,
				quantity,
				price 
			)VALUES(
				$1,$2,$3,$4,$5,$6
			)
			`

	_, err := o.db.Exec(
		ctx,
		query,
		id,
		req.OrderId,
		req.UserId,
		req.BookId,
		req.Quantity,
		req.Price,
	)
	if err != nil {

		o.log.Error("err on db CreateOrderItem", logger.Error(err))
		return nil, err
	}

	resp, err := o.GetOrderedItemById(context.Background(), &product_service.GetByIdReq{Id: id.String()})

	if err != nil {

		o.log.Error("err on db GetOrderItemById", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (o *orderedItemRepo) GetOrderedItemById(ctx context.Context, req *product_service.GetByIdReq) (*product_service.OrderItem, error) {

	var resp product_service.OrderItem
	qury := `
		SELECT 
			order_item_id,
			order_id,
			user_id,
			book_id,
			quantity,
			price 
		FROM 
			order_items
		WHERE
			order_item_id = $1
	`

	err := o.db.QueryRow(
		ctx,
		qury,
		req.Id,
	).Scan(
		&resp.OrderItemId,
		&resp.OrderId,
		&resp.UserId,
		&resp.BookId,
		&resp.Quantity,
		&resp.Price,
	)

	if err != nil {

		o.log.Error("err on db GetOrderItemByOrderId scan", logger.Error(err))
		return nil, err
	}

	return &resp, nil
}

func (o *orderedItemRepo) GetOrderedItemsByOrdreId(ctx context.Context, req *product_service.GetByIdReq) (*product_service.OrderItemGetListResp, error) {

	var row product_service.OrderItemGetListResp
	var resp product_service.OrderItem

	qury := `
		SELECT 
			*
		FROM 
			order_items
		WHERE
			order_id = $1
	`

	rows, err := o.db.Query(ctx, qury, req.Id)

	for rows.Next() {

		rows.Scan(
			&resp.OrderItemId,
			&resp.OrderId,
			&resp.UserId,
			&resp.BookId,
			&resp.Quantity,
			&resp.Price,
			&resp.CreatedAt,
			&resp.UpdatedAt,
			&resp.DeletedAt,
		)

		row.OrderItem = append(row.OrderItem, &resp)

		row.Count++
	}

	if err != nil {

		o.log.Error("err on db GetOrderItemByOrderId", logger.Error(err))
		return nil, err
	}

	return &row, nil
}

func (o *orderedItemRepo) GetOrderedItems(ctx context.Context, req *product_service.GetListReq) (*product_service.OrderItemGetListResp, error) {

	offset := (req.Page - 1) * req.Limit

	var row product_service.OrderItemGetListResp
	var resp product_service.OrderItem

	qury := `
		SELECT 
			*
		FROM 
			order_items
		WHERE 
    		deleted_at IS NULL
		LIMIT $1 OFFSET $2;
	`

	rows, err := o.db.Query(ctx, qury, req.Limit,offset)

	for rows.Next() {

		rows.Scan(
			&resp.OrderItemId,
			&resp.OrderId,
			&resp.UserId,
			&resp.BookId,
			&resp.Quantity,
			&resp.Price,
			&resp.CreatedAt,
			&resp.UpdatedAt,
			&resp.DeletedAt,
		)

		row.Count++

		row.OrderItem = append(row.OrderItem, &resp)

	}

	if err != nil {

		o.log.Error("err on db GetOrderItems", logger.Error(err))
		return nil, err
	}

	return &row, nil
}


func (o *orderedItemRepo) UpdateOrderedItem(ctx context.Context, req *product_service.OrderItemUpdate) (*product_service.OrderItem, error) {

	time :=time.Now()

	query := `
			UPDATE
				order_items
			SET
				order_id = $1,
				user_id = $2,
				book_id = $3,
				quantity = $4,
				price = $5,
				update_at = $6
		WHERE 
				order_item_id = $7,
		AND  
				deleted_at is null
			`

	_, err := o.db.Exec(
		ctx,
		query,
		req.OrderId,
		req.UserId,
		req.BookId,
		req.Quantity,
		req.Price,
		time,
		req.OrderItemId,
	)
	if err != nil {

		o.log.Error("err on db UpdateOrderItem", logger.Error(err))
		return nil, err
	}

	resp, err := o.GetOrderedItemById(context.Background(), &product_service.GetByIdReq{Id: req.OrderItemId})

	if err != nil {

		o.log.Error("err on db GetOrderItemById", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (o *orderedItemRepo) DeleteOrderedItem(ctx context.Context, req *product_service.DeleteReq) (*product_service.Empty, error) {

	time :=time.Now()

	query := `
			UPDATE
				order_items
			SET
				deleted_at = $1,
		WHERE 
				order_item_id = $2,
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

		o.log.Error("err on db DeleteOrderItem", logger.Error(err))
		return nil, err
	}

	

	return &product_service.Empty{}, nil
}