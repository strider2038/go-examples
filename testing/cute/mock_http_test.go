package main_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/muonsoft/api-testing/assertjson"
	"github.com/ozontech/cute"
)

type MockRoundTripper struct {
	handler http.Handler
}

func (m *MockRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	recorder := httptest.NewRecorder()
	m.handler.ServeHTTP(recorder, request)

	response := recorder.Result()
	response.Request = request

	return response, nil
}

var roundTripper = &MockRoundTripper{
	handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"key": "value"}`))
	}),
}

var mockBuilder = cute.NewHTTPTestMaker(
	cute.WithCustomHTTPRoundTripper(roundTripper),
)

func Test_MockHTTP(t *testing.T) {
	mockBuilder.NewTestBuilder().
		Title("Single test with default T").
		Tag("single_test").
		Description("some_description").
		Create().
		RequestBuilder(
			cute.WithMethod(http.MethodPost),
			cute.WithURI("/api/endpoint"),
		).
		ExpectStatus(http.StatusOK).
		AssertBodyT(assertJSON(func(json *assertjson.AssertJSON) {
			json.Node().Print()
			json.Node("key").IsString().EqualTo("value")
		})).
		ExecuteTest(context.Background(), t)
}
