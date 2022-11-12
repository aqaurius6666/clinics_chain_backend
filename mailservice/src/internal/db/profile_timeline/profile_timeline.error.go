package profile_timeline

import "golang.org/x/xerrors"

var (
	prefix        = "profile_timeline"
	ErrNotFound   = xerrors.Errorf("%s: record not found", prefix)
	ErrInsertFail = xerrors.Errorf("%s: insert failed", prefix)
)
