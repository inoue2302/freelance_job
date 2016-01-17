package base

import (
	"github.com/revel/revel"
	"net/http"
)

/**
 * ApiBaseController
 * Api基底コントローラー
 */
type ApiBaseController struct {
	*revel.Controller
}

/**
 * ErrorResponse
 * エラーレスポンス構造体
 */
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

/**
 * Response
 * レスポンス構造体
 */
type Response struct {
	Results interface{} `json:"results"`
}

/**
 * HandleBadRequestError
 * 不正パラメーターハンドラ
 *
 * @param text string
 * @return json
 */
func (c *ApiBaseController) HandleBadRequestError(text string) revel.Result {
	c.Response.Status = http.StatusBadRequest
	r := ErrorResponse{c.Response.Status, text}
	return c.RenderJson(r)
}

/**
 * HandleNotFoundError
 * 404ハンドラ
 *
 * @param text string
 * @return json
 */
func (c *ApiBaseController) HandleNotFoundError(text string) revel.Result {
	c.Response.Status = http.StatusNotFound
	r := ErrorResponse{c.Response.Status, text}
	return c.RenderJson(r)
}

/**
 * HandleInternalServerError
 * 500ハンドラ
 *
 * @param text string
 * @return json
 */
func (c *ApiBaseController) HandleInternalServerError(text string) revel.Result {
	c.Response.Status = http.StatusInternalServerError
	r := ErrorResponse{c.Response.Status, text}
	return c.RenderJson(r)
}
