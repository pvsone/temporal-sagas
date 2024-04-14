package activities

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/activity"
)

func BookCar(ctx context.Context, carId string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("BookCar activity started", "carId", carId)

	time.Sleep(1 * time.Second)

	result := fmt.Sprintf("Booked car: %s", carId)
	return result, nil
}

func BookHotel(ctx context.Context, hotelId string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("BookHotel activity started", "hotelId", hotelId)

	time.Sleep(1 * time.Second)

	result := fmt.Sprintf("Booked hotel: %s", hotelId)
	return result, nil
}

func BookFlight(ctx context.Context, flightId string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("BookFlight activity started", "flightId", flightId)

	time.Sleep(1 * time.Second)

	result := fmt.Sprintf("Booked flight: %s", flightId)
	return result, nil
}
