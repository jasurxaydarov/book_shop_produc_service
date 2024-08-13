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


func (o *ProductService) CreateOrdered_Item(ctx context.Context, req *product_service.OrderItemCreateReq) (*product_service.OrderItem, error){
	log:=logger.NewLogger("",logger.LevelDebug)

	log.Debug("errrrrrrrrrrrrr")
	
	resp, err := o.storage.GetOrderedItemRepo().CreateOrderedItem(ctx, req)
	
	if err != nil {
	
		fmt.Println(err)
		return nil, err
	}
	
	return resp, nil
}

func (o *ProductService) DeleteOrdered_Item(context.Context, *product_service.DeleteReq) (*product_service.Empty, error){

	return nil,nil
}
func (o *ProductService) GetOrdered_Item(ctx context.Context,req *product_service.GetByIdReq) (*product_service.OrderItem, error){
	log:=logger.NewLogger("",logger.LevelDebug)

	log.Debug("errrrrrrrrrrrrr")
	
	resp, err := o.storage.GetOrderedItemRepo().GetOrderedItemById(ctx, req)
	
	if err != nil {
	
		fmt.Println(err)
		return nil, err
	}
	
	return resp, nil
}
func (o *ProductService) GetOrdered_Items(context.Context, *product_service.GetListReq) (*product_service.OrderItemGetListResp, error){

	return nil,nil
}
func (o *ProductService)UpdateOrdered_Item(context.Context, *product_service.OrderItemCreateReq) (*product_service.OrderItem, error){

	return nil,nil
}

func (o *ProductService)GetOrdered_ItemByOrderId(ctx context.Context,req *product_service.GetByIdReq) (*product_service.OrderItemGetListResp, error){

	log:=logger.NewLogger("",logger.LevelDebug)

	log.Debug("errrrrrrrrrrrrr")
	
	resp, err := o.storage.GetOrderedItemRepo().GetOrderedItemByOrdreId(ctx, req)
	
	if err != nil {
	
		fmt.Println(err)
		return nil, err
	}
	
	return resp, nil
}



func (c *ProductService)CreateCategory(ctx context.Context,req *product_service.CategoryCreateReq) (*product_service.Category, error){

	log:=logger.NewLogger("",logger.LevelDebug)
	log.Debug("errrrrrrrrr")
	resp, err := c.storage.GetCategoryRepo().CreateCategory(ctx,req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (c *ProductService) GetCategory(ctx context.Context,req *product_service.GetByIdReq) (*product_service.Category, error){
	resp, err := c.storage.GetCategoryRepo().GetCategoryById(ctx,req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (c *ProductService) GetCategories(context.Context, *product_service.GetListReq) (*product_service.CategoryGetListResp, error){
	
	return nil,nil
}

func (c *ProductService) UpdateCategory(context.Context, *product_service.CategoryUpdateReq) (*product_service.Category, error){
	
	return nil,nil
}
func (c *ProductService) DeleteCategory(context.Context, *product_service.DeleteReq) (*product_service.Empty, error){

	return nil,nil
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

func (b *ProductService) DeleteBook(context.Context, *product_service.DeleteReq) (*product_service.Empty, error) {

	return nil, nil
}

func (b *ProductService) GetBooks(context.Context, *product_service.GetListReq) (*product_service.BookGetListResp, error) {

	return nil, nil
}
func (b *ProductService) UpdateBook(context.Context, *product_service.BookUpdateReq) (*product_service.Book, error) {

	return nil, nil
}




func (a *ProductService) CreateAuth(ctx context.Context, req *product_service.AuthorUpdateReq) (*product_service.Author, error) {

	resp, err := a.storage.GetAuthRepo().CreateAuth(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}
func (a *ProductService) DeleteAuth(ctx context.Context, req *product_service.DeleteReq) (*product_service.Empty, error) {

	return nil, nil
}
func (a *ProductService) GetAuth(ctx context.Context, req *product_service.GetByIdReq) (*product_service.Author, error) {

	resp, err := a.storage.GetAuthRepo().GetAuthById(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil

}
func (a *ProductService) GetAuths(context.Context, *product_service.GetListReq) (*product_service.AuthorGetListResp, error) {

	return nil, nil
}
func (a *ProductService) UpdateAuth(context.Context, *product_service.AuthorUpdateReq) (*product_service.Author, error) {

	return nil, nil
}


func (o *ProductService) CreateOrder(ctx context.Context, req *product_service.OrderCreateReq) (*product_service.Order, error) {

	log:=logger.NewLogger("",logger.LevelDebug)

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

func (o *ProductService) DeleteOrder(context.Context, *product_service.DeleteReq) (*product_service.Empty, error) {

	return nil, nil
}

func (o *ProductService) Getorders(context.Context, *product_service.GetListReq) (*product_service.OrderGetListResp, error) {

	return nil, nil
}
func (o *ProductService) Updateorder(context.Context, *product_service.OrderUpdateReq) (*product_service.Order, error) {

	return nil, nil
}

