import axios from "@/axios";
import { ChangePasswordForm } from "@/components/Forms/ChangePasswordForm";
import { UpdateProfileForm } from "@/components/Forms/UpdateProfileForm";
import MainLayout from "@/components/Layout/MainLayout";
import { AuthWraper, useAuth } from "@/providers/authProvider";
import { LoadingButton } from "@mui/lab";
import { Container, Divider, Paper, Stack, Typography } from "@mui/material";
import { enqueueSnackbar } from "notistack";
import { useMutation } from "react-query";

function Profile() {
    const {logout} = useAuth()
    const mutation = useMutation(
        () => {
            return axios.delete("/auth");
        },
        {
            onSuccess(data, variables, context) {
                logout()
                enqueueSnackbar({
                    message: "Succesfully deleted profile",
                    variant: "success",
                });
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

    return (
        <MainLayout>
            <AuthWraper>
                <Container maxWidth="md" sx={{ py: 3 }}>
                    <Paper>
                        <Stack sx={{ p: 3 }}>
                            <Typography variant="h4" gutterBottom sx={{ m: 0 }}>
                                Profile
                            </Typography>
                        </Stack>
                        <Divider />
                        <UpdateProfileForm />
                        <Divider />
                        <ChangePasswordForm />
                        <Divider />
                        <Stack p={3}>
                            <Typography variant="h6" gutterBottom>
                                Delete profile
                            </Typography>
                            <Typography variant="body2" gutterBottom>
                                This action is irreversible. All your data will
                                be deleted.
                                </Typography>
                                <LoadingButton
                                    variant="contained"
                                    color="error"
                                    loading={mutation.isLoading}
                                    onClick={()=>mutation.mutate()}
                                >
                                    Delete profile
                                </LoadingButton>
                            </Stack>
                    </Paper>
                </Container>
            </AuthWraper>
        </MainLayout>
    );
}

export default Profile;
