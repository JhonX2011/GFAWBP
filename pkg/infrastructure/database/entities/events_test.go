package entities

import (
	"testing"

	"github.com/JhonX2011/GFAWBP/pkg/domain/models/internal_structs/entities"
	"github.com/JhonX2011/GFAWBP/test/doubles/database"
	gt "github.com/JhonX2011/GFAWBP/test/generic"
)

type eventsScenery struct {
	gt.GenericTest
	events Events
}

func givenEventsScenery(t *testing.T) *eventsScenery {
	t.Parallel()
	return &eventsScenery{
		events: Events{},
	}
}

func (s *eventsScenery) whenToDomainIsExecuted() {
	s.AResult = s.events.ToDomain()
}

func (s *eventsScenery) whenFromDomainIsExecuted() {
	s.events.FromDomain(database.GetDataEvents())
}

func TestToDomainIsExecuted(t *testing.T) {
	u := givenEventsScenery(t)
	u.whenToDomainIsExecuted()
	u.ThenIsType(t, &entities.Events{}, u.AResult)
	u.ThenNotEmpty(t)
}

func TestFromDomainIsExecuted(t *testing.T) {
	u := givenEventsScenery(t)
	u.whenFromDomainIsExecuted()
	u.ThenNoHaveError(t)
}
