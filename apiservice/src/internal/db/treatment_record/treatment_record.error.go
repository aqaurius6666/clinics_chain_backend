package treatment_record

import "golang.org/x/xerrors"

var (
	prefix        = "treatment_record"
	ErrNotFound   = xerrors.Errorf("%s: record not found", prefix)
	ErrInsertFail = xerrors.Errorf("%s: insert failed", prefix)
)
