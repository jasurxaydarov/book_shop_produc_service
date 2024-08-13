package storage

import (
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/book_shop_produc_service/storage/postgres"
	"github.com/saidamir98/udevs_pkg/logger"
)

type StorageRepoI interface{
	GetOrderedItemRepo()postgres.OrderedItemRepoI
	GetCategoryRepo()postgres.CategoryRepoI
	GetAuthRepo()postgres.AuthRepoI
	GetBookRepo()postgres.BookRepoI
	GetOrderRepo()postgres.OrderRepoI


}

type storageRepo struct{
	ordeRepo postgres.OrderRepoI
	orderedItemRepo postgres.OrderedItemRepoI
	authRepo postgres.AuthRepoI
	bookRepo postgres.BookRepoI
	categoryRepo postgres.CategoryRepoI
	
}

func NewOrderedItemRepo(db *pgx.Conn,log logger.LoggerI)StorageRepoI{

	return &storageRepo{
		orderedItemRepo: postgres.NewOrderedItemRepo(db,log),
		ordeRepo: postgres.NewOrderRepo(db,log),
		categoryRepo: postgres.NewCategoryRepo(db,log),
		authRepo: postgres.NewAuthRepo(db,log),
		bookRepo: postgres.NewBookRepo(db,log),
	}
}


func (s *storageRepo)GetOrderedItemRepo()postgres.OrderedItemRepoI{

	return s.orderedItemRepo
}

func (s *storageRepo)GetAuthRepo()postgres.AuthRepoI{

	return s.authRepo
}

func (s *storageRepo)GetBookRepo()postgres.BookRepoI{

	return s.bookRepo
}

func (s *storageRepo)GetCategoryRepo()postgres.CategoryRepoI{

	return s.categoryRepo
}
func (s *storageRepo)GetOrderRepo()postgres.OrderRepoI{

	return s.ordeRepo
}