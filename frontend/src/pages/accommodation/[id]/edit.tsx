import axios from "@/axios";
import { EditAccomodationForm } from "@/components/Forms/EditAccomodationForm";
import MainLayout from "@/components/Layout/MainLayout";
import BackdropLoader from "@/components/Loaders/backdropLoader";
import { AuthWraper, useAuth } from "@/providers/authProvider";
import { Alert, Container, Paper, Stack, Typography } from "@mui/material";
import { useRouter } from "next/router";
import { useQuery } from "react-query";

export default function EditAccomodation() {
    const router = useRouter();
    const { user } = useAuth();
    const { id } = router.query;
    const { data, isLoading, error } = useQuery(
        ["accommodation", id],
        () => {
            return axios.get(`/accommodation/${id}`);
        },
        { enabled: !!id }
    );
    const accommodation = data?.data?.accommodation;

    if (user && accommodation && user.id != accommodation.user_id)
        router.push(`/accommodation/${accommodation.id}`);

    return (
        <MainLayout>
            <AuthWraper roles={["H"]}>
                <Container maxWidth="md" sx={{ py: 3 }}>
                    <Paper>
                        {isLoading ? (
                            <BackdropLoader text="Loading" />
                        ) : error ? (
                            <Alert severity="error">
                                Error while fetching accommodation
                            </Alert>
                        ) : (
                            accommodation && (
                                <EditAccomodationForm
                                    accommodation={accommodation}
                                />
                            )
                        )}
                    </Paper>
                </Container>
            </AuthWraper>
        </MainLayout>
    );
}
