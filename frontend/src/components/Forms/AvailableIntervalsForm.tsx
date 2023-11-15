import axios from "@/axios";
import { LoadingButton } from "@mui/lab";
import { Button, Stack, Typography } from "@mui/material";
import dayjs from "dayjs";
import router from "next/router";
import { enqueueSnackbar } from "notistack";
import {
    CheckboxElement,
    DatePickerElement,
    useForm,
} from "react-hook-form-mui";
import { useMutation } from "react-query";

interface AvailableIntervalFormProps {
    accomodationId: string;
    onSuccess?: () => void;
}

export function AvailableIntervalForm({
    accomodationId,
    onSuccess,
}: AvailableIntervalFormProps) {
    const { control, handleSubmit } = useForm({
        defaultValues: {
            start_date: dayjs(),
            end_date: dayjs().add(1, "day"),
            is_available: false,
        },
    });

    const mutation = useMutation(
        (data: any) => {
            return axios.put(
                `/accommodation/${accomodationId}/availability`,
                data
            );
        },
        {
            onSuccess(data, variables, context) {
                enqueueSnackbar({
                    message: "Succesfully updated availability",
                    variant: "success",
                });
                if (onSuccess) onSuccess();
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
        <Stack spacing={1}>
            <Typography variant="h6" gutterBottom>
                Update availability
            </Typography>
            <Stack direction="row" spacing={3}>
                <DatePickerElement
                    control={control}
                    name="start_date"
                    label="Interval start:"
                    sx={{flexGrow:1}}
                />
                <DatePickerElement
                    control={control}
                    name="end_date"
                    label="Interval end:"
                    sx={{flexGrow:1}}
                />
            </Stack>
            <CheckboxElement
                control={control}
                name="is_available"
                label="Is available in this interval?"
            />
            <Stack direction={"row"} spacing={2} justifyContent={"end"}>
                <LoadingButton
                    variant="contained"
                    color="primary"
                    loading={mutation.isLoading}
                    onClick={submit}
                >
                    Update
                </LoadingButton>
            </Stack>
        </Stack>
    );
}
