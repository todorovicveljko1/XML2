import { number } from "yup";

export type Reservation = {
    id: string;
    accommodation_id: string;
    start_date: string;
    end_date: string;
    user_id: string;
    price: number;
    number_of_guests: number;
    status: string;
}


export type ReservationList = {
    reservations: Reservation[];
}