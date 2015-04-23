package routing

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMethod(t *testing.T) {
	Convey("Given GET metcher", t, func() {
		var handlerCalled bool
		mockHandler := func(resp http.ResponseWriter, req *http.Request) {
			handlerCalled = true
		}
		matcher := GETFunc(mockHandler)

		Convey("When GET request is matched", func() {
			request, _ := http.NewRequest("GET", "/match", nil)
			recorder := httptest.NewRecorder()

			result := matcher("", recorder, request)

			So(result, ShouldBeTrue)
			So(handlerCalled, ShouldBeTrue)
		})

		Convey("When non-GET request is matched", func() {
			request, _ := http.NewRequest("POST", "/notmatch", nil)
			recorder := httptest.NewRecorder()

			result := matcher("", recorder, request)

			So(result, ShouldBeFalse)
			So(handlerCalled, ShouldBeFalse)
		})
	})

	Convey("Given POST metcher", t, func() {
		var handlerCalled bool
		mockHandler := func(resp http.ResponseWriter, req *http.Request) {
			handlerCalled = true
		}
		matcher := POSTFunc(mockHandler)

		Convey("When POST request is matched", func() {
			request, _ := http.NewRequest("POST", "/match", nil)
			recorder := httptest.NewRecorder()

			result := matcher("", recorder, request)

			So(result, ShouldBeTrue)
			So(handlerCalled, ShouldBeTrue)
		})

		Convey("When non-POST request is matched", func() {
			request, _ := http.NewRequest("PUT", "/match", nil)
			recorder := httptest.NewRecorder()

			result := matcher("", recorder, request)

			So(result, ShouldBeFalse)
			So(handlerCalled, ShouldBeFalse)
		})
	})

	Convey("Given PUT metcher", t, func() {
		var handlerCalled bool
		mockHandler := func(resp http.ResponseWriter, req *http.Request) {
			handlerCalled = true
		}
		matcher := PUTFunc(mockHandler)

		Convey("When PUT request is matched", func() {
			request, _ := http.NewRequest("PUT", "/match", nil)
			recorder := httptest.NewRecorder()

			result := matcher("", recorder, request)

			So(result, ShouldBeTrue)
			So(handlerCalled, ShouldBeTrue)
		})

		Convey("When non-PUT request is matched", func() {
			request, _ := http.NewRequest("DELETE", "/match", nil)
			recorder := httptest.NewRecorder()

			result := matcher("", recorder, request)

			So(result, ShouldBeFalse)
			So(handlerCalled, ShouldBeFalse)
		})
	})

	Convey("Given DELETE metcher", t, func() {
		var handlerCalled bool
		mockHandler := func(resp http.ResponseWriter, req *http.Request) {
			handlerCalled = true
		}
		matcher := DELETEFunc(mockHandler)

		Convey("When DELETE request is matched", func() {
			request, _ := http.NewRequest("DELETE", "/match", nil)
			recorder := httptest.NewRecorder()

			result := matcher("", recorder, request)

			So(result, ShouldBeTrue)
			So(handlerCalled, ShouldBeTrue)
		})

		Convey("When non-DELETE request is matched", func() {
			request, _ := http.NewRequest("GET", "/match", nil)
			recorder := httptest.NewRecorder()

			result := matcher("", recorder, request)

			So(result, ShouldBeFalse)
			So(handlerCalled, ShouldBeFalse)
		})
	})
}
