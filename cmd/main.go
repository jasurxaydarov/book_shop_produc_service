package main

import (
	"context"
	"fmt"
	"net"

	"github.com/jasurxaydarov/book_shop_produc_service/genproto/product_service"
	"github.com/jasurxaydarov/book_shop_produc_service/pkg/db"
	"github.com/jasurxaydarov/book_shop_produc_service/service"
	"github.com/jasurxaydarov/book_shop_produc_service/storage"
	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc"
)

func main() {

	log := logger.NewLogger("", logger.LevelDebug)

	pgxConn, err := db.ConnDB(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(pgxConn)

	storage := storage.NewOrderedItemRepo(pgxConn, log)

	service := service.NewOrderedItemService(storage)

	listen, err := net.Listen("tcp", "localhost:8001")

	if err != nil {
		fmt.Println(err)
		return
	}

	server := grpc.NewServer()

	product_service.RegisterProductServiceServer(server, service)

	log.Debug("server serve on :8001")

	if err = server.Serve(listen); err != nil {
		log.Error(err.Error())
		return

	}

}
