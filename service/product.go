package service

import (
	"context"
	"fmt"

	"github.com/jasurxaydarov/book_shop_produc_service/genproto/product_service"
	"github.com/jasurxaydarov/book_shop_produc_service/storage"
	"github.com/saidamir98/udevs_pkg/logger"
)

type ProductService struct {
	storage storage.StorageRepoI
	product_service.UnimplementedProductServiceServer
}

func NewOrderedItemService(storage storage.StorageRepoI) *ProductService {

	return &ProductService{storage: storage}
}

func (o *ProductService) CreateOrdered_Item(ctx context.Context, req *product_service.OrderItemCreateReq) (*product_service.OrderItem, error) {

	resp, err := o.storage.GetOrderedItemRepo().CreateOrderedItem(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (o *ProductService) DeleteOrdered_Item(ctx context.Context,req *product_service.DeleteReq) (*product_service.Empty, error) {

	resp, err := o.storage.GetOrderedItemRepo().DeleteOrderedItem(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}
func (o *ProductService) GetOrdered_Item(ctx context.Context, req *product_service.GetByIdReq) (*product_service.OrderItem, error) {

	resp, err := o.storage.GetOrderedItemRepo().GetOrderedItemById(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}
func (o *ProductService) GetOrdered_Items(ctx context.Context,req *product_service.GetListReq) (*product_service.OrderItemGetListResp, error) {
	resp, err := o.storage.GetOrderedItemRepo().GetOrderedItems(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}
func (o *ProductService) UpdateOrdered_Item(ctx context.Context,req *product_service.OrderItemUpdate) (*product_service.OrderItem, error) {

	resp, err := o.storage.GetOrderedItemRepo().UpdateOrderedItem(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (o *ProductService) GetOrdered_ItemByOrderId(ctx context.Context, req *product_service.GetByIdReq) (*product_service.OrderItemGetListResp, error) {

	log := logger.NewLogger("", logger.LevelDebug)

	log.Debug("errrrrrrrrrrrrr")

	resp, err := o.storage.GetOrderedItemRepo().GetOrderedItemsByOrdreId(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (c *ProductService) CreateCategory(ctx context.Context, req *product_service.CategoryCreateReq) (*product_service.Category, error) {

	resp, err := c.storage.GetCategoryRepo().CreateCategory(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (c *ProductService) GetCategory(ctx context.Context, req *product_service.GetByIdReq) (*product_service.Category, error) {
	resp, err := c.storage.GetCategoryRepo().GetCategoryById(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (c *ProductService) GetCategories(ctx context.Context, req *product_service.GetListReq) (*product_service.CategoryGetListResp, error) {

	resp, err := c.storage.GetCategoryRepo().GetCategories(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (c *ProductService) UpdateCategory(ctx context.Context, req *product_service.CategoryUpdateReq) (*product_service.Category, error) {

	resp, err := c.storage.GetCategoryRepo().UpdateCategory(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}
func (c *ProductService) DeleteCategory(ctx context.Context, req *product_service.DeleteReq) (*product_service.Empty, error) {

	resp, err := c.storage.GetCategoryRepo().DeleteCategory(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (b *ProductService) CreateBook(ctx context.Context, req *product_service.BookCreateReq) (*product_service.Book, error) {

	resp, err := b.storage.GetBookRepo().CreateBook(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (b *ProductService) GetBook(ctx context.Context, req *product_service.GetByIdReq) (*product_service.Book, error) {

	resp, err := b.storage.GetBookRepo().GetBookById(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (b *ProductService) DeleteBook(ctx context.Context, req *product_service.DeleteReq) (*product_service.Empty, error) {
	resp, err := b.storage.GetBookRepo().DeleteBook(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (b *ProductService) GetBooks(ctx context.Context, req *product_service.GetListReq) (*product_service.BookGetListResp, error) {

	resp, err := b.storage.GetBookRepo().GetBooks(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}
func (b *ProductService) UpdateBook(ctx context.Context, req *product_service.BookUpdateReq) (*product_service.Book, error) {
	resp, err := b.storage.GetBookRepo().UpdateBook(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (a *ProductService) CreateAuth(ctx context.Context, req *product_service.AuthorCreateReq) (*product_service.Author, error) {

	resp, err := a.storage.GetAuthRepo().CreateAuth(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}
func (a *ProductService) DeleteAuth(ctx context.Context, req *product_service.DeleteReq) (*product_service.Empty, error) {
	resp, err := a.storage.GetAuthRepo().DeleteAuth(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}
func (a *ProductService) GetAuth(ctx context.Context, req *product_service.GetByIdReq) (*product_service.Author, error) {

	resp, err := a.storage.GetAuthRepo().GetAuthById(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil

}
func (a *ProductService) GetAuths(ctx context.Context, req *product_service.GetListReq) (*product_service.AuthorGetListResp, error) {
	resp, err := a.storage.GetAuthRepo().GetAuths(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}
func (a *ProductService) UpdateAuth(ctx context.Context, req *product_service.AuthorUpdateReq) (*product_service.Author, error) {
	resp, err := a.storage.GetAuthRepo().UpdateAuth(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (o *ProductService) CreateOrder(ctx context.Context, req *product_service.OrderCreateReq) (*product_service.Order, error) {

	log := logger.NewLogger("", logger.LevelDebug)

	log.Debug("errrrrrrrrrrrrr")

	resp, err := o.storage.GetOrderRepo().CreateOrder(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (o *ProductService) GetOrder(ctx context.Context, req *product_service.GetByIdReq) (*product_service.Order, error) {

	resp, err := o.storage.GetOrderRepo().GetOrderById(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (o *ProductService) DeleteOrder(ctx context.Context, req *product_service.DeleteReq) (*product_service.Empty, error) {
	resp, err := o.storage.GetOrderRepo().DeleteOrder(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (o *ProductService) Getorders(ctx context.Context, req *product_service.GetListReq) (*product_service.OrderGetListResp, error) {

	resp, err := o.storage.GetOrderRepo().GetOrders(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}
func (o *ProductService) Updateorder(ctx context.Context, req *product_service.OrderUpdateReq) (*product_service.Order, error) {
	resp, err := o.storage.GetOrderRepo().UpdateOrder(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}
