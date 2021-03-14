package service

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/yigithanbalci/amazon-clone-account-service/dbclient"
	"github.com/yigithanbalci/amazon-clone-account-service/model"
)

var whenStr = "when the request is handled by the router"

func TestGetAccountForWrongPath(t *testing.T) {
	convey.Convey("given a http request for /invalid/123", t, func() {
		req := httptest.NewRequest("GET", "/invalid/123", nil)
		resp := httptest.NewRecorder()

		convey.Convey(whenStr, func() {
			NewRouter().ServeHTTP(resp, req)

			convey.Convey("then the response should be a 404", func() {
				convey.So(resp.Code, convey.ShouldEqual, 404)
			})
		})
	})
}

func TestGetAccount(t *testing.T) {
	mockRepo := &dbclient.MockBoltClient{}

	mockRepo.On("QueryAccount", "123").Return(model.Account{ID: "123", Name: "person_123"}, nil)
	mockRepo.On("QueryAccount", "456").Return(model.Account{}, fmt.Errorf("some error"))

	Dbclient = mockRepo

	convey.Convey("given a http request for /accounts/123", t, func() {
		req := httptest.NewRequest("GET", "/accounts/123", nil)
		resp := httptest.NewRecorder()

		convey.Convey(whenStr, func() {
			NewRouter().ServeHTTP(resp, req)

			convey.Convey("then the response should be a 200", func() {
				convey.So(resp.Code, convey.ShouldEqual, 200)

				account := model.Account{}
				json.Unmarshal(resp.Body.Bytes(), &account)
				convey.So(account.ID, convey.ShouldEqual, "123")
				convey.So(account.Name, convey.ShouldEqual, "person_123")
			})
		})
	})

	convey.Convey("given a http request for /accounts/456", t, func() {
		req := httptest.NewRequest("GET", "/accounts/456", nil)
		resp := httptest.NewRecorder()

		convey.Convey(whenStr, func() {
			NewRouter().ServeHTTP(resp, req)

			convey.Convey("then the response should be a 404", func() {
				convey.So(resp.Code, convey.ShouldEqual, 404)
			})
		})
	})
}
