from datetime import timedelta

from temporalio import workflow
from temporalio.common import RetryPolicy

with workflow.unsafe.imports_passed_through():
    from activities import BookVacationInput, book_car, book_flight, book_hotel


@workflow.defn
class BookWorkflow:
    @workflow.run
    async def run(self, input: BookVacationInput):
        compensations = []

        activity_args = {
            "start_to_close_timeout": timedelta(seconds=10),
            "retry_policy": RetryPolicy(
                non_retryable_error_types=["Exception"],
            ),
        }

        try:
            compensations.append("undo_book_car")
            car = await workflow.execute_activity(
                book_car,
                input,
                **activity_args,
            )
            compensations.append("undo_book_hotel")
            hotel = await workflow.execute_activity(
                book_hotel,
                input,
                **activity_args,
            )

            compensations.append("undo_book_flight")
            flight = await workflow.execute_activity(
                book_flight,
                input,
                start_to_close_timeout=timedelta(seconds=10),
                retry_policy=RetryPolicy(
                    initial_interval=timedelta(seconds=1),
                    maximum_interval=timedelta(seconds=1),
                    maximum_attempts=input.attempts,
                    non_retryable_error_types=["Exception"],
                ),
            )
            return car + " " + hotel + " " + flight
        except Exception:
            for compensation in reversed(compensations):
                await workflow.execute_activity(
                    compensation,
                    input,
                    start_to_close_timeout=timedelta(seconds=10),
                )

            return "Booking cancelled"
