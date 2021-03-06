package cat

import (
	"net/http"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"github.com/catcurd/models"
	"github.com/catcurd/services"
)

type Handler struct {
	Services services.Services
}

func (h Handler) Get(ctx *gofr.Context) (interface{}, error) {
	resp, err := h.Services.Get(ctx)

	if err != nil {
		return nil, errors.DB{Err: errors.Error("Internal DB error")}
	}

	res := models.Response{
		Cat:        resp,
		Message:    "Success",
		StatusCode: http.StatusOK,
	}

	return res, nil
}

func (h Handler) Create(ctx *gofr.Context) (interface{}, error) {
	var c models.Cat

	if err := ctx.Bind(&c); err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	if c.ID == "0" {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	resp, err := h.Services.Create(ctx, c)
	if err != nil {
		return nil, errors.Error("Internal DB error")
	}

	res := models.Response{
		Cat:        resp,
		Message:    "created successful",
		StatusCode: http.StatusCreated,
	}

	return res, nil
}

func (h Handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	if id == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	_, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.InvalidParam{
			Param: []string{"id"},
		}
	}

	res, err := h.Services.GetByID(ctx, id)
	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "cat",
			ID:     id,
		}
	}

	resp := models.Response{
		Cat:        res,
		Message:    "successful",
		StatusCode: http.StatusOK,
	}

	return resp, nil
}

func (h Handler) Delete(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	if id == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	_, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.InvalidParam{
			Param: []string{"id"},
		}
	}

	if err := h.Services.Delete(ctx, id); err != nil {
		return nil, errors.DB{Err: err}
	}

	res := models.Response{
		Cat:        nil,
		Message:    "successful",
		StatusCode: http.StatusOK,
	}

	return res, nil
}

func (h Handler) Update(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	_, err := strconv.Atoi(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	var c models.Cat

	if err = ctx.Bind(&c); err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	if c.ID == "0" {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	resp, err := h.Services.Update(ctx, c)
	if err != nil {
		return nil, errors.Error("Internal DB error")
	}

	res := models.Response{
		Cat:        resp,
		Message:    "update successful",
		StatusCode: http.StatusCreated,
	}

	return res, nil
}
