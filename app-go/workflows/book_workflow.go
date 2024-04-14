package workflows

import (
	"fmt"
	"temporal-sagas/activities"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type BookVacationInput struct {
	BookUserId   string `json:"book_user_id"`
	BookCarId    string `json:"book_car_id"`
	BookHotelId  string `json:"book_hotel_id"`
	BookFlightId string `json:"book_flight_id"`
	Attempts     int    `json:"attempts"`
}

func BookWorkflow(ctx workflow.Context, input BookVacationInput) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Book workflow started")

	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    1 * time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    30 * time.Second,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, activityOptions)

	var car string
	err := workflow.ExecuteActivity(ctx, activities.BookCar, input.BookCarId).Get(ctx, &car)
	if err != nil {
		return "", err
	}

	logger.Info("Sleeping for 1 second...")
	workflow.Sleep(ctx, 1*time.Second)

	var hotel string
	err = workflow.ExecuteActivity(ctx, activities.BookHotel, input.BookHotelId).Get(ctx, &hotel)
	if err != nil {
		return "", err
	}

	var flight string
	err = workflow.ExecuteActivity(ctx, activities.BookFlight, input.BookFlightId).Get(ctx, &flight)
	if err != nil {
		return "", err
	}

	output := fmt.Sprintf("%s %s %s", car, hotel, flight)
	return output, nil
}
