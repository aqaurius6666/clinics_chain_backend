package lib

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minh1611/clinics_chain_management/apiservice/src/internal/utils/e"
	"golang.org/x/xerrors"
)

func SetBody(c *gin.Context) error {
	if c.Request.Method == http.MethodGet {
		return nil
	}
	buf := &bytes.Buffer{}

	_, err := buf.ReadFrom(c.Request.Body)
	if err != nil {
		return xerrors.Errorf("%w", e.ErrMissingBody)
	}
	c.Set("body", buf.Bytes())
	c.Request.Body = io.NopCloser(buf)

	return nil
}

func GetBody(g *gin.Context, target interface{}) error {
	var bbody []byte
	var err error
	ibody, ok := g.Get("body")
	if !ok {
		return xerrors.Errorf("%w", e.ErrMissingBody)
	} else {
		bbody = ibody.(json.RawMessage)
	}

	err = json.Unmarshal(bbody, target)
	if err != nil {
		return xerrors.Errorf("%w", e.ErrMissingBody)
	}
	return nil

}

func ToUUID(id interface{}) (uuid.UUID, error) {
	switch tmp := id.(type) {
	case uuid.UUID:
		return tmp, nil
	case string:
		tmp2, err := uuid.Parse(tmp)
		if err != nil {
			return uuid.Nil, e.ErrIdInvalidFormat
		}
		return tmp2, nil
	default:
		return uuid.Nil, e.ErrIdInvalidFormat
	}
}