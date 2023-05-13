export interface Accommodation {
    id: string;
    name: string;
    location: string;
    amenity: string[];
    photo_url: string[];
    max_guests: number;
    min_guests: number;
    default_price: number;
    user_id: string;
    is_price_per_night: boolean;
    is_manual: boolean;
}

export interface CreateAccomodationRequest {
    name: string;
    location: string;
    amenity: string[];
    photo_url: string[];
    max_guests: number;
    min_guests: number;
    default_price: number;
    user_id: string;
    is_price_per_night: boolean;
    is_manual: boolean;
}

export interface AvailableInterval {
    start_date: string;
    end_date: string;
    is_available: boolean;
}

export interface PriceInterval {
    start_date: string;
    end_date: string;
    price: number;
}

export interface GetAccommodationResponse {
    accommodation: Accommodation;
    available_intervals: AvailableInterval[];
    price_intervals: PriceInterval[];
}

export interface AddAvailabilityRequest {
    id: string;
    availability: AvailableInterval;
}

export interface AddPriceRequest {
    id: string;
    price: PriceInterval;
}