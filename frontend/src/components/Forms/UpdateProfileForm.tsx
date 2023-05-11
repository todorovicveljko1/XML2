import axios from "@/axios";
import { useAuth } from "@/providers/authProvider";
import { LoadingButton } from "@mui/lab";
import { Stack, Typography } from "@mui/material";
import { enqueueSnackbar } from "notistack";
import { useEffect } from "react";
import { useForm } from "react-hook-form";
import { TextFieldElement } from "react-hook-form-mui";
import { useMutation } from "react-query";

export function UpdateProfileForm() {
    const { user } = useAuth();
    const { control, handleSubmit, formState, reset } = useForm({
        defaultValues: {
            username: user?.username ?? "",
            first_name: user?.first_name ?? "",
            last_name: user?.last_name ?? "",
            email: user?.email ?? "",
            place_of_living: user?.place_of_living ?? "",
        },
    });

    useEffect(() => {
        reset(user);
    }, [user, reset]);

    const mutation = useMutation(
        (data: any) => {
            return axios.put("/auth", data);
        },
        {
            onSuccess(data, variables, context) {
                enqueueSnackbar({
                    message: "Succesfully updated profile",
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
    const submit = handleSubmit((data) => {
        mutation.mutate(data);
    });
    
    return (
        <Stack spacing={3} p={3}>
            <Typography variant="h6" gutterBottom>
                Personal informations
            </Typography>
            <TextFieldElement
                control={control}
                name="username"
                label="Username"
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

            <LoadingButton
                variant="contained"
                loading={mutation.isLoading}
                onClick={submit}
            >
                Save
            </LoadingButton>
        </Stack>
    );
}
