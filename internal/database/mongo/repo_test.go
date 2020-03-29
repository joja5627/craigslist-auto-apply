package mongo

import (
	"craigslist-auto-apply/internal/scrape"
	"github.com/google/uuid"
	"github.com/tj/assert"
	"testing"
)

//https://github.com/donvito/learngo/tree/master/mongo-microservice

func TestShouldSaveMongoEntity(t *testing.T) {
	repo, _ := NewMongoRepository()
	err := repo.Add(&scrape.Listing{
		ID:                  uuid.UUID{},
		StateCode:           "",
		Title:               "",
		Url:                 "",
		Query:               "",
		ContactInfoUrl:      "",
		ListingInfoResponse: "",
		EmailResponse:       "",
		Email:               "",
	})
	assert.NoError(t,err)
	listings := repo.GetAll()
	assert.True(t,len(listings) > 0)
}