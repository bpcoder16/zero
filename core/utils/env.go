package utils

import (
	"os"
	"path/filepath"
	"runtime"
)

func RootPath() string {
	rootPath, _ := os.Getwd()
	return rootPath
}

func ZeroRootPath() string {
	// 获取当前文件的路径
	_, currentFile, _, _ := runtime.Caller(0)
	return filepath.Dir(filepath.Dir(filepath.Dir(currentFile)))
}
