package api

import (
	"net/http"

	"github.com/rick-you/uvid/dtos"

	"github.com/labstack/echo/v4"
)

func bindDashApi(server Server) {
	api := &dashApi{server}
	rg := server.App.Group("/dash")

	rg.GET("/pvs", api.pageviews)
	rg.GET("/pvs/interval", api.pageviewInterval)

	rg.GET("/uvs/interval", api.uniqueVisitorInterval)

	rg.GET("/errors", api.errors)
	rg.GET("/errors/interval", api.errorInterval)

	rg.GET("/https/errors", api.httpErrorInterval)
	rg.GET("/https/errors/interval", api.httpErrorInterval)

	rg.GET("/sessions", api.sessions)
	rg.GET("/performances", api.avgPerformance)
	rg.GET("/events/group", api.eventGroup)

	rg.GET("/metric/count", api.metricCount)
}

type dashApi struct {
	Server
}

func (api *dashApi) pageviews(c echo.Context) error {
	body := &dtos.SpanFilterDTO{}
	if err := dtos.BindAndValidateDTO(c, body); err != nil {
		return err
	}

	pageviews, err := api.Dao.FindPageViews(api.Dao.SpanFilter(body.Start, body.End))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, pageviews)
}

func (api *dashApi) pageviewInterval(c echo.Context) error {
	body := &dtos.TimeIntervalSpanDTO{}
	if err := dtos.BindAndValidateDTO(c, body); err != nil {
		return err
	}

	interval, err := api.Dao.FindPageViewInterval(api.Dao.SpanFilter(body.Start, body.End), body.Unit)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, interval)
}

func (api *dashApi) metricCount(c echo.Context) error {
	body := &dtos.SpanFilterDTO{}
	if err := dtos.BindAndValidateDTO(c, body); err != nil {
		return err
	}
	spanFilter := api.Dao.SpanFilter(body.Start, body.End)

	Pv, err := api.Dao.FindPageViewCount(spanFilter)
	if err != nil {
		return err
	}

	Uv, err := api.Dao.FindUniqueVisitorCount(spanFilter)
	if err != nil {
		return err
	}

	HttpError, err := api.Dao.FindHTTPErrorCount(spanFilter)
	if err != nil {
		return err
	}

	JsError, err := api.Dao.FindJSErrorCount(spanFilter)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &dtos.CountDTO{
		Pv:        Pv,
		Uv:        Uv,
		JsError:   JsError,
		HttpError: HttpError,
	})
}

func (api *dashApi) uniqueVisitorInterval(c echo.Context) error {
	body := &dtos.TimeIntervalSpanDTO{}
	if err := dtos.BindAndValidateDTO(c, body); err != nil {
		return err
	}

	interval, err := api.Dao.FindUniqueVisitorInterval(api.Dao.SpanFilter(body.Start, body.End), body.Unit)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, interval)
}

func (api *dashApi) avgPerformance(c echo.Context) error {
	body := &dtos.SpanFilterDTO{}
	if err := dtos.BindAndValidateDTO(c, body); err != nil {
		return err
	}

	interval, err := api.Dao.FindAveragePerformanceInterval(api.Dao.SpanFilter(body.Start, body.End))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, interval)
}

func (api *dashApi) sessions(c echo.Context) error {
	body := &dtos.SpanFilterDTO{}
	if err := dtos.BindAndValidateDTO(c, body); err != nil {
		return err
	}

	sessions, err := api.Dao.FindSessions(api.Dao.SpanFilter(body.Start, body.End))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, sessions)
}

func (api *dashApi) eventGroup(c echo.Context) error {
	body := &dtos.SpanFilterDTO{}
	if err := dtos.BindAndValidateDTO(c, body); err != nil {
		return err
	}

	interval, err := api.Dao.FindEventInterval(api.Dao.SpanFilter(body.Start, body.End))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, interval)
}

func (api *dashApi) errors(c echo.Context) error {
	body := &dtos.SpanFilterDTO{}
	if err := dtos.BindAndValidateDTO(c, body); err != nil {
		return err
	}

	errors, err := api.Dao.FindJSErrors(api.Dao.SpanFilter(body.Start, body.End))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, errors)
}

func (api *dashApi) errorInterval(c echo.Context) error {
	body := &dtos.TimeIntervalSpanDTO{}
	if err := dtos.BindAndValidateDTO(c, body); err != nil {
		return err
	}

	interval, err := api.Dao.FindJSErrorInterval(api.Dao.SpanFilter(body.Start, body.End), body.Unit)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, interval)
}

func (api *dashApi) httpErrors(c echo.Context) error {
	body := &dtos.SpanFilterDTO{}
	if err := dtos.BindAndValidateDTO(c, body); err != nil {
		return err
	}

	errors, err := api.Dao.FindHTTPErrors(api.Dao.SpanFilter(body.Start, body.End))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, errors)
}

func (api *dashApi) httpErrorInterval(c echo.Context) error {
	body := &dtos.TimeIntervalSpanDTO{}
	if err := dtos.BindAndValidateDTO(c, body); err != nil {
		return err
	}

	interval, err := api.Dao.FindHTTPErrorInterval(api.Dao.SpanFilter(body.Start, body.End), body.Unit)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, interval)
}
