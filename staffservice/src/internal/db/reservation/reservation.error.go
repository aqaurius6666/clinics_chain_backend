package reservation

import "golang.org/x/xerrors"

var (
	prefix        = "reservation"
	ErrNotFound   = xerrors.Errorf("%s: record not found", prefix)
	ErrInsertFail = xerrors.Errorf("%s: insert failed", prefix)
)
