package controller

import (
	"exchangeapp/global"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func LikeArticle(ctx *gin.Context) {
	articleId := ctx.Param("id")

	likeKey := "article:" + articleId + ":likes"

	if err := global.RedisDB.Incr(likeKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {"error" : err})
		return ;
	}
	ctx.JSON(http.StatusOK, gin.H{"message" : "successful"})

}

func GetArticleLikes(ctx *gin.Context) {
	articleId := ctx.Param("id")

	likeKey := "article:" + articleId + ":likes"

	likes, err := global.RedisDB.Get(likeKey).Result()

	if err == redis.Nil {
		likes = "0"
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H {"error" : err})
	}
	ctx.JSON(http.StatusOK, gin.H{"likes" : likes})
}