package upload

import (
	"github.com/bpcoder16/zero/contrib/aliyun/oss"
	"github.com/gin-gonic/gin"
	"net/http"
)

// MultiAliyunDefaultGinHandlerFunc 批量上传默认方法
func MultiAliyunDefaultGinHandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从表单中获取多个文件
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve form data"})
			return
		}

		// 获取表单中的所有文件 (key 为 "files")
		files := form.File["files"]

		// 遍历所有文件
		imageURLs := make([]string, 0, len(files))
		for _, file := range files {
			var ossPath string
			ossPath, err = oss.SimpleUpload(file, "tmp/tmp")
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"err": err,
				})
				return
			}
			imageURLs = append(imageURLs, ossPath)
		}

		c.JSON(http.StatusOK, gin.H{
			"imageURLs": imageURLs,
		})
		return
	}
}
