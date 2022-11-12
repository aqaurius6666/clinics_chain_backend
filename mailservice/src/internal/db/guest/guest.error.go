package guest

import "golang.org/x/xerrors"

var (
	prefix        = "guest"
	ErrNotFound   = xerrors.Errorf("%s: record not found", prefix)
	ErrInsertFail = xerrors.Errorf("%s: insert failed", prefix)
)
