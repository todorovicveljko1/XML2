import axios from "@/axios";
import MainLayout from "@/components/Layout/MainLayout";
import { AuthWraper, GuestWraper } from "@/providers/authProvider";
import { useToken } from "@/providers/tokenProvider";
import { RegisterDto } from "@/types/user";
import { LoadingButton } from "@mui/lab";
import {
    Alert,
    Container,
    FormControl,
    InputLabel,
    MenuItem,
    Paper,
    Select,
    Stack,
    TextField,
    Typography,
} from "@mui/material";
import { useRouter } from "next/router";
import { useSnackbar } from "notistack";
import { use, useState } from "react";
import { useMutation } from "react-query";

// role options for G as Guest and H as Host
const ROLE_OPTIONS = [
    { value: "G", label: "Guest" },
    { value: "H", label: "Host" },
];

function Register() {
    const router = useRouter();
    const { enqueueSnackbar, closeSnackbar } = useSnackbar();
    const [username, setUsername] = useState("");
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [placeOfLiving, setPlaceOfLiving] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [role, setRole] = useState("G");
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
                            <TextField
                                error={mutation.isError}
                                label="Username:"
                                required
                                fullWidth
                                value={username}
                                onChange={(e) => {
                                    mutation.isError && mutation.reset();
                                    setUsername(e.target.value);
                                }}
                            />{" "}
                            <TextField
                                error={mutation.isError}
                                label="First name:"
                                required
                                fullWidth
                                value={firstName}
                                onChange={(e) => {
                                    mutation.isError && mutation.reset();
                                    setFirstName(e.target.value);
                                }}
                            />
                            <TextField
                                error={mutation.isError}
                                label="Last name:"
                                required
                                fullWidth
                                value={lastName}
                                onChange={(e) => {
                                    mutation.isError && mutation.reset();
                                    setLastName(e.target.value);
                                }}
                            />
                            <TextField
                                error={mutation.isError}
                                label="Email:"
                                required
                                fullWidth
                                value={email}
                                onChange={(e) => {
                                    mutation.isError && mutation.reset();
                                    setEmail(e.target.value);
                                }}
                            />
                            <TextField
                                error={mutation.isError}
                                label="Password:"
                                type="password"
                                required
                                fullWidth
                                value={password}
                                onChange={(e) => {
                                    mutation.isError && mutation.reset();
                                    setPassword(e.target.value);
                                }}
                            />
                            <TextField
                                error={mutation.isError}
                                label="Place of living:"
                                required
                                fullWidth
                                value={placeOfLiving}
                                onChange={(e) => {
                                    mutation.isError && mutation.reset();
                                    setPlaceOfLiving(e.target.value);
                                }}
                            />
                            <FormControl fullWidth>
                                <InputLabel id="demo-simple-select-label">
                                    Role
                                </InputLabel>
                                <Select label="Role"
                                    labelId="demo-simple-select-label"
                                    required
                                    fullWidth
                                    value={role}
                                    onChange={(e) => {
                                        mutation.isError && mutation.reset();
                                        setRole(e.target.value);
                                    }}
                                >
                                    {ROLE_OPTIONS.map((option) => (
                                        <MenuItem
                                            key={option.value}
                                            value={option.value}
                                        >
                                            {" "}
                                            {option.label}{" "}
                                        </MenuItem>
                                    ))}
                                </Select>
                            </FormControl>
                            <LoadingButton
                                variant="contained"
                                loading={mutation.isLoading}
                                onClick={() =>
                                    mutation.mutate({
                                        username,
                                        email,
                                        password,
                                        first_name: firstName,
                                        last_name: lastName,
                                        place_of_living: placeOfLiving,
                                        role,
                                    })
                                }
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
