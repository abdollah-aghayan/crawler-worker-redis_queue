package httprouter

import (
	"crawler/domain"
	"crawler/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CrawlerRequest crawler request interface
type CrawlerRequest interface {
	Fetch(c *gin.Context)
}

//crawlerRoute
type crawlerRoute struct {
	CrawlerRequest
}

var crawlerRoutes = newCrawlerRoute()

func newCrawlerRoute() CrawlerRequest {
	return &crawlerRoute{}
}

/**
 * Routes Method
 */

//getUserInfo get user info by id
func (u *crawlerRoute) Fetch(c *gin.Context) {

	req := domain.Request{}
	c.BindJSON(&req)

	err := req.ValidateRequest()

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	// call logic
	logic := logic.New()
	err = logic.FetchUrls(c, req)
	if err != nil {
		c.JSON(err.Code, err)
	}

	res := map[string]string{
		"success": "true",
	}

	c.JSON(http.StatusOK, res)
	return
}
