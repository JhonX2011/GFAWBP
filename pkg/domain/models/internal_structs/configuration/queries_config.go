package configuration

import "fmt"

type QueriesConfig struct {
	QueryID    string `json:"query_id"`
	QueryValue string `json:"query_value"`
}

type Queries struct {
	Queries []QueriesConfig `json:"queries"`
}

func (q Queries) Get(queryID string) (QueriesConfig, error) {
	for _, query := range q.Queries {
		if query.QueryID == queryID {
			return query, nil
		}
	}

	return QueriesConfig{}, fmt.Errorf("service not found [%s]", queryID)
}
