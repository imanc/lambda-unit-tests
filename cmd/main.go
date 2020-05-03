package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/imanc/lambda-unit-testing/infrastructure/pkg/service"
)

func handlerClosure(monthService service.MonthService) func(ctx context.Context) (string, error) {
	return func(ctx context.Context) (string, error) {
		currentMonth := monthService.CurrentMonth()
		response, err := json.Marshal(currentMonth)
		if err != nil {
			return "", err
		}

		return string(response), nil
	}
}
func main() {
	monthService := service.MonthService{}

	lambda.Start(handlerClosure(monthService))
}
