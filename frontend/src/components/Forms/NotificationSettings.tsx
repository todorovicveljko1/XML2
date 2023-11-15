import axios from "@/axios";
import { LoadingButton } from "@mui/lab";
import { Alert, Stack, Typography } from "@mui/material";
import { enqueueSnackbar } from "notistack";
import { useEffect } from "react";
import { useForm } from "react-hook-form";
import { SwitchElement } from "react-hook-form-mui";
import { useMutation, useQuery } from "react-query";
import BackdropLoader from "../Loaders/backdropLoader";
import { useAuth } from "@/providers/authProvider";
const defaultValues = {
    reservation_created: true,
    reservation_canceled: true,
    reservation_accepted: true,
    reservation_rejected: true,
    rating_modified: true,
};
const PREETY_NAMES: Record<string, string> = {
    reservation_created: "Reservation created",
    reservation_canceled: "Reservation canceled",
    reservation_accepted: "Reservation approved",
    reservation_rejected: "Reservation rejected",
    rating_modified: "Rating modified",
};

const ROLE_FOR_NOTIFICATIONS:  Record<string, string[]> = {
    reservation_created: ["H"],
    reservation_canceled: ["H"],
    reservation_accepted: ["G"],
    reservation_rejected: ["G"],
    rating_modified: ["H"],
}


export function NotificationSettings() {
    const {hasAccess} = useAuth();
    const { control, handleSubmit, formState, reset } = useForm<
        Record<string, boolean>
    >({
        defaultValues,
    });

    const { data, isLoading, error } = useQuery(
        ["notification-settings"],
        () => {
            return axios.get("/notification/settings");
        }
    );

    const mutation = useMutation(
        (data: any) => {
            return axios.put("/notification/settings", data);
        },
        {
            onSuccess(data, variables, context) {
                enqueueSnackbar({
                    message: "Successfully updated notification settings",
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
        // transform data to match backend
        const transformedData = Object.keys(data).map((key: any) => {
            return {
                type: key,
                enabled: data[key],
            };
        });
        mutation.mutate(transformedData);
    });

    useEffect(() => {
        if (data?.data) {
            const transformedData = data.data.reduce((acc: any, curr: any) => {
                acc[curr.type] = curr.enabled ?? false;
                return acc;
            }, {});
            reset({ ...defaultValues, ...transformedData });
        }
    }, [data, reset]);

    return (
        <Stack spacing={3} p={3}>
            <Typography variant="h6" gutterBottom>
                Notification Settings
            </Typography>
            {isLoading ? (
                <BackdropLoader />
            ) : error ? (
                <Alert severity="error">
                    Error while getting notification settings
                </Alert>
            ) : (
                Object.keys(defaultValues).filter(
                    (key) => hasAccess(ROLE_FOR_NOTIFICATIONS[key])
                ).map((key) => (
                    <SwitchElement
                        key={key}
                        control={control}
                        name={key}
                        label={PREETY_NAMES[key]}
                        labelPlacement="start"
                        sx={{justifyContent: "space-between", alignItems: "center", borderBottom: "1px solid #ccc"}}
                    />
                ))
            )}
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
