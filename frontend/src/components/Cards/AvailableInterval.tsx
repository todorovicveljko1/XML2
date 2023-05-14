import { Accommodation, AvailableInterval } from "@/types/accomodation";
import {
    Alert,
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
import { AvailableIntervalForm } from "../Forms/AvailableIntervalsForm";
import { AuthShow, useAuth } from "@/providers/authProvider";
import dayjs from "dayjs";

interface AvailableIntervalProps {
    accommodation: Accommodation;
    intervals: AvailableInterval[];
    onUpdate?: () => void;
}

export function AvailableInterval({
    intervals,
    accommodation,
    onUpdate,
}: AvailableIntervalProps) {
    const { user } = useAuth();
    return (
        <Paper sx={{ p: 3 }}>
            <AuthShow roles={["H"]}>
                {user && user.id == accommodation.user_id && (
                    <>
                        <AvailableIntervalForm
                            accomodationId={accommodation.id}
                            onSuccess={onUpdate}
                        />
                        <Divider sx={{ my: 3 }} />
                    </>
                )}
            </AuthShow>
            <Typography variant="h6" gutterBottom>
                Unavailable intervals
            </Typography>
            <Typography variant="body2" color="text.secondary" gutterBottom>
                Intervals when host is not rentting accommodation.
            </Typography>
            {intervals.length == 0 ? (
                <Alert color="info">This accommodation is available!</Alert>
            ) : (
                <TableContainer component={Paper}>
                    <Table sx={{ minWidth: 650 }} aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell>Start date</TableCell>
                                <TableCell>End Date</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {intervals.map((row) => (
                                <TableRow
                                    key={row.start_date}
                                    sx={{
                                        "&:last-child td, &:last-child th": {
                                            border: 0,
                                        },
                                    }}
                                >
                                    <TableCell>
                                        {dayjs(
                                            row.start_date.split(" ")[0]
                                        ).format("DD/MM/YYYY")}
                                    </TableCell>
                                    <TableCell>
                                        {dayjs(
                                            row.end_date.split(" ")[0]
                                        ).format("DD/MM/YYYY")}
                                    </TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
            )}
        </Paper>
    );
}
