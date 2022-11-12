package patient

import "golang.org/x/xerrors"

var (
	prefix        = "patient"
	ErrNotFound   = xerrors.Errorf("%s: record not found", prefix)
	ErrInsertFail = xerrors.Errorf("%s: insert failed", prefix)
)
