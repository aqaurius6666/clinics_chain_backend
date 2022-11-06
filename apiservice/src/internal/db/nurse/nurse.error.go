package nurse

import "golang.org/x/xerrors"

var (
	prefix        = "nurse"
	ErrNotFound   = xerrors.Errorf("%s: record not found", prefix)
	ErrInsertFail = xerrors.Errorf("%s: insert failed", prefix)
)
