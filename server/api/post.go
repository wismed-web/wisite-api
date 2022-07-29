package api

import (
	"github.com/labstack/echo/v4"
	"github.com/wismed-web/wisite-api/server/api/post"
)

// register to main echo Group

// "/api/post"
func PostHandler(e *echo.Group) {

	var mGET = map[string]echo.HandlerFunc{
		"/template":            post.Template,
		"/ids":                 post.IdBatch,
		"/ids-all":             post.IdAll,
		"/one":                 post.GetOne,
		"/own/ids":             post.IdOwn,
		"/follower/ids":        post.Followers,
		"/thumbsup/status/:id": post.ThumbsUpStatus,
	}

	var mPOST = map[string]echo.HandlerFunc{
		"/upload": post.Upload,
	}

	var mPUT = map[string]echo.HandlerFunc{}

	var mDELETE = map[string]echo.HandlerFunc{
		"/del/one":   post.DelOne,
		"/erase/one": post.EraseOne,
	}

	var mPATCH = map[string]echo.HandlerFunc{
		"/thumbsup/:id": post.ThumbsUp,
	}

	// ------------------------------------------------------- //

	methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH"}

	mRegAPIs := map[string]map[string]echo.HandlerFunc{
		"GET":    mGET,
		"POST":   mPOST,
		"PUT":    mPUT,
		"DELETE": mDELETE,
		"PATCH":  mPATCH,
		// others...
	}

	mRegMethod := map[string]func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route{
		"GET":    e.GET,
		"POST":   e.POST,
		"PUT":    e.PUT,
		"DELETE": e.DELETE,
		"PATCH":  e.PATCH,
		// others...
	}

	for _, m := range methods {
		mAPI, method := mRegAPIs[m], mRegMethod[m]
		for path, handler := range mAPI {
			if handler == nil {
				continue
			}
			method(path, handler)
		}
	}
}
