package externalservices

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	externalentity "github.com/MarceloZardoBR/go-api-frame/domain/external-entity"
	"github.com/MarceloZardoBR/go-api-frame/domain/interfaces"
	"github.com/MarceloZardoBR/go-api-frame/infra/config"
	"github.com/MarceloZardoBR/go-api-frame/infra/utils"
	"github.com/pkg/errors"
)

type elasticSearchService struct {
	cfg *config.Config
}

// NewElasticSearchService returns a service instance
func NewElasticSearchService(cfg *config.Config) interfaces.ElasticSearchService {
	return &elasticSearchService{
		cfg: cfg,
	}
}

func (s *elasticSearchService) GetByQuery(query string, index string) (retVal externalentity.ElasticSearchResponse, err error) {
	endpoint := fmt.Sprintf("%s/%s/_search", s.cfg.ESEndpoint, index)

	queryParams := map[string]string{
		"pretty":      "true",
		"filter_path": "hits.hits._source",
	}

	var bodyValues = externalentity.ElasticSearchBody{
		Query: externalentity.Query{
			MultiMatch: externalentity.MultiMatch{
				Query:  query,
				Fields: []string{"*"},
			},
		},
	}

	reqTimeout := time.Second * time.Duration(s.cfg.ESRequestTimeout)

	response, _, err := utils.RequestAPIGETWithBody(endpoint, reqTimeout, bodyValues, queryParams)
	if err != nil {
		return retVal, errors.Wrap(err, "elasticsearch.GetByQuery error")
	}

	err = json.Unmarshal(response, &retVal)
	if err != nil {
		return retVal, errors.WithStack(err)
	}

	return retVal, nil
}

func (s *elasticSearchService) GetByIndexAndFilterByField(index, query string, fields []string) (err error) {
	if len(fields) == 0 {
		err = errors.Errorf("%s: %s", http.StatusText(http.StatusBadRequest), "search fields can't be null")
		return errors.Wrap(err, "")
	}

	//Create the field validates at controller.

	return nil
}
