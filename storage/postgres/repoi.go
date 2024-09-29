package postgres

import (
	"context"

	"github.com/jasurxaydarov/book_shop_produc_service/genproto/product_service"
)

type OrderedItemRepoI interface {
	CreateOrderedItem(ctx context.Context, req *product_service.OrderItemCreateReq) (*product_service.OrderItem, error)
	GetOrderedItemById(ctx context.Context, req *product_service.GetByIdReq) (*product_service.OrderItem, error)
	GetOrderedItemsByOrdreId(ctx context.Context, req *product_service.GetByIdReq) (*product_service.OrderItemGetListResp, error)
	GetOrderedItems(ctx context.Context, req *product_service.GetListReq) (*product_service.OrderItemGetListResp, error) 
	UpdateOrderedItem(ctx context.Context, req *product_service.OrderItemUpdate) (*product_service.OrderItem, error) 
	DeleteOrderedItem(ctx context.Context, req *product_service.DeleteReq) (*product_service.Empty, error)
}

type OrderRepoI interface {
	CreateOrder(ctx context.Context, req *product_service.OrderCreateReq) (*product_service.Order, error)
	GetOrderById(ctx context.Context, req *product_service.GetByIdReq) (*product_service.Order, error)
	GetOrders(ctx context.Context, req *product_service.GetListReq) (*product_service.OrderGetListResp, error)
	UpdateOrder(ctx context.Context, req *product_service.OrderUpdateReq) (*product_service.Order, error)
	DeleteOrder(ctx context.Context, req *product_service.DeleteReq) (*product_service.Empty, error)
}

type AuthRepoI interface {
	CreateAuth(ctx context.Context, req *product_service.AuthorCreateReq) (*product_service.Author, error)
	GetAuthById(ctx context.Context, req *product_service.GetByIdReq) (*product_service.Author, error)
	GetAuths(ctx context.Context, req *product_service.GetListReq) (*product_service.AuthorGetListResp, error)
	UpdateAuth(ctx context.Context, req *product_service.AuthorUpdateReq) (*product_service.Author, error)
	DeleteAuth(ctx context.Context, req *product_service.DeleteReq) (*product_service.Empty, error)
}

type BookRepoI interface {
	CreateBook(ctx context.Context, req *product_service.BookCreateReq) (*product_service.Book, error)
	GetBookById(ctx context.Context, req *product_service.GetByIdReq) (*product_service.Book, error)
	GetBooks(ctx context.Context, req *product_service.GetListReq) (*product_service.BookGetListResp, error)
	UpdateBook(ctx context.Context, req *product_service.BookUpdateReq) (*product_service.Book, error)
	DeleteBook(ctx context.Context, req *product_service.DeleteReq) (*product_service.Empty, error)
}

type CategoryRepoI interface {
	CreateCategory(ctx context.Context, req *product_service.CategoryCreateReq) (*product_service.Category, error)
	GetCategoryById(ctx context.Context, req *product_service.GetByIdReq) (*product_service.Category, error)
	GetCategories(ctx context.Context, req *product_service.GetListReq) (*product_service.CategoryGetListResp, error)
	UpdateCategory(ctx context.Context, req *product_service.CategoryUpdateReq) (*product_service.Category, error)
	DeleteCategory(ctx context.Context, req *product_service.DeleteReq) (*product_service.Empty, error)
}
