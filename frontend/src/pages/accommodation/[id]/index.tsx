import axios from "@/axios";
import MainLayout from "@/components/Layout/MainLayout";
import BackdropLoader from "@/components/Loaders/backdropLoader";
import { AuthShow, useAuth } from "@/providers/authProvider";
import {
    Alert,
    Box,
    Button,
    Chip,
    Container,
    Grid,
    Paper,
    Stack,
    Typography,
} from "@mui/material";
import { useRouter } from "next/router";
import { useQuery } from "react-query";

export default function AccommodationPage() {
    const router = useRouter();
    const { id } = router.query;
    const { user } = useAuth();
    const { data, isLoading, error } = useQuery(
        ["accommodation", id],
        () => {
            return axios.get(`/accommodation/${id}`);
        },
        { enabled: !!id }
    );
    const accommodation = data?.data?.accommodation;
    return (
        <MainLayout>
            {isLoading ? (
                <BackdropLoader text="Loading" />
            ) : error ? (
                <Alert severity="error">
                    Error while fetching accommodation
                </Alert>
            ) : (
                accommodation && (
                    <Container maxWidth="md" sx={{ marginTop: 3 }}>
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
                                    <Typography
                                        variant="h6"
                                        component="h2"
                                        color="green"
                                    >
                                        {accommodation.default_price}${" "}
                                        {accommodation.is_price_per_night
                                            ? "/ night"
                                            : "/ guest"}
                                    </Typography>
                                </Stack>
                                <Stack flexWrap={"wrap"} direction={"row"}>
                                    <Chip
                                        color="info"
                                        label={`GUESTS: ${accommodation.min_guests} - ${accommodation.max_guests}`}
                                        sx={{ mt: 1, mr: 1 }}
                                    />
                                    {accommodation.amenity.map(
                                        (amenity: string) => (
                                            <Chip
                                                key={amenity}
                                                label={amenity}
                                                sx={{ mt: 1, mr: 1 }}
                                            />
                                        )
                                    )}
                                </Stack>
                                <Grid container>
                                    {accommodation.photo_url.map(
                                        (photo: string) => (
                                            <Grid item lg={4} key={photo}>
                                                <Box
                                                    component={"img"}
                                                    width={"100%"}
                                                    src={photo}
                                                ></Box>
                                            </Grid>
                                        )
                                    )}
                                </Grid>
                                <Stack
                                    direction={"row"}
                                    justifyContent={"space-between"}
                                    alignItems={"center"}
                                >
                                    <Button
                                        onClick={() => router.back()}
                                    >
                                        Back
                                    </Button>
                                    <Stack
                                        spacing={2}
                                        direction={"row-reverse"}
                                    >
                                        <AuthShow roles={["G"]}>
                                            <Button variant="contained">
                                                Reserve
                                            </Button>
                                        </AuthShow>
                                        <AuthShow roles={["H"]}>
                                            {user &&
                                                user.id ==
                                                    accommodation.user_id && (
                                                    <>
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
                    </Container>
                )
            )}
        </MainLayout>
    );
}
