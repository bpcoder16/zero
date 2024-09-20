package captcha

import (
	"context"
	"fmt"
	"github.com/bpcoder16/zero/contrib/captcha"
	"github.com/bpcoder16/zero/core/utils"
	"github.com/bpcoder16/zero/redis"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

const (
	zeroCaptchaRedisKey = "zero:captcha:%s"
)

func generateRandCode(ctx context.Context, n int) (uuidStr, randCodeShow string) {
	randCodeShow = utils.RandStr(n)
	randCodeLower := strings.ToLower(randCodeShow)

	uuidStr = uuid.New().String()
	redis.DefaultClient().SetEx(ctx, fmt.Sprintf(zeroCaptchaRedisKey, uuidStr), randCodeLower, time.Minute)
	return
}

func GetImageDefaultGinHandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuidStr, randCodeShow := generateRandCode(c, 5)
		imageBytes := captcha.ImageBytes(200, 100, randCodeShow)
		c.Writer.Header().Set("Content-Type", "image/png")
		c.Writer.Header().Set("X-UID", uuidStr)
		_, _ = c.Writer.Write(imageBytes)
	}
}

func CheckDefaultGinHandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		xUid := c.DefaultQuery("xUid", "")
		code := c.DefaultQuery("code", "")
		c.JSON(http.StatusOK, gin.H{
			"xUid":  xUid,
			"code":  code,
			"check": Check(c, xUid, code),
		})
	}
}

func Check(ctx context.Context, uuidStr, randCode string) bool {
	redisKey := fmt.Sprintf(zeroCaptchaRedisKey, uuidStr)
	redisRandCode, err := redis.DefaultClient().Get(ctx, redisKey).Result()
	if err != nil {
		return false
	}
	redis.DefaultClient().Del(ctx, redisKey)
	return redisRandCode == strings.ToLower(randCode)
}
