package app

import (
	"net/http"
	"dumper/conf"
	"github.com/pkg/errors"
)

func filtraToken(r *http.Request) error {
	if r.URL.Query().Get("token") != conf.Token() {
		return errors.New("token inválido")
	}
	return nil
}
