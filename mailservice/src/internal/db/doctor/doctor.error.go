package doctor

import "golang.org/x/xerrors"

var (
	prefix        = "doctor"
	ErrNotFound   = xerrors.Errorf("%s: record not found", prefix)
	ErrInsertFail = xerrors.Errorf("%s: insert failed", prefix)
)

