package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/imanc/lambda-unit-testing/internal/pkg/service"
	"github.com/stretchr/testify/assert"
)

type ServiceMock struct {
}

// CurrentMonth returns the current Month
func (s ServiceMock) CurrentMonth() service.Month {

	return service.Month{
		Name: "May",
	}
}

func TestHandler(t *testing.T) {
	expectedResponse := "\"{\\\"Name\\\":\\\"May\\\"}\""
	handlerClosure := handlerClosure(ServiceMock{})
	lambdaHandler := lambda.NewHandler(handlerClosure)
	response, err := lambdaHandler.Invoke(context.TODO(), nil)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, string(response))
}
