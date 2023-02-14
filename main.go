package main

import (
	"context"
	"flag"
	"fmt"
	"gin-simple-project/global/variable"
	"gin-simple-project/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis"

	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"
)

var env = flag.String("env", "test", "环境标识")

func init() {
	flag.Parse()
	viper.SetConfigName(*env)
	viper.AddConfigPath("./config") //添加配置文件所在的路径
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("配置文件读取失败: %s\n", err)
	}
	//models.InitDB()
	//initRedis()
}

func main() {
	gin.SetMode(viper.GetString("server.runMode"))
	endPoint := fmt.Sprintf(":%d", viper.GetInt("server.httpPort"))
	server := &http.Server{
		Addr:         endPoint,
		Handler:      routers.InitRouter(),
		ReadTimeout:  time.Second * 60,
		WriteTimeout: time.Second * 60,
	}
	go func() {
		// 服务连接
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listening: %s\n", err)
		}
	}()
	log.Printf("[info] start http server listening%s", endPoint)
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting...")
}

// 初始化连接
func initRedis() (err error) {
	variable.Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"), // no password set
		DB:       0,                                 // use default DB
	})

	_, err = variable.Rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
