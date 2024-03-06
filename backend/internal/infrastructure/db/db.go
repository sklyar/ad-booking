package db

import (
	"github.com/sklyar/go-transact"
	"github.com/sklyar/go-transact/txsql"
)

// TxManager is an alias for transact.Manager.
type TxManager = transact.Manager

// Handler is an alias for txsql.DB.
type Handler = txsql.DB

// Row is an alias for txsql.Row.
type Row = txsql.Row
