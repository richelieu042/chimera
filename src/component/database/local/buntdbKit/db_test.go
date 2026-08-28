package buntdbKit

import (
	"testing"

	"github.com/richelieu042/chimera/v3/src/idKit"
	"github.com/richelieu042/chimera/v3/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/buntdb"
)

func TestOpen(t *testing.T) {
	logrusKit.MustSetUp(nil)

	// Open the data.db file. It will be created if it doesn't exist.
	db, err := Open("_test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	key := "key"
	value := idKit.NewUlid()

	// (1) To set a value you must open a read/write transaction:
	err = db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, value, nil)
		return err
	})
	if err != nil {
		panic(err)
	}
	logrus.Info("Manager to set.")

	// (2) To get the value
	var val string
	err = db.View(func(tx *buntdb.Tx) error {
		var err error
		val, err = tx.Get(key)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	logrus.WithField("val", val).Info("Manager to get.")
}
