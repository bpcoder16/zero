package db

import "time"

type ZeroCallbackRecord struct {
	Id               int        `gorm:"column:id"`
	CallbackFuncName string     `gorm:"column:callback_func_name"`
	Params           string     `gorm:"column:params"`
	RetryCnt         int        `gorm:"column:retry_cnt"`
	Status           *int       `gorm:"column:status"`
	Remark           string     `gorm:"column:remark"`
	ExpectedAt       *time.Time `gorm:"column:expected_at"`
}

var (
	StatusCallbackException = -1
	StatusCallbackPending   = 0
	StatusCallbackSuccess   = 1
)
