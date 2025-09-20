package main

import (
	"context"
	"errors"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (app *application) fetchBoxIdFromParams(ctx context.Context) (uint64, error) {
	var id uint64

	boxId := chi.URLParamFromCtx(ctx, "boxId")
	if boxId == "" {
		// failed to fetch boxId from URL parameters
		return id, errors.New("Failed to fetch boxId parameter from request context")
	}

	id, err := strconv.ParseUint(boxId, 10, 64)
	if err != nil {
		// failed to convert the boxId URL parameter to an unsigned integer, invalid parameter
		return id, err
	}

	return id, nil
}
