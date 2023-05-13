import axios from "@/axios";
import { AccommodationInfoLarge } from "@/components/Cards/AccommodationInfoLarge";
import { AvailableInterval } from "@/components/Cards/AvailableInterval";
import { PriceIntervalCard } from "@/components/Cards/PriceInterval";
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
    const { data, isLoading, error, refetch } = useQuery(
        ["accommodation", id],
        () => {
            return axios.get(`/accommodation/${id}`);
        },
        { enabled: !!id }
    );
    const accommodation = data?.data?.accommodation;
    const available_intervals = data?.data?.available_intervals ?? [];
    const price_intervals = data?.data?.price_intervals ?? [];

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
                    <>
                        <Container maxWidth="md" sx={{ marginTop: 3 }}>
                            <Stack spacing={3}>
                                <AccommodationInfoLarge
                                    accommodation={accommodation}
                                />
                                <AvailableInterval
                                    accommodationId={accommodation.id}
                                    intervals={available_intervals}
                                    onUpdate={()=>refetch()}
                                />
                                <PriceIntervalCard
                                    accommodationId={accommodation.id}
                                    intervals={price_intervals}
                                    onUpdate={()=>refetch()}
                                />
                            </Stack>
                        </Container>
                    </>
                )
            )}
        </MainLayout>
    );
}
