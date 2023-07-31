package main

import (
	"admin/core"
	_ "admin/docs"
	"admin/migration"
	"admin/route"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	if core.DB != nil {
		db, _ := core.DB.DB()
		defer db.Close()
	}

	migration.DBMigrate(core.DB)
	migration.DBSeed(core.DB)

	core.InitLogger()

	r := route.InitRouter()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", core.Config.ServerPort),
		Handler:      r,
		ReadTimeout:  time.Duration(core.Config.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(core.Config.WriteTimeout) * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			core.Logger.Error(fmt.Sprintf("listen: %s", err))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	core.Logger.Info("Server shutdown ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		core.Logger.Error(fmt.Sprintf("Server shutdown: %s\n", err))
	}
	core.Logger.Info("Server exiting")
}
