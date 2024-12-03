package handlers

import (
	"library/internal/config"
	"library/internal/storage"
	"log/slog"

	_ "library/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run() {
	cfg, err := config.New()
	if err != nil {
		slog.Error("ERROR", slog.Any("err", err))
	}

	storage, err := storage.New(cfg)
	if err != nil {
		slog.Error("ERROR", slog.Any("err", err))
	}

	library := &Library{Storage: storage}

	router := gin.Default()

	// Вместо SelectApiFunc вставьте функцию, реализующую апи из 2-го пункта ТЗ
	router.GET("/info/:group/:song", SelectApiFunc)
	router.GET("/library/:group/:song", library.ReturnSongWithName)
	router.GET("/library", library.ReturnLibrary)

	router.POST("/songs", library.SaveSong)

	router.DELETE("/songs/:id", library.DeleteSong)
	router.DELETE("/library", library.DeleteDb)

	router.PATCH("/songs/:id/:group/:song", library.UpdateSong)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(":8080"); err != nil {
		slog.Error("ERROR", slog.Any("err", err))
	}
	slog.Info("Server run")

}

func SelectApiFunc(c *gin.Context) {

}
