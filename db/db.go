package db

import "context"

const DBNAME = "command-constructor"

type Dropper interface {
	Drop(context.Context) error
}
