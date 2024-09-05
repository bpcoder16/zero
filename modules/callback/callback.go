package callback

import (
	"errors"
	"github.com/bpcoder16/zero/modules/callback/db"
	"github.com/bpcoder16/zero/mysql"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func AddCallbackTask(callbackFuncName, params string, retryCnt int, expectedAt time.Time) error {
	return AddCallbackTaskTransactionFunc(callbackFuncName, params, retryCnt, expectedAt)(mysql.MasterDB())
}

func AddCallbackTaskTransactionFunc(callbackFuncName, params string, retryCnt int, expectedAt time.Time) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		if retryCnt <= 0 || retryCnt > 100 {
			return errors.New("callback retryCnt is " + strconv.Itoa(retryCnt))
		}
		return tx.Create(&db.ZeroCallbackRecord{
			CallbackFuncName: callbackFuncName,
			Params:           params,
			RetryCnt:         retryCnt,
			Status:           &db.StatusCallbackPending,
			ExpectedAt:       &expectedAt,
		}).Error
	}
}
