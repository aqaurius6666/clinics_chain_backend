package revisitation

import "golang.org/x/xerrors"

var (
	prefix        = "revisitation"
	ErrNotFound   = xerrors.Errorf("%s: record not found", prefix)
	ErrInsertFail = xerrors.Errorf("%s: insert failed", prefix)
)
