package staff_working_time

import "golang.org/x/xerrors"

var (
	prefix        = "staff_working_time"
	ErrNotFound   = xerrors.Errorf("%s: record not found", prefix)
	ErrInsertFail = xerrors.Errorf("%s: insert failed", prefix)
)
