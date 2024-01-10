package transaction

import (
	"GinTemplate/common/errs"
	"GinTemplate/pkg/connections/database"
	"GinTemplate/pkg/connections/database/gorms"
)

type Transaction struct {
	conn database.DbConn
}

func NewTransaction() *Transaction {
	return &Transaction{
		conn: gorms.NewTransaction(),
	}
}

func (t *Transaction) Action(f func(conn database.DbConn) *errs.BError) *errs.BError {
	t.conn.Begin()
	err := f(t.conn)
	if err != nil {
		t.conn.Rollback()
		return err
	}
	t.conn.Commit()
	return nil
}
