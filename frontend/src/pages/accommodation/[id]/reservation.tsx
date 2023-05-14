import axios from "@/axios";
import MainLayout from "@/components/Layout/MainLayout";
import BackdropLoader from "@/components/Loaders/backdropLoader";
import {
    Alert,
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
import { useMutation, useQuery } from "react-query";

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

export default function Reservations() {
    const router = useRouter();
    const { id } = router.query;

    const { data, isLoading, error, refetch } = useQuery(
        ["accommodation", id],
        () => {
            return axios.get(`/accommodation/${id}/reservation`);
        },
        { enabled: !!id }
    );
    const reservations = data?.data.reservations ?? [];

    const mutation = useMutation(
        (data: { id: string; status: string }) =>
            axios.put(`/accommodation/${id}/reservation/${data.id}`, {
                status: data.status,
            }),
        {
            onSuccess: () => {
                enqueueSnackbar("Reservation updated", { variant: "success" });
                refetch();
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
            <Container maxWidth="md" sx={{ py: 3 }}>
                {isLoading ? (
                    <BackdropLoader />
                ) : error ? (
                    <Alert severity="error">
                        Error while getting reservations
                    </Alert>
                ) : (
                    data &&
                    data.data.reservations && (
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
                                                <TableCell>
                                                    Start date
                                                </TableCell>
                                                <TableCell>End Date</TableCell>
                                                <TableCell>
                                                    Number of Guests
                                                </TableCell>
                                                <TableCell>Price</TableCell>
                                                <TableCell>Status</TableCell>

                                                <TableCell>Actions</TableCell>
                                            </TableRow>
                                        </TableHead>
                                        <TableBody>
                                            {reservations.map((row: any) => (
                                                <TableRow
                                                    hover
                                                    key={row.start_date}
                                                    sx={{
                                                        "&:last-child td, &:last-child th":
                                                            {
                                                                border: 0,
                                                            },
                                                    }}
                                                    onClick={() =>
                                                        console.log(row)
                                                    }
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
                                                        {row.number_of_guests}
                                                    </TableCell>
                                                    <TableCell>
                                                        {row.price} $
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
                                                        {row.status ==
                                                            "PENDING" && (
                                                            <>
                                                                <Button
                                                                    variant="contained"
                                                                    color="success"
                                                                    onClick={() =>
                                                                        mutation.mutate(
                                                                            {
                                                                                id: row.id,
                                                                                status: "APPROVED",
                                                                            }
                                                                        )
                                                                    }
                                                                >
                                                                    Approve
                                                                </Button>
                                                                <Button
                                                                    variant="contained"
                                                                    color="error"
                                                                    sx={{
                                                                        ml: 1,
                                                                    }}
                                                                    onClick={() =>
                                                                        mutation.mutate(
                                                                            {
                                                                                id: row.id,
                                                                                status: "REJECTED",
                                                                            }
                                                                        )
                                                                    }
                                                                >
                                                                    Reject
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
                    )
                )}
            </Container>
        </MainLayout>
    );
}
