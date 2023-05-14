import { AuthShow, useAuth } from "@/providers/authProvider";
import { Accommodation } from "@/types/accomodation";
import {
    Paper,
    Stack,
    Typography,
    Chip,
    Grid,
    Box,
    Button,
} from "@mui/material";
import { useRouter } from "next/router";

interface AccommodationInfoLargeProps {
    accommodation: Accommodation;
}

export function AccommodationInfoLarge({
    accommodation,
}: AccommodationInfoLargeProps) {
    const router = useRouter();
    const { user } = useAuth();
    return (
        <Paper sx={{ p: 3 }}>
            <Stack spacing={2}>
                <Stack
                    direction={"row"}
                    justifyContent={"space-between"}
                    alignItems={"center"}
                >
                    <Typography variant="h4" component="h1">
                        {accommodation.name}
                    </Typography>
                    <Typography variant="h6" component="h2" color="green">
                        {accommodation.default_price}${" "}
                        {accommodation.is_price_per_night
                            ? "/ night"
                            : "/ guest"}
                    </Typography>
                </Stack>
                <Stack flexWrap={"wrap"} direction={"row"}>
                    <Chip
                        size="small"
                        color="info"
                        label={`GUESTS: ${accommodation.min_guests} - ${accommodation.max_guests}`}
                        sx={{ mt: 1, mr: 1 }}
                    />
                    {accommodation.amenity.map((amenity: string) => (
                        <Chip
                            size="small"
                            key={amenity}
                            label={amenity}
                            sx={{ mt: 1, mr: 1 }}
                        />
                    ))}
                </Stack>
                <Grid container>
                    {accommodation.photo_url.map((photo: string) => (
                        <Grid item lg={4} key={photo}>
                            <Box
                                component={"img"}
                                width={"100%"}
                                src={photo}
                            ></Box>
                        </Grid>
                    ))}
                </Grid>
                <Stack
                    direction={"row"}
                    justifyContent={"space-between"}
                    alignItems={"center"}
                >
                    <Button onClick={() => router.back()}>Back</Button>
                    <Stack spacing={2} direction={"row-reverse"}>
                        <AuthShow roles={["H"]}>
                            {user && user.id == accommodation.user_id && (
                                <>
                                    {" "}
                                    <Button
                                        variant="contained"
                                        onClick={() =>
                                            router.push(
                                                `/accommodation/${accommodation.id}/reservation`
                                            )
                                        }
                                    >
                                        Reservations
                                    </Button>
                                    <Button
                                        variant="contained"
                                        onClick={() =>
                                            router.push(
                                                `/accommodation/${accommodation.id}/edit`
                                            )
                                        }
                                    >
                                        Edit
                                    </Button>
                                </>
                            )}
                        </AuthShow>
                    </Stack>
                </Stack>
            </Stack>
        </Paper>
    );
}
