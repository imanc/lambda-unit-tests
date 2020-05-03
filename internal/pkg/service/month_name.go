package service

import "time"

// Service interface
type Service interface {
	CurrentMonth() Month
}

//Month groups some data about current month
type Month struct {
	Name string `json="name"`
}

// MonthService implements Service interface
type MonthService struct {
}

// NewMonthService return a new MonthService
func NewMonthService() MonthService {
	return MonthService{}
}

// CurrentMonth returns the current Month
func (s MonthService) CurrentMonth() Month {
	currentTime := time.Now()

	return Month{
		Name: currentTime.Month().String(),
	}
}
