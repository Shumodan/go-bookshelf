package main

import (
	"embed"
	_ "embed"

	"io/fs"
	"log"
	"net/http"

	"github.com/Shumodan/go-bookshelf/config"
	"github.com/Shumodan/go-bookshelf/logger"
	"github.com/Shumodan/go-bookshelf/middleware"
	"github.com/Shumodan/go-bookshelf/migration"
	"github.com/Shumodan/go-bookshelf/mycontext"
	"github.com/Shumodan/go-bookshelf/repository"
	"github.com/Shumodan/go-bookshelf/router"
	"github.com/labstack/echo/v4"
)

//go:embed zaplogger.develop.yml
var zfile []byte

//go:embed public/*
var embededFiles embed.FS

func getFileSystem() http.FileSystem {
	log.Print("using embed mode")
	fsys, err := fs.Sub(embededFiles, "public")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}

func main() {
	e := echo.New()

	conf, _ := config.Load()
	logger := logger.NewLogger(zfile)
	logger.GetZapLogger().Infof("Loaded configuration")

	rep := repository.NewBookRepository(logger, conf)
	context := mycontext.NewContext(rep, conf, logger)

	migration.CreateDatabase(context)
	migration.InitMasterData(context)

	router.Init(e, context)
	assetHandler := http.FileServer(getFileSystem())
	e.GET("/", echo.WrapHandler(assetHandler))
	e.GET("/*", echo.WrapHandler(assetHandler))
	logger.GetZapLogger().Infof("Served the static contents. ")

	middleware.InitLoggerMiddleware(e, context)
	middleware.InitSessionMiddleware(e, context)

	if err := e.Start(":8080"); err != nil {
		logger.GetZapLogger().Errorf(err.Error())
	}

	defer rep.Close()
}
