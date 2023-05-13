import axios from "@/axios";
import { AccommodationCard } from "@/components/Cards/AccommodationCard";
import { AccommodationFilterForm } from "@/components/Forms/AccommodationFilterForm";
import MainLayout from "@/components/Layout/MainLayout";
import BackdropLoader from "@/components/Loaders/backdropLoader";
import { buildGetQuery } from "@/helpers/buildGetQuery";
import { AuthShow, useAuth } from "@/providers/authProvider";
import { Alert, Container, Fab, Grid } from "@mui/material";
import dayjs from "dayjs";
import { useRouter } from "next/router";
import { useState } from "react";
import { useQuery } from "react-query";

export default function Home() {
    const router = useRouter();
    const [query, setQuery] = useState("");
    const { data, isLoading, error } = useQuery(
        ["accommodation", query],
        () => {
            return axios.get("/accommodation" + query);
        }
    );
    return (
        <MainLayout>
            <Container sx={{ pt: 3, px: 3}} maxWidth={false}>
                <AccommodationFilterForm
                    onFilter={(data) =>
                        setQuery(
                            "?" +
                                buildGetQuery({
                                    ...data,
                                    start_date: dayjs(
                                        data.start_date
                                    ).toISOString(),
                                    end_date: dayjs(
                                        data.end_date
                                    ).toISOString(),
                                })
                        )
                    }
                />
            </Container>
            {isLoading ? (
                <BackdropLoader text="Loading" />
            ) : error ? (
                <Alert severity="error">
                    Error while fetching accommodations
                </Alert>
            ) : (
                <Grid p={3} spacing={3} container>
                    {data && (
                        <>
                            {data.data.accommodations.map(
                                (accommodation: any) => (
                                    <Grid
                                        key={accommodation.id}
                                        item
                                        xs={12}
                                        sm={6}
                                        md={4}
                                        lg={3}
                                    >
                                        <AccommodationCard
                                            accommodation={accommodation}
                                        />
                                    </Grid>
                                )
                            )}
                        </>
                    )}
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
