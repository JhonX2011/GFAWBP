package entities

import (
	"testing"

	"github.com/JhonX2011/GFAWBP/pkg/domain/models/internal_structs/database/entities"
	"github.com/JhonX2011/GFAWBP/test/doubles/database"
	gt "github.com/JhonX2011/GFAWBP/test/generic"
)

type inventoriesScenery struct {
	gt.GenericTest
	inventories Inventories
}

func givenInventoriesScenery(t *testing.T) *inventoriesScenery {
	t.Parallel()
	return &inventoriesScenery{
		inventories: Inventories{},
	}
}

func (s *inventoriesScenery) whenToDomainInventoriesIsExecuted() {
	s.AResult = s.inventories.ToDomain()
}

func (s *inventoriesScenery) whenFromDomainInventoriesIsExecuted() {
	s.inventories.FromDomain(database.GetDataInventories())
}

func TestToDomainInventoriesIsExecuted(t *testing.T) {
	u := givenInventoriesScenery(t)
	u.whenToDomainInventoriesIsExecuted()
	u.ThenIsType(t, &entities.Inventories{}, u.AResult)
	u.ThenNotEmpty(t)
}

func TestFromDomainInventoriesIsExecuted(t *testing.T) {
	u := givenInventoriesScenery(t)
	u.whenFromDomainInventoriesIsExecuted()
	u.ThenNoHaveError(t)
}
