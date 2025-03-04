import axios from "@/axios";
import { AMENETIES } from "@/consts";
import { useAuth } from "@/providers/authProvider";
import { Accommodation } from "@/types/accomodation";
import { LoadingButton } from "@mui/lab";
import { Button, Stack, Typography } from "@mui/material";
import { useRouter } from "next/router";
import { enqueueSnackbar } from "notistack";
import { useForm } from "react-hook-form";
import { CheckboxElement, MultiSelectElement, TextFieldElement } from "react-hook-form-mui";
import { useMutation } from "react-query";

interface EditAccomodationFormProps {
    accommodation: Accommodation;
}

export function EditAccomodationForm({accommodation}: EditAccomodationFormProps)  {
    const router = useRouter();
    const { control, handleSubmit, formState, reset } = useForm({
        defaultValues: {
            name: accommodation.name,
            location: accommodation.location,
            amenity: accommodation.amenity,
            photo_url: accommodation.photo_url.join(","),
            max_guests: accommodation.max_guests.toString(),
            min_guests: accommodation.min_guests.toString(),
            default_price: accommodation.default_price.toString(),
            is_price_per_night: accommodation.is_price_per_night,
            is_manual: accommodation.is_manual,
        },
    });

    const mutation = useMutation(
        (data: any) => {
            return axios.put(`/accommodation/${accommodation.id}`, data);
        },
        {
            onSuccess(data, variables, context) {
                enqueueSnackbar({
                    message: "Succesfully updated accommodation",
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
        const jsonData = {
            ...data,
            amenity: data.amenity,
            photo_url: data.photo_url.split(","),
            max_guests: parseInt(data.max_guests),
            min_guests: parseInt(data.min_guests),
            default_price: parseFloat(data.default_price),
        };
        mutation.mutate(jsonData);
    });

    return (
        <Stack spacing={3} p={3}>
            <Typography variant="h6" gutterBottom>
                Edit accomodation
            </Typography>
            <TextFieldElement
                control={control}
                name="name"
                label="Name"
                required
            />

            <TextFieldElement
                control={control}
                name="location"
                label="Location"
                required
            />
            <MultiSelectElement
                control={control}
                name="amenity"
                label="Amenity"
                required
                options={AMENETIES}
            />
            <TextFieldElement
                control={control}
                name="photo_url"
                label="Photo url"
                required
            />
            <TextFieldElement
                control={control}
                name="min_guests"
                label="Min guests"
                required
                type="number"
            />
            <TextFieldElement
                control={control}
                name="max_guests"
                label="Max guests"
                required
                type="number"
            />
            <TextFieldElement
                control={control}
                name="default_price"
                label="Default price"
                required
                type="number"
            />
            <CheckboxElement
                control={control}
                name="is_price_per_night"
                label="Price is calculated per night?"
            />
            <CheckboxElement
                control={control}
                name="is_manual"
                label="Reservations are manually managed?"
            />
            <Stack direction={"row"} spacing={2}  justifyContent={"end"}>
                <Button onClick={() => router.back()}>Cancle</Button>
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
