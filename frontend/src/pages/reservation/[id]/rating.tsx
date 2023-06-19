import axios from "@/axios";
import MainLayout from "@/components/Layout/MainLayout";
import { AuthWraper } from "@/providers/authProvider";
import { Button, Container, Paper, Stack, Typography } from "@mui/material";
import { useRouter } from "next/router";
import { enqueueSnackbar } from "notistack";
import { useForm } from "react-hook-form";
import { useMutation } from "react-query";
import { TextFieldElement } from "react-hook-form-mui";
import { LoadingButton } from "@mui/lab";

export default function AddRating() {
    const router = useRouter();
    const { id } = router.query;
    const { control, handleSubmit } = useForm({
        defaultValues: {
            host_rating: 0,
            accommodation_rating: 0,
        },
    });

    const mutation = useMutation(
        (data: any) => {
            return axios.put(`/reservation/${id}/rating`, data);
        },
        {
            onSuccess(data, variables, context) {
                enqueueSnackbar({
                    message: "Succesfully added rating",
                    variant: "success",
                });

                router.push("/reservation");
            },
            onError(error: any, variables, context) {
                enqueueSnackbar({
                    message:
                        error?.response?.data?.error ??
                        error?.message ??
                        "Error",
                    variant: "error",
                });
            },
        }
    );

    const submit = handleSubmit((data) => {
        mutation.mutate(data);
    });

    return (
        <MainLayout>
            <AuthWraper roles={["G"]}>
                <Container maxWidth="xs" sx={{ py: 3 }}>
                    <Paper sx={{ p: 3 }}>
                        <Stack spacing={3}>
                            <Typography variant="h6" gutterBottom>
                                Modify rating
                            </Typography>
                            <TextFieldElement
                                name="host_rating"
                                control={control}
                                label="Host rating"
                                type="number"
                                required
                            />
                            <TextFieldElement
                                name="accommodation_rating"
                                control={control}
                                label="Accommodation rating"
                                type="number"
                                required
                            />
                            <Stack
                                direction={"row"}
                                spacing={2}
                                justifyContent={"space-between"}
                            >
                                <Button onClick={() => router.back()}>
                                    Back
                                </Button>
                                <LoadingButton
                                    variant="contained"
                                    color="primary"
                                    loading={mutation.isLoading}
                                    onClick={submit}
                                >
                                    Modify Rating
                                </LoadingButton>
                            </Stack>
                        </Stack>
                    </Paper>
                </Container>
            </AuthWraper>
        </MainLayout>
    );
}
