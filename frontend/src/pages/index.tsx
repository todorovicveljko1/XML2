import axios from "@/axios";
import { AccommodationCard } from "@/components/Cards/AccommodationCard";
import MainLayout from "@/components/Layout/MainLayout";
import BackdropLoader from "@/components/Loaders/backdropLoader";
import { AuthShow, useAuth } from "@/providers/authProvider";
import { Alert, Container, Fab, Grid } from "@mui/material";
import { useRouter } from "next/router";
import { useQuery } from "react-query";

export default function Home() {
    const router = useRouter();
    const { data, isLoading, error } = useQuery("accommodation", () => {
        return axios.get("/accommodation");
    });
    return (
        <MainLayout>
            {isLoading ? (
                <BackdropLoader text="Loading" />
            ) : error ? (
                <Alert severity="error">
                    Error while fetching accommodations
                </Alert>
            ) : (
                <Grid p={3} spacing={3} container>
                    {data &&
                        data.data.accommodations.map((accommodation: any) => (
                            <Grid key={accommodation.id} item xs={12} sm={6} md={4} lg={3}>
                                <AccommodationCard
                                    accommodation={accommodation}
                                />
                            </Grid>
                        ))}
                </Grid>
            )}
            <AuthShow roles={["H"]}>
                <Fab
                    color="primary"
                    sx={{
                        position: "fixed",
                        bottom: 16,
                        right: 16,
                        fontSize: "3rem",
                    }}
                    onClick={() => router.push("/accommodation/create")}
                >
                    +
                </Fab>
            </AuthShow>
        </MainLayout>
    );
}
