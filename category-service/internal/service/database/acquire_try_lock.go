package database

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type lockType int

const (
	_ lockType = iota
	LockTypeTask
)

func AcquireTryLock(ctx context.Context, tx *sqlx.Tx, lockID lockType, bar int32) (bool, error) {
	var isAcquired bool
	err := tx.GetContext(ctx, &isAcquired, fmt.Sprintf("select pg_try_advisory_xact_lock(%d, %d)", lockID, bar))
	return isAcquired, err
}
