package handlers

import (
	"encoding/json"
	"fama/numbers/ports/mocks"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type numbersHandlerSuite struct {
	suite.Suite
	manager *mocks.NumbersManager
	handler *NumbersHandler
}

func TestNumbersHandlerSuiteInit(t *testing.T) {
	suite.Run(t, new(numbersHandlerSuite))
}

func (s *numbersHandlerSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
	s.manager = new(mocks.NumbersManager)
	s.handler = newNumbersHandler(s.manager)
}

func (s *numbersHandlerSuite) TestNumbersHandler_ToWords_OK() {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	url := fmt.Sprintf("/numbers/%s/words", "12")
	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	c.Params = []gin.Param{{
		Key:   "number",
		Value: "12",
	}}

	expectedWords := "twelve"
	s.manager.Mock.On("ToWords", mock.Anything, mock.Anything).
		Return(expectedWords, nil).
		Once()

	s.handler.ToWords(c)
	s.Equal(http.StatusOK, recorder.Code)

	var response NumberToWordsResponse
	err := json.NewDecoder(recorder.Body).Decode(&response)
	s.Nil(err)

	expectedLang := "en"
	s.Equal(StatusOK, response.Status)
	s.Equal(expectedWords, response.Words)
	s.Equal(expectedLang, response.Lang)
}
