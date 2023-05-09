import axios from "@/axios";
import MainLayout from "@components/Layout/MainLayout";
import { GuestWraper } from "@providers/authProvider";
import { useToken } from "@providers/tokenProvider";
import { LoginDto } from "@/types/user";
import { LoadingButton } from "@mui/lab";
import {
    Alert,
    Container,
    Paper,
    Stack,
    Typography,
} from "@mui/material";
import { useState } from "react";
import { useMutation } from "react-query";
import { useForm } from "react-hook-form";
import { TextFieldElement } from "react-hook-form-mui";

function Login() {
    const { setToken } = useToken();
    const {control, handleSubmit, formState} = useForm({
        defaultValues: {
            username: "",
            password: "",
        },
        
    });
    const [error, setError] = useState<string | null>(null);
    const mutation = useMutation(
        (data: LoginDto) => {
            return axios.post("/auth/login", data);
        },
        {
            onSuccess(data, variables, context) {
                if (data.status === 200) {
                    setToken(data.data.token);
                }
            },
            onError(error, variables, context) {
                setError("Wrong credentials");
            },
        }
    );

    const submit = handleSubmit((data) => {
        mutation.mutate(data);
    });

    return (
        <MainLayout>
            <GuestWraper>
                <Container maxWidth="xs" sx={{ marginTop: "15vh" }}>
                    <Paper sx={{ p: 3 }}>
                            <Stack spacing={3}>
                                <Typography
                                    variant="h5"
                                    textAlign={"center"}
                                    gutterBottom
                                    mb={2}
                                >
                                    Login
                                </Typography>
                                {error && (
                                    <Alert severity="error">{error}</Alert>
                                )}

                                <TextFieldElement
                                    name="username"
                                    label="Username"
                                    required
                                    control={control}
                                    
                                />
                                <TextFieldElement
                                    name="password"
                                    label="Password"
                                    type="password"
                                    required
                                    control={control}
                                />
                                <LoadingButton
                                    variant="contained"
                                    loading={mutation.isLoading}
                                    onClick={submit}
                                    disabled={!formState.isValid}
                                >
                                    {" "}
                                    Login{" "}
                                </LoadingButton>
                            </Stack>
                    </Paper>
                </Container>
            </GuestWraper>
        </MainLayout>
    );
}

export default Login;


