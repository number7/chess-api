package service_test

import (
	"errors"
	"testing"

	data "github.com/number7/chess-api/cmd/chess-api/data-layer/structs"
	"github.com/number7/chess-api/cmd/chess-api/service-layer"
	servicelayerfakes "github.com/number7/chess-api/cmd/chess-api/service-layer/service-layerfakes"
	"github.com/number7/chess-api/cmd/chess-api/service-layer/structs"
	"github.com/stretchr/testify/assert"
)

//go:generate counterfeiter . GetRepo

func TestGet(t *testing.T) {
	t.Run("returns and error when repo returns an error", func(t *testing.T) {
		fakeRepo := new(servicelayerfakes.FakeGetRepo)
		fakeRepo.GetReturns(data.GameData{}, errors.New("boom; head shot"))

		testObject := service.New(fakeRepo)

		_, err := testObject.Get(structs.GetRequest{})

		assert.Error(t, err)
	})
}
