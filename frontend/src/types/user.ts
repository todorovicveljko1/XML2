export type UserType = {
    id: string;
    username: string;
    first_name: string;
    last_name: string;
    email: string;
    place_of_living: string;
    role: string;
};

export type LoginDto = {
    username: string;
    password: string;
};

export type RegisterDto = {
    username: string;
    first_name: string;
    last_name: string;
    email: string;
    place_of_living: string;
    role: string;
    password: string;
};