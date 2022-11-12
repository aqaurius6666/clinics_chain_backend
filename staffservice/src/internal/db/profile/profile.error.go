package profile

import "golang.org/x/xerrors"

var (
	prefix        = "profile"
	ErrNotFound   = xerrors.Errorf("%s: record not found", prefix)
	ErrInsertFail = xerrors.Errorf("%s: insert failed", prefix)
)
