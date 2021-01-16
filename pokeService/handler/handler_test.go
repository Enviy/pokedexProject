package handler

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"

	httpMocks "github.com/Enviy/pokedexProject/pokeService/.gen/mocks/http"

	"go.uber.org/zap"
)

type mockClients struct {
	mockResponseWriter *httpMocks.MockResponseWriter
}

func setup(t *testing.T) (Handler, mockClients) {
	handler := New(zap.S())

	ctrl := gomock.NewController(t)
	mockRespWriter := httpMocks.NewMockResponseWriter(ctrl)

	return handler, mockClients{
		mockResponseWriter: mockRespWriter,
	}
}

func TestHandler_Root(t *testing.T) {
	tests := []struct {
		Name        string
		WriteStatus int
	}{
		{
			Name:        "Success",
			WriteStatus: http.StatusOK,
		},
	}

	for i := 0; i < len(tests); i++ {
		currTest := tests[i]
		t.Run(currTest.Name, func(t *testing.T) {
			handler, mockClients := setup(t)

			mockClients.mockResponseWriter.EXPECT().WriteHeader(currTest.WriteStatus).Times(1)

			handler.Root(mockClients.mockResponseWriter, &http.Request{})
		})
	}
}
