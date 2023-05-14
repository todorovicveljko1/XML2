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
import { enqueueSnackbar } from "notistack";
import { useEffect, useState } from "react";
import { useMutation, useQuery } from "react-query";

const DEFAULT_QUERY = {
    location: "",
    num_guests: 1,
    start_date: dayjs(),
    end_date: dayjs().add(6, "day"),
    amenity: [],
    show_my: false,
};

function BuildUrlQuert(data: any) {
    return (
        "?" +
        buildGetQuery({
            ...data,
            start_date:
                data.start_date &&
                dayjs(data.start_date as string).toISOString(),
            end_date:
                data.end_date && dayjs(data.end_date as string).toISOString(),
            show_my: data.show_my ? 1 : 0,
        })
    );
}

export default function Home() {
    const router = useRouter();
    const [query, setQuery] = useState<any>(DEFAULT_QUERY);
    const { data, isLoading, error } = useQuery(
        ["accommodation", query],
        () => {
            return axios.get("/accommodation" + BuildUrlQuert(query));
        },
        {
            keepPreviousData: true,
        }
    );

    const createReservationMutation = useMutation(
        (data: any) => {
            return axios.post(`/accommodation/${data.accommodationId}/reservation`, {
                start_date: dayjs(data.start_date).toISOString(),
                end_date: dayjs(data.end_date).toISOString(),
                num_guests: data.num_guests,
                price: data.price,
            });
        }
    )

    const handleCreateReservation = (accommodationId: string, price:number) => {
        createReservationMutation.mutate({
            accommodationId: accommodationId,
            start_date: query.start_date,
            end_date: query.end_date,
            num_guests: query.num_guests,
            price:price,
        }, {
            onSuccess: () => {
                enqueueSnackbar("Reservation created", {
                    variant: "success",
                });
                router.push("/reservation");
            }
        })
    }

    return (
        <MainLayout>
            <Container sx={{ pt: 3, px: 3 }} maxWidth={false}>
                <AccommodationFilterForm
                    isLoading={isLoading}
                    onFilter={(data) => setQuery(data)}
                />
            </Container>
            {isLoading ? (
                <BackdropLoader text="Loading" />
            ) : error ? (
                <Alert severity="error">
                    Error while fetching accommodations
                </Alert>
            ) : (
                <Grid p={3} spacing={3} container alignItems="stretch">
                    {data?.data?.accommodations ? (
                        <>
                            {data.data.accommodations.map(
                                (accommodation: any) => (
                                    <Grid
                                        key={accommodation.id}
                                        item
                                        xs={12}
                                        md={4}
                                        lg={3}
                                    >
                                        <AccommodationCard
                                            accommodation={accommodation}
                                            onCreateReservation={handleCreateReservation}
                                        />
                                    </Grid>
                                )
                            )}
                        </>
                    ) : (
                        <Grid item xs={12}>
                            <Alert severity="info">
                                No accommodations found
                            </Alert>
                        </Grid>
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
