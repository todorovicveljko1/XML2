import { Accommodation, AvailableInterval, PriceInterval } from "@/types/accomodation";
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
import { AuthShow, useAuth } from "@/providers/authProvider";
import dayjs from "dayjs";
import { PriceIntervalForm } from "../Forms/PriceIntervalForms";

interface PriceIntervalProps {
    accommodation: Accommodation;
    intervals: PriceInterval[];
    onUpdate?: () => void;
}

export function PriceIntervalCard({
    intervals,
    accommodation,
    onUpdate,
}: PriceIntervalProps) {
    const { user } = useAuth();
    return (
        <Paper sx={{ p: 3 }}>
            <AuthShow roles={["H"]}>
            {user && user.id == accommodation.user_id && (
                    <>
                <PriceIntervalForm
                    accomodationId={accommodation.id}
                    onSuccess={onUpdate}
                />
                <Divider sx={{ my: 3 }} />
                </>
            )}
            </AuthShow>
            <Typography variant="h6" gutterBottom>
                Modified prices intervals
            </Typography>
            {intervals.length == 0 ? (
                <Alert color="info">This accommodation does not have modifued prices!</Alert>
            ) : (
                <TableContainer component={Paper}>
                    <Table sx={{ minWidth: 650 }} aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell>Start date</TableCell>
                                <TableCell>End Date</TableCell>
                                <TableCell align="right">Price</TableCell>
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
                                    <TableCell align="right">
                                        {row.price} $
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
