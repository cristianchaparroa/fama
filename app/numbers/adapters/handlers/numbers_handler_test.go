package handlers

import (
	"encoding/json"
	"errors"
	"fama/core/handlers"
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
	expectedLang := "en"

	s.manager.Mock.On("ToWords", mock.Anything, mock.Anything).
		Return(expectedWords, nil).
		Once()

	s.handler.ToWords(c)
	s.Equal(http.StatusOK, recorder.Code)

	var response NumberToWordsResponse
	err := json.NewDecoder(recorder.Body).Decode(&response)
	s.Nil(err)

	s.Equal(StatusOK, response.Status)
	s.Equal(expectedWords, response.Words)
	s.Equal(expectedLang, response.Lang)
}

func (s *numbersHandlerSuite) TestNumbersHandler_ToWords_OK_With_Lang() {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	url := fmt.Sprintf("/numbers/%s/words?lang=fr", "12")
	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	c.Params = []gin.Param{
		{
			Key:   "number",
			Value: "12",
		},
	}
	expectedWords := "douze"
	expectedLang := "fr"

	s.manager.Mock.On("ToWords", mock.Anything, mock.Anything).
		Return(expectedWords, nil).
		Once()

	s.handler.ToWords(c)
	s.Equal(http.StatusOK, recorder.Code)

	var response NumberToWordsResponse
	err := json.NewDecoder(recorder.Body).Decode(&response)
	s.Nil(err)

	s.Equal(StatusOK, response.Status)
	s.Equal(expectedWords, response.Words)
	s.Equal(expectedLang, response.Lang)
}

func (s *numbersHandlerSuite) TestNumbersHandler_ToWords_WithoutSupportedLang() {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	url := fmt.Sprintf("/numbers/%s/words?lang=de", "12")
	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	c.Params = []gin.Param{
		{
			Key:   "number",
			Value: "12",
		},
	}
	expectedWords := ""
	expectedError := "lang not supported"
	s.manager.Mock.On("ToWords", mock.Anything, mock.Anything).
		Return(expectedWords, errors.New(expectedError)).
		Once()

	s.handler.ToWords(c)
	s.Equal(http.StatusUnprocessableEntity, recorder.Code)

	var errResponse handlers.Error
	err := json.NewDecoder(recorder.Body).Decode(&errResponse)
	s.Nil(err)

	s.Equal(http.StatusUnprocessableEntity, errResponse.Status)
	s.Equal(expectedError, errResponse.Title)
}
