package main_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/muonsoft/api-testing/assertjson"
	"github.com/ozontech/cute"
)

func Test_Single_1(t *testing.T) {
	cute.NewTestBuilder().
		Title("Single test with default T").
		Tag("single_test").
		Description("some_description").
		Parallel().
		Create().
		RequestRepeat(3).
		RequestBuilder(
			cute.WithMethod(http.MethodGet),
			cute.WithURI("https://jsonplaceholder.typicode.com/posts/1/comments"),
			cute.WithMarshalBody(struct {
				Name string `json:"name"`
			}{
				Name: "Vasya Pupkin",
			}),
		).
		ExpectExecuteTimeout(10*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBodyT(assertJSON(func(json *assertjson.AssertJSON) {
			json.Node().ForEach(func(node *assertjson.AssertNode) {
				node.Assert(func(json *assertjson.AssertJSON) {
					json.Node("id").IsInteger()
					json.Node("postId").IsInteger()
					json.Node("name").IsString().WithLengthGreaterThan(1)
					json.Node("email").IsEmail()
					json.Node("body").IsString().WithLengthGreaterThan(1)
				})
			})
		})).
		ExecuteTest(context.Background(), t)
}

func assertJSON(jsonAssert assertjson.JSONAssertFunc) func(t cute.T, body []byte) error {
	return func(t cute.T, body []byte) error {
		fmt.Println(string(body))

		assertjson.Has(TAdapter{T: t}, body, jsonAssert)
		return nil
	}
}

type TAdapter struct {
	cute.T
}

func (T TAdapter) Helper() {}
