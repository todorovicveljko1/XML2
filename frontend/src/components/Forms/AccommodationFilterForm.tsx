import { AMENETIES } from "@/consts";
import { AuthShow } from "@/providers/authProvider";
import { Button, Stack } from "@mui/material";
import dayjs from "dayjs";
import { useForm } from "react-hook-form";
import {
    CheckboxElement,
    DatePickerElement,
    MultiSelectElement,
    TextFieldElement,
} from "react-hook-form-mui";
import * as Yup from "yup";
import { yupResolver } from "@hookform/resolvers/yup";
import { LoadingButton } from "@mui/lab";

const YUP_SCHEMA = Yup.object().shape({
    location: Yup.string(),
    num_guests: Yup.number()
        .min(1)
        .required()
        .transform((_, val) => (val === Number(val) ? val : null)),
    start_date: Yup.date().min(dayjs().startOf("day").toDate()).required(),
    // validate end date is after start date
    end_date: Yup.date().min(dayjs().startOf("day").toDate()).required(),
    amenity: Yup.array().of(Yup.string()),
    show_my: Yup.boolean(),
    show_super_hosts: Yup.boolean(),
    price_min: Yup.number()
        .min(0)
        .transform((_, val) => (val === Number(val) ? val : null)),
    price_max: Yup.number()
        .min(0)
        .transform((_, val) => (val === Number(val) ? val : null)),
});

interface AccommodationFilterFormProps {
    onFilter: (data: any) => void;
    isLoading: boolean;
}
export function AccommodationFilterForm({
    onFilter,
    isLoading,
}: AccommodationFilterFormProps) {
    const { control, handleSubmit } = useForm({
        defaultValues: {
            location: "",
            num_guests: 1,
            start_date: dayjs(),
            end_date: dayjs().add(6, "day"),
            amenity: [],
            show_my: false,
            show_super_hosts: false,
            price_min: 0,
            price_max: 0,
        },
        resolver: yupResolver(YUP_SCHEMA),
    });

    const submit = handleSubmit((data) => {
        onFilter(data);
    });

    return (
        <Stack direction="row" alignItems={"center"} flexWrap={"wrap"}>
            <TextFieldElement
                sx={{ mr: 2, mb: 2 }}
                control={control}
                name="location"
                label="Location"
                size="small"
            />

            <DatePickerElement
                sx={{ mr: 2, mb: 2 }}
                control={control}
                name="start_date"
                label="Interval start:"
                inputProps={{ size: "small" }}
                required
                disablePast
            />
            <DatePickerElement
                sx={{ mr: 2, mb: 2 }}
                control={control}
                name="end_date"
                label="Interval end:"
                inputProps={{ size: "small" }}
                required
                disablePast
            />

            <TextFieldElement
                sx={{ mr: 2, mb: 2 }}
                type="number"
                control={control}
                name="num_guests"
                label="Number of guests"
                size="small"
            />
            <MultiSelectElement
                sx={{ mr: 2, mb: 2 }}
                control={control}
                name="amenity"
                label="Amenity"
                options={AMENETIES}
                size="small"
            />
            <AuthShow roles={["H"]}>
                <CheckboxElement
                    sx={{ mr: 2 }}
                    control={control}
                    name="show_my"
                    label="Show my accommodations"
                    size="small"
                />
            </AuthShow>
            <CheckboxElement
                sx={{ mr: 2 }}
                control={control}
                name="show_super_hosts"
                label="Show super hosts"
                size="small"
            />
            <TextFieldElement
                sx={{ mr: 2, mb: 2 }}
                type="number"
                control={control}
                name="price_min"
                label="Price min"
                size="small"
            />
            <TextFieldElement
                sx={{ mr: 2, mb: 2 }}
                type="number"
                control={control}
                name="price_max"
                label="Price max"
                size="small"
            />
            <LoadingButton
                sx={{ mr: 2, mb: 2 }}
                loading={isLoading}
                variant="contained"
                onClick={() => submit()}
            >
                Filter
            </LoadingButton>
        </Stack>
    );
}
