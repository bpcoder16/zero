package mysql

import (
	"github.com/bpcoder16/zero/contrib/orm"
	"github.com/bpcoder16/zero/core/log"
	"gorm.io/gorm"
)

var defaultMySQLGormDBManager *orm.GormDBManager

func SetManager(configPath string, logger *log.Helper) {
	defaultMySQLGormDBManager = orm.NewGormDBManager(configPath, logger)
}

func MasterDB() *gorm.DB {
	return defaultMySQLGormDBManager.MasterDB()
}

func SlaveDB() *gorm.DB {
	return defaultMySQLGormDBManager.SlaveDB()
}
