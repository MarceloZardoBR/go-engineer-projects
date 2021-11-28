package titlessearchengine

import (
	"net/http"

	externalentity "github.com/MarceloZardoBR/go-api-frame/domain/external-entity"
	"github.com/MarceloZardoBR/go-api-frame/domain/interfaces"
	"github.com/MarceloZardoBR/go-api-frame/server/viewmodels"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	elasticSearchService interfaces.ElasticSearchService
}

func NewController(elasticSearchService interfaces.ElasticSearchService) *Controller {
	return &Controller{
		elasticSearchService: elasticSearchService,
	}
}

func (c *Controller) SearchIndexWithQuery(ctx *fiber.Ctx) error {

	indexValue := ctx.Params("origin")
	queryValue := ctx.Query("query")

	if indexValue == "" && queryValue == "" {
		return ctx.Status(http.StatusUnprocessableEntity).SendString("Search values can't be empty")
	}

	resultSet, err := c.elasticSearchService.GetByQuery(queryValue, indexValue)
	if err != nil {
		// Jogar erro
	}

	titles := parseResultSetToViewModel(resultSet)

	return ctx.JSON(titles)
}

func parseResultSetToViewModel(retVal externalentity.ElasticSearchResponse) (titleList viewmodels.Titles) {
	for _, source := range retVal.Hits.HitsValues {
		title := viewmodels.Title(source.Source)
		titleList.Titles = append(titleList.Titles, title)
	}

	return titleList
}
