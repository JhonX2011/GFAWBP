package entities

import (
	"testing"

	"github.com/JhonX2011/GFAWBP/pkg/domain/models/internal_structs/entities"
	"github.com/JhonX2011/GFAWBP/test/doubles/database"
	gt "github.com/JhonX2011/GFAWBP/test/generic"
)

type movableScenery struct {
	gt.GenericTest
	movables Movables
}

func givenMovableScenery(t *testing.T) *movableScenery {
	t.Parallel()
	return &movableScenery{
		movables: Movables{
			Inventories: make([]Inventories, 2),
		},
	}
}

func (s *movableScenery) whenToDomainMovablesIsExecuted() {
	s.AResult = s.movables.ToDomain()
}

func (s *movableScenery) whenFromDomainMovablesIsExecuted() {
	s.movables.FromDomain(database.GetDataDomainMovables())
}

func TestToDomainMovablesIsExecuted(t *testing.T) {
	u := givenMovableScenery(t)
	u.whenToDomainMovablesIsExecuted()
	u.ThenIsType(t, &entities.Movables{}, u.AResult)
	u.ThenNotEmpty(t)
}

func TestFromDomainMovablesIsExecuted(t *testing.T) {
	u := givenMovableScenery(t)
	u.whenFromDomainMovablesIsExecuted()
	u.ThenNoHaveError(t)
}
