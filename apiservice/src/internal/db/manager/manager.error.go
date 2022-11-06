package manager

import "golang.org/x/xerrors"

var (
	prefix        = "manager"
	ErrNotFound   = xerrors.Errorf("%s: record not found", prefix)
	ErrInsertFail = xerrors.Errorf("%s: insert failed", prefix)
)
