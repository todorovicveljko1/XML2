import axios from "@/axios";
import MainLayout from "@/components/Layout/MainLayout";
import { AuthWraper } from "@/providers/authProvider";
import {
    Button,
    Chip,
    Container,
    Divider,
    Paper,
    Stack,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    Typography,
} from "@mui/material";
import dayjs from "dayjs";
import { useRouter } from "next/router";
import { enqueueSnackbar } from "notistack";
import { useMutation } from "react-query";

const reservations = [
    {
        start_date: "2021-10-01 00:00:00",
        end_date: "2021-10-10 00:00:00",
        status: "PENDING",
    },
    {
        start_date: "2021-10-15 00:00:00",
        end_date: "2021-10-20 00:00:00",
        status: "APPROVED",
    },
];
//PENDING, APPROVED, REJECTED, CANCELED, AUTO_REJECTED
const STATUS_TO_COLOR: Record<
    string,
    "info" | "success" | "error" | "warning"
> = {
    PENDING: "info",
    APPROVED: "success",
    REJECTED: "error",
    CANCELED: "warning",
    AUTO_REJECTED: "error",
};

function ShowCancleCondition(reservation: any) {
    return (
        reservation.status == "PENDING" ||
        (reservation.status == "APPROVED" &&
            dayjs(reservation.start_date.split(" ")[0]) <
                dayjs().add(-1, "day"))
    );
}

export default function Reservations() {
    const router = useRouter();
    const { id } = router.query;
    const mutation = useMutation(
        (data: { id: string; status: string }) =>
            axios.put(`/accommodation/${id}/reservation/${data.id}`, {
                status: data.status,
            }),
        {
            onSuccess: () => {
                enqueueSnackbar("Reservation updated", { variant: "success" });
            },
            onError: () => {
                enqueueSnackbar("Error while updating reservation", {
                    variant: "error",
                });
            },
        }
    );

    return (
        <MainLayout>
            <AuthWraper roles={["G"]}>
                <Container maxWidth="md" sx={{ py: 3 }}>
                    <Paper sx={{ p: 3 }}>
                        <Stack spacing={3}>
                            <Typography
                                variant="h4"
                                component="h1"
                                gutterBottom
                            >
                                Reservations
                            </Typography>
                            <TableContainer component={Paper}>
                                <Table
                                    sx={{ minWidth: 650 }}
                                    aria-label="simple table"
                                >
                                    <TableHead>
                                        <TableRow>
                                            <TableCell>Start date</TableCell>
                                            <TableCell>End Date</TableCell>
                                            <TableCell>Status</TableCell>
                                            <TableCell>Actions</TableCell>
                                        </TableRow>
                                    </TableHead>
                                    <TableBody>
                                        {reservations.map((row) => (
                                            <TableRow
                                                hover
                                                key={row.start_date}
                                                sx={{
                                                    "&:last-child td, &:last-child th":
                                                        {
                                                            border: 0,
                                                        },
                                                }}
                                                onClick={() => console.log(row)}
                                            >
                                                <TableCell>
                                                    {dayjs(
                                                        row.start_date.split(
                                                            " "
                                                        )[0]
                                                    ).format("DD/MM/YYYY")}
                                                </TableCell>
                                                <TableCell>
                                                    {dayjs(
                                                        row.end_date.split(
                                                            " "
                                                        )[0]
                                                    ).format("DD/MM/YYYY")}
                                                </TableCell>
                                                <TableCell>
                                                    <Chip
                                                        color={
                                                            STATUS_TO_COLOR[
                                                                row.status
                                                            ]
                                                        }
                                                        label={row.status}
                                                    />
                                                </TableCell>
                                                <TableCell>
                                                    {ShowCancleCondition(
                                                        row
                                                    ) && (
                                                        <>
                                                            <Button
                                                                variant="contained"
                                                                color="success"
                                                                onClick={() =>
                                                                    mutation.mutate(
                                                                        {
                                                                            id: "1",
                                                                            status: "CANCELLED",
                                                                        }
                                                                    )
                                                                }
                                                            >
                                                                Cancel
                                                            </Button>
                                                        </>
                                                    )}
                                                </TableCell>
                                            </TableRow>
                                        ))}
                                    </TableBody>
                                </Table>
                            </TableContainer>
                            <Stack direction="row">
                                <Button onClick={() => router.back()}>
                                    Back
                                </Button>
                            </Stack>
                        </Stack>
                    </Paper>
                </Container>
            </AuthWraper>
        </MainLayout>
    );
}
