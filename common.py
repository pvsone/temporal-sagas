from dataclasses import dataclass

@dataclass
class BookVacationInput:
    book_user_id: str
    book_car_id: str
    book_hotel_id: str
    book_flight_id: str
    attempts: int
