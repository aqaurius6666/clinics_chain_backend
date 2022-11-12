package reservation_comment

import "golang.org/x/xerrors"

var (
	prefix        = "reservation_comment"
	ErrNotFound   = xerrors.Errorf("%s: record not found", prefix)
	ErrInsertFail = xerrors.Errorf("%s: insert failed", prefix)
)
