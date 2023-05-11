import axios from "@/axios";
import { LoadingButton } from "@mui/lab";
import { Stack, Typography } from "@mui/material";
import { enqueueSnackbar } from "notistack";
import { useForm } from "react-hook-form";
import { TextFieldElement } from "react-hook-form-mui";
import { useMutation } from "react-query";

export function ChangePasswordForm() {
    const { control, handleSubmit, formState, reset } = useForm({
        defaultValues: {
            old_password: "",
            new_password: "",
        },
    });

    const mutation = useMutation(
        (data: any) => {
            return axios.put("/auth/change-password", data);
        },
        {
            onSuccess(data, variables, context) {
                enqueueSnackbar({
                    message: "Succesfully updated password",
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
                Change Password
            </Typography>
            <TextFieldElement
                control={control}
                name="old_password"
                label="Old Password"
                type="password"
                required
            />
            <TextFieldElement
                control={control}
                name="new_password"
                label="New Password"
                type="password"
                required
            />
            <LoadingButton
                variant="contained"
                loading={mutation.isLoading}
                onClick={submit}
            >
                Change password
            </LoadingButton>
        </Stack>
    );
}
