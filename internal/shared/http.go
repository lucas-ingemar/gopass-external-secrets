package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

// FIXME: Here should errors be handeled
func HttpWriteResponse(ctx context.Context, w http.ResponseWriter, response Response, err error) error {
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		status := http.StatusInternalServerError
		rErr, ok := err.(ResponseError)
		if ok {
			status = rErr.HttpStatus()
		}
		// NOTE: ctx does not work here..
		log.Err(rErr).Int("status", status).Msg("error response")
		w.WriteHeader(status)
		_, wErr := fmt.Fprintf(w, "%d error", status)
		return wErr

	} else {
		w.WriteHeader(http.StatusOK)
	}
	return json.NewEncoder(w).Encode(&response)
}
