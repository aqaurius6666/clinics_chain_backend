package psql

import (
	"context"

	"github.com/minh1611/clinics_chain_management/apiservice/src/internal/db/user"
)

func (u *ServerCDBRepo) InsertUser(ctx context.Context, value *user.User) (*user.User, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if err := u.Db.WithContext(ctx).Create(value).Error; err != nil {
		return nil, err
	}

	return value, nil
}
