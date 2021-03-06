package rest_test

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/untoldwind/routing/rest"
)

func TestHttpError(t *testing.T) {
	Convey("Given an InternalServerError", t, func() {
		httpErr := rest.InternalServerError(errors.New("Something is wrong"))

		Convey("When error is send as json", func() {
			response := httptest.NewRecorder()
			httpErr.Send(response, rest.JsonResponseEncoder)

			var received rest.HTTPError
			err := json.Unmarshal(response.Body.Bytes(), &received)

			So(err, ShouldBeNil)
			So(received.Code, ShouldEqual, 500)
			So(received.Type, ShouldEndWith, "/500")
			So(received.Details, ShouldEqual, "Something is wrong")
		})
	})
}
