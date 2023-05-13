import { AMENETIES } from "@/consts";
import { AuthShow } from "@/providers/authProvider";
import { Button, Stack } from "@mui/material";
import dayjs from "dayjs";
import { useForm } from "react-hook-form";
import { CheckboxElement, DatePickerElement, MultiSelectElement, TextFieldElement } from "react-hook-form-mui";

interface AccommodationFilterFormProps {
    onFilter: (data: any) => void;
}
export function AccommodationFilterForm({
    onFilter,
}: AccommodationFilterFormProps) {
    const { control, handleSubmit } = useForm({
        defaultValues: {
            location: "",
            num_guests: "",
            start_date: dayjs(),
            end_date:dayjs().add(1,'week'),
            amenity: [],
            show_my: false,
        },
    });

    const submit = handleSubmit((data) => {
        onFilter(data);
    });

    return (
        <Stack direction="row" spacing={3} alignItems={"center"}>
            <TextFieldElement
                control={control}
                name="location"
                label="Location"
                size="small"
            />

            <DatePickerElement
                control={control}
                name="start_date"
                label="Interval start:"
                inputProps={{size:"small"}}
                
            />
            <DatePickerElement
                control={control}
                name="end_date"
                label="Interval end:"
                inputProps={{size:"small"}}
                
            />

            <TextFieldElement
                type="number"
                control={control}
                name="num_guests"
                label="Number of guests"
                size="small"
                
            />
            <MultiSelectElement
                control={control}
                name="amenity"
                label="Amenity"
                options={AMENETIES}
                size="small"
            />
            <AuthShow roles={['H']}>
            <CheckboxElement
                control={control}
                name="show_my"
                label="Show my accommodations"
                size="small"
            />
            </AuthShow>
            <Button variant="contained" onClick={()=>submit()}>Filter</Button>
        </Stack>
    );
}
