package lib

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
)

func GetBody(g *gin.Context, target interface{}) error {
	jsonData, err := ioutil.ReadAll(g.Request.Body)
	if err != nil {
		return xerrors.Errorf("%w", "Cannot read the body")
	}
	err = json.Unmarshal(jsonData, target)
	if err != nil {
		return xerrors.Errorf("%w", "Body wrong format")
	}
	return nil
}
