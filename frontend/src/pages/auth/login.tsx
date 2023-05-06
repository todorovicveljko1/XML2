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
    TextField,
    Typography,
} from "@mui/material";
import { useState } from "react";
import { useMutation } from "react-query";

function Login() {
    const { setToken } = useToken();
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState(false);
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
                setError(true);
            },
        }
    );

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
                            <TextField
                                error={error}
                                label="Username:"
                                required
                                fullWidth
                                value={username}
                                onChange={(e) => {
                                    error && setError(false);
                                    setUsername(e.target.value);
                                }}
                            />
                            <TextField
                                error={error}
                                label="Password:"
                                type="password"
                                required
                                fullWidth
                                value={password}
                                onChange={(e) => {
                                    error && setError(false);
                                    setPassword(e.target.value);
                                }}
                            />
                            {error && (
                                <Alert severity="error" variant="outlined">
                                    Wrong credentials
                                </Alert>
                            )}
                            <LoadingButton
                                variant="contained"
                                loading={mutation.isLoading}
                                onClick={() =>
                                    mutation.mutate({ username, password })
                                }
                            >
                                Login
                            </LoadingButton>
                        </Stack>
                    </Paper>
                </Container>
            </GuestWraper>
        </MainLayout>
    );
}

export default Login;
