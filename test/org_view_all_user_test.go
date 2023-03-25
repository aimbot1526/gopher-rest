package test

import (
	"gopher-rest/pkg/routes"
	"gopher-rest/pkg/utils"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestOrgViewAllUser(t *testing.T) {

	token, tokenErr := utils.GenerateNewAccessToken()

	if tokenErr != nil {
		t.Fatal(tokenErr)
	}

	tests := []struct {
		description   string
		route         string
		method        string
		token         string
		body          io.Reader
		expectedError bool
		expectedCode  int
		testType      string
	}{
		{
			description:   "List Org users.",
			route:         "/api/v1/all/3",
			method:        "GET",
			body:          nil,
			expectedError: false,
			expectedCode:  200,
			testType:      "Success",
		},
		{
			description:   "List Org users without orgId.",
			route:         "/api/v1/all",
			method:        "GET",
			body:          nil,
			expectedError: false,
			expectedCode:  404,
			testType:      "Failure",
		},
	}

	app := fiber.New()
	routes.PrivateRoutes(app)

	for _, test := range tests {

		req := httptest.NewRequest(test.method, test.route, test.body)

		req.Header.Set("Authorization", token)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)

		assert.Equalf(t, false, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}

}
