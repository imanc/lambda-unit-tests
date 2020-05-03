package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	t.Run("Unable to get IP", func(t *testing.T) {

		_, err := handler(context.TODO(), events.APIGatewayProxyRequest{})
		if err == nil {
			t.Fatal("Error")
		}
	})

}
