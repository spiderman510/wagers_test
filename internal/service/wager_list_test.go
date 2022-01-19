package service

import (
	"net/http"
	"testing"
	"wagers/internal/entity"
	"wagers/internal/errors"

	"github.com/gogo/status"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
)

var DatabaseError = status.Error(codes.Internal, "database error")

func TestListWagerService(t *testing.T) {
	t.Run("Should return PageRequiredError if page <= 0", func(t *testing.T) {
		s := CreateWagerServiceTestContext()
		_, err := s.WagerService.ListWager(&entity.ListRequest{
			Page: -1,
		})
		assert.Equal(t, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.PageRequiredError,
		}, err)
	})
	t.Run("Should return LimitRequiredError if limit <= 0", func(t *testing.T) {
		s := CreateWagerServiceTestContext()
		_, err := s.WagerService.ListWager(&entity.ListRequest{
			Page:  1,
			Limit: -1,
		})
		assert.Equal(t, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.LimitRequiredError,
		}, err)
	})
	t.Run("Should return ListWagerError if query list is failed to execute", func(t *testing.T) {
		s := CreateWagerServiceTestContext()
		request := &entity.ListRequest{
			Page:  1,
			Limit: 10,
		}
		s.WagerRepository.Mock.On("Query", (request.Page-1)*request.Limit, request.Limit).Return(nil, DatabaseError)
		_, err := s.WagerService.ListWager(request)
		assert.Equal(t, &errors.Errors{
			Code:    http.StatusInternalServerError,
			Message: errors.ListWagerError,
		}, err)
	})

	t.Run("Should return list wagers", func(t *testing.T) {
		s := CreateWagerServiceTestContext()
		request := &entity.ListRequest{
			Page:  1,
			Limit: 10,
		}
		wagers := []*entity.Wager{
			CreateWagerTestData(),
		}
		s.WagerRepository.Mock.On("Query", (request.Page-1)*request.Limit, request.Limit).Return(wagers, nil)
		actualWagers, err := s.WagerService.ListWager(request)
		assert.Nil(t, err)
		assert.Equal(t, wagers, actualWagers)
	})
}
