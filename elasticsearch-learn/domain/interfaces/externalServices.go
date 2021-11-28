package interfaces

import externalentity "github.com/MarceloZardoBR/go-api-frame/domain/external-entity"

type ElasticSearchService interface {
	GetByQuery(query string, index string) (retVal externalentity.ElasticSearchResponse, err error)
	GetByIndexAndFilterByField(index, query string, fields []string) (err error)
}
