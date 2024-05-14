package db

import (
	"database/sql"
)

// trong day can co tat ca chuc nang cuar queries va chuc nang Transfer
type Store interface {
	Querier //lam cho store cos tat ca chuc nang
	// TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB //dai dien cho database
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
// 	tx, err := store.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}

// 	q := New(tx)
// 	err = fn(q)
// 	if err != nil {
// 		if rbErr := tx.Rollback(); rbErr != nil {
// 			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
// 		}
// 		return err
// 	}
// 	return tx.Commit()
// }

// // TransferTxParams contains the input parameters of the transfer transaction
// // type TransferTxParams struct {
// // 	FromAccountID int64 `json:"from_account_id"`
// // 	ToAccountID   int64 `json:"to_account_id"`
// // 	Amount        int64 `json:"amount"`
// // }

// // TransferTxResult is the result of the tranfer transaction
// // type TransferTxResult struct {
// // 	Transfer    Transfer `json:"transfer"`
// // 	FromAccount Account  `json:"from_account"` //Tai khoan tu sau khi so du duoc cap nhat
// // 	ToAccount   Account  `json:"to_account"`   //Tai kho	an den tu sau khi so du duoc cap nhat
// // 	FromEntry   Entry    `json:"from_entry"`   //Muc nhap cua tai khoan tu ghi lai so tien do dang chuyen ra ngoai
// // 	ToEntry     Entry    `json:"to_entry"`     //muc nhao cua tai khoan ghi lai so tien do dang chuyen vao
// // }

// // it crates a transfer record, add account entries and update accounts balance within a single database transaction
// // func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
// // 	var result TransferTxResult

// // 	err := store.execTx(ctx, func(q *Queries) error {
// // 		var err error

// // 		return nil
// // 	})
// // 	return result, err
// // }

// //type SQLStore struct {
// //	connPool *pgxpool.Pool
// //	*Queries
// //}
// //
// //func NewStore(connPool *pgxpool.Pool) Store {
// //	return &SQLStore{
// //		connPool: connPool,
// //		Queries:  New(connPool),
// //	}
// //}
