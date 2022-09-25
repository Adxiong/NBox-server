/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-07-31 23:05:56
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-09-24 22:19:07
 */
package bootstrap

import (
	"context"
	"log"
	. "nbox/src/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"

	"github.com/gin-gonic/gin"
)

func Start(ctx context.Context) {
	MustInit(ctx)

	r := gin.Default()

	// 初始化基于redis的存储引擎
	// 参数说明：
	//    第1个参数 - redis最大的空闲连接数
	//    第2个参数 - 数通信协议tcp或者udp
	//    第3个参数 - redis地址, 格式，host:port
	//    第4个参数 - redis密码
	//    第5个参数 - session加密密钥
	store, errRedis := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte("secret"))
	if errRedis != nil {
		log.Fatalln(errRedis)
	}
	r.Use(sessions.Sessions("session", store))

	// 注册路由
	RegisterController(ctx, r)

	if err := r.Run(":8000"); err != nil {
		log.Printf("startup service failed, err:%v\n", err)
	}
}
