import { AuthShow, useAuth } from "@/providers/authProvider";
import { Accommodation } from "@/types/accomodation";
import {
    Card,
    CardContent,
    Typography,
    CardActions,
    Button,
    CardActionArea,
    CardMedia,
    Stack,
    Chip,
} from "@mui/material";
import { useRouter } from "next/router";

interface AccommodationCardProps {
    accommodation: Accommodation;
    onCreateReservation: (accommodationId: string, price: number) => void;
}

export function AccommodationCard({
    accommodation,
    onCreateReservation,
}: AccommodationCardProps) {
    const { user } = useAuth();
    const router = useRouter();
    return (
        <Card sx={{ minWidth: 275 }}>
            <CardActionArea
                onClick={() =>
                    router.push(`/accommodation/${accommodation.id}`)
                }
            >
                {accommodation.photo_url.length != 0 && (
                    <CardMedia
                        component={"img"}
                        width={"100%"}
                        image={accommodation.photo_url[0]}
                        alt={accommodation.name}
                    ></CardMedia>
                )}
                <CardContent>
                    <Stack spacing={1}>
                        <Stack
                            direction={"row"}
                            justifyContent={"space-between"}
                            alignItems={"center"}
                        >
                            <Typography variant="h5" component="div">
                                {accommodation.name}
                            </Typography>
                            <Typography variant="h6" color="green">
                                {accommodation.price
                                    ? `${accommodation.price}$ total`
                                    : `${accommodation.default_price}$${
                                          accommodation.is_price_per_night
                                              ? "/ night"
                                              : "/ guest"
                                      }`}
                            </Typography>
                        </Stack>
                        <Typography variant="body2" color="text.secondary">
                            {accommodation.location}
                        </Typography>
                        <Stack flexWrap={"wrap"} direction={"row"}>
                            <Chip
                                size="small"
                                color="info"
                                label={`GUESTS: ${accommodation.min_guests} - ${accommodation.max_guests}`}
                                sx={{ mt: 1, mr: 1 }}
                            />
                            {accommodation.amenity.map((amenity) => (
                                <Chip
                                    size="small"
                                    key={amenity}
                                    label={amenity}
                                    sx={{ mt: 1, mr: 1 }}
                                />
                            ))}
                        </Stack>
                    </Stack>
                </CardContent>
            </CardActionArea>
            <AuthShow roles={["G"]}>
                <CardActions sx={{ px: 2, pb: 2, float: "right" }}>
                    <Button
                        variant="contained"
                        size="small"
                        onClick={() =>
                            onCreateReservation(
                                accommodation.id,
                                accommodation?.price ??
                                    accommodation.default_price
                            )
                        }
                    >
                        Create Reserve
                    </Button>
                </CardActions>
            </AuthShow>
            <AuthShow roles={["H"]}>
                {user && user.id == accommodation.user_id && (
                    <CardActions sx={{ px: 2, pb: 2, float: "right" }}>
                        <Button
                            variant="contained"
                            size="small"
                            onClick={() =>
                                router.push(
                                    `/accommodation/${accommodation.id}/reservation`
                                )
                            }
                        >
                            Reservations
                        </Button>
                    </CardActions>
                )}
            </AuthShow>
        </Card>
    );
}
