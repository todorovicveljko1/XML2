import axios from "@/axios";
import MainLayout from "@/components/Layout/MainLayout";
import { GuestWraper } from "@/providers/authProvider";
import { RegisterDto } from "@/types/user";
import { LoadingButton } from "@mui/lab";
import { Alert, Container, Paper, Stack, Typography } from "@mui/material";
import { useRouter } from "next/router";
import { useSnackbar } from "notistack";
import { TextFieldElement, SelectElement, useForm } from "react-hook-form-mui";
import { useMutation } from "react-query";

// role options for G as Guest and H as Host
const ROLE_OPTIONS = [
    { id: "G", label: "Guest" },
    { id: "H", label: "Host" },
];

function Register() {
    const router = useRouter();
    const { enqueueSnackbar, closeSnackbar } = useSnackbar();
    const { control, handleSubmit, formState } = useForm({
        defaultValues: {
            username: "",
            password: "",
            first_name: "",
            last_name: "",
            email: "",
            place_of_living: "",
            role: "",
        },
    });

    const mutation = useMutation(
        (data: RegisterDto) => {
            return axios.post("/auth/register", data);
        },
        {
            onSuccess(data, variables, context) {
                enqueueSnackbar({
                    message: "Succesfully created new account",
                    variant: "success",
                });
                router.push("/auth/login");
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
            <GuestWraper>
                <Container maxWidth="xs" sx={{ marginTop: "5vh" }}>
                    <Paper sx={{ p: 3 }}>
                        <Stack spacing={3}>
                            <Typography
                                variant="h5"
                                textAlign={"center"}
                                gutterBottom
                                mb={2}
                            >
                                Register
                            </Typography>
                            <TextFieldElement
                                control={control}
                                name="username"
                                label="Username"
                                required
                            />
                            <TextFieldElement
                                control={control}
                                name="password"
                                label="Password"
                                type="password"
                                required
                            />
                            <TextFieldElement
                                control={control}
                                name="first_name"
                                label="First name"
                                required
                            />
                            <TextFieldElement
                                control={control}
                                name="last_name"
                                label="Last name"
                                required
                            />
                            <TextFieldElement
                                control={control}
                                name="email"
                                label="Email"
                                required
                            />
                            <TextFieldElement
                                control={control}
                                name="place_of_living"
                                label="Place of living"
                                required
                            />
                            <SelectElement
                                control={control}
                                name="role"
                                label="Role"
                                options={ROLE_OPTIONS}
                                required
                            />
                            <LoadingButton
                                variant="contained"
                                loading={mutation.isLoading}
                                onClick={submit}
                            >
                                Register
                            </LoadingButton>
                        </Stack>
                    </Paper>
                </Container>
            </GuestWraper>
        </MainLayout>
    );
}

export default Register;
