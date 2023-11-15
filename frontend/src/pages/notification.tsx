import axios from "@/axios";
import MainLayout from "@/components/Layout/MainLayout";
import BackdropLoader from "@/components/Loaders/backdropLoader";
import { AuthWraper } from "@/providers/authProvider";
import {
    Alert,
    Button,
    Container,
    Divider,
    Paper,
    Stack,
    Typography,
} from "@mui/material";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import { useQuery } from "react-query";

type NotificationType = {
    id: string;
    type: string;
    resource_id: string;
    body: string;
    user_id: string;
    is_read: boolean;
    created_at: string;
};

dayjs.extend(relativeTime);

export default function Notification() {
    const { data, isLoading, error, refetch } = useQuery(
        ["notification"],
        () => {
            return axios.get("/notification");
        }
    );

    const mutate = (id: string) => {
        axios.get("/notification/" + id + "/read").then((res) => {
            console.log(res);
            refetch();
        });
    };

    return (
        <MainLayout>
            <AuthWraper>
                <Container maxWidth="lg" sx={{ py: 3 }}>
                    {isLoading ? (
                        <BackdropLoader />
                    ) : error ? (
                        <Alert severity="error">
                            Error while getting reservations
                        </Alert>
                    ) : !data?.data ? (
                        <Alert severity="info">No notifications</Alert>
                    ) : (
                        <Paper sx={{ p: 3 }}>
                            <Stack spacing={2}>
                                <Stack>
                                    <Typography variant="h4" component="h1">
                                        Notifications
                                    </Typography>
                                </Stack>

                                {data?.data?.map(
                                    (notification: NotificationType) => (
                                        <Stack key={notification.id}>
                                            <Stack
                                                direction={"row"}
                                                spacing={3}
                                                alignItems={"start"}
                                                justifyContent={"space-between"}
                                            >
                                                <Typography>
                                                    {notification.body}
                                                </Typography>
                                                <Stack sx={{ width: "10rem" }}>
                                                    <Button
                                                        size="small"
                                                        color="error"
                                                        onClick={() =>
                                                            mutate(
                                                                notification.id
                                                            )
                                                        }
                                                    >
                                                        Dismiss
                                                    </Button>
                                                    <Typography
                                                        fontWeight={"bold"}
                                                        noWrap={true}
                                                    >
                                                        {dayjs().to(
                                                            dayjs(
                                                                notification.created_at,
                                                                "YYYY-MM-DD HH:mm:ss.SSS Z [UTC]"
                                                            )
                                                        )}
                                                    </Typography>
                                                </Stack>
                                            </Stack>
                                            <Divider sx={{ mt: 1 }} />
                                        </Stack>
                                    )
                                )}
                            </Stack>
                        </Paper>
                    )}
                </Container>
            </AuthWraper>
        </MainLayout>
    );
}
