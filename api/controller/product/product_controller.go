package product

import (
	"context"
	"net/http"

	ut "github.com/go-playground/universal-translator"
	"github.com/krobus00/technical-test-rest-api/api/service/product"
	"github.com/krobus00/technical-test-rest-api/infrastructure"
	"github.com/krobus00/technical-test-rest-api/model"
	"github.com/krobus00/technical-test-rest-api/util"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

const (
	tag = "[ProductController]"
)

type Controller struct {
	fx.In

	Logger         infrastructure.Logger
	ProductService product.ProductService
	Translator     *ut.UniversalTranslator
}

func (c *Controller) HandleCreateProduct(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()

	payload := new(model.CreateProductRequest)
	if err := eCtx.Bind(payload); err != nil {
		return err
	}

	if err := eCtx.Validate(payload); err != nil {
		trans := util.TranslatorFromRequestHeader(eCtx, c.Translator)
		return echo.NewHTTPError(http.StatusBadRequest, util.BuildValidationErrors(err, trans))
	}

	ctx = context.WithValue(ctx, "userID", eCtx.Get("userID").(string))
	resp, err := c.ProductService.Store(ctx, payload)
	if err != nil {
		return err
	}

	dataResp := model.DataResponse{
		Data: resp,
	}

	return eCtx.JSON(http.StatusOK, dataResp)
}

func (c *Controller) HandleGetProductDetail(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()

	payload := new(model.GetProductDetailRequest)
	if err := eCtx.Bind(payload); err != nil {
		return err
	}

	if err := eCtx.Validate(payload); err != nil {
		trans := util.TranslatorFromRequestHeader(eCtx, c.Translator)
		return echo.NewHTTPError(http.StatusBadRequest, util.BuildValidationErrors(err, trans))
	}

	ctx = context.WithValue(ctx, "userID", eCtx.Get("userID").(string))
	ctx = context.WithValue(ctx, "role", eCtx.Get("role").(string))
	resp, err := c.ProductService.FindProductByID(ctx, payload)
	if err != nil {
		return err
	}
	dataResp := model.DataResponse{
		Data: resp,
	}

	return eCtx.JSON(http.StatusOK, dataResp)
}

func (c *Controller) HandleGetAllProduct(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()

	payload := new(model.PaginationRequest)
	if err := eCtx.Bind(payload); err != nil {
		return err
	}

	if err := eCtx.Validate(payload); err != nil {
		trans := util.TranslatorFromRequestHeader(eCtx, c.Translator)
		return echo.NewHTTPError(http.StatusBadRequest, util.BuildValidationErrors(err, trans))
	}

	ctx = context.WithValue(ctx, "userID", eCtx.Get("userID").(string))
	ctx = context.WithValue(ctx, "role", eCtx.Get("role").(string))
	resp, err := c.ProductService.FindAll(ctx, payload)
	if err != nil {
		return err
	}

	dataResp := model.DataResponse{
		Data: resp,
	}

	return eCtx.JSON(http.StatusOK, dataResp)
}

func (c *Controller) HandleUpdateProduct(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()

	payload := new(model.UpdateProductRequest)
	if err := eCtx.Bind(payload); err != nil {
		return err
	}

	if err := eCtx.Validate(payload); err != nil {
		trans := util.TranslatorFromRequestHeader(eCtx, c.Translator)
		return echo.NewHTTPError(http.StatusBadRequest, util.BuildValidationErrors(err, trans))
	}

	ctx = context.WithValue(ctx, "userID", eCtx.Get("userID").(string))
	ctx = context.WithValue(ctx, "role", eCtx.Get("role").(string))
	resp, err := c.ProductService.Update(ctx, payload)
	if err != nil {
		return err
	}

	dataResp := model.DataResponse{
		Data: resp,
	}

	return eCtx.JSON(http.StatusOK, dataResp)
}

func (c *Controller) HandleDeleteProduct(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()

	payload := new(model.DeleteProductRequest)
	if err := eCtx.Bind(payload); err != nil {
		return err
	}

	if err := eCtx.Validate(payload); err != nil {
		trans := util.TranslatorFromRequestHeader(eCtx, c.Translator)
		return echo.NewHTTPError(http.StatusBadRequest, util.BuildValidationErrors(err, trans))
	}

	ctx = context.WithValue(ctx, "userID", eCtx.Get("userID").(string))
	ctx = context.WithValue(ctx, "role", eCtx.Get("role").(string))
	err := c.ProductService.Delete(ctx, payload)
	if err != nil {
		return err
	}
	resp := &model.BasicResponse{
		Message: "OK",
	}
	return eCtx.JSON(http.StatusOK, resp)
}
