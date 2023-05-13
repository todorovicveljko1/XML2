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
}

export function AccommodationCard({ accommodation }: AccommodationCardProps) {
    const router = useRouter();
    return (
        <Card sx={{ minWidth: 275 }}>
            <CardActionArea onClick={()=>router.push(`/accommodation/${accommodation.id}`)}>
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
                            <Typography
                                gutterBottom
                                variant="h5"
                                component="div"
                            >
                                {accommodation.name}
                            </Typography>
                            <Typography variant="h6" color="green">
                                {accommodation.default_price}${" "}
                                {accommodation.is_price_per_night
                                    ? "/ night"
                                    : "/ guest"}
                            </Typography>
                        </Stack>
                        <Typography variant="body2" color="text.secondary">
                            {accommodation.location}
                        </Typography>
                        <Stack flexWrap={"wrap"} direction={"row"}>
                            <Chip
                                color="info"
                                label={`GUESTS: ${accommodation.min_guests} - ${accommodation.max_guests}`}
                                sx={{mt:1, mr:1}}
                            />
                            {accommodation.amenity.map((amenity) => (
                                <Chip key={amenity} label={amenity} sx={{mt:1, mr:1}}/>
                            ))}
                        </Stack>
                    </Stack>
                </CardContent>
            </CardActionArea>
        </Card>
    );
}
