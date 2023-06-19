import { AuthShow, useAuth } from "@/providers/authProvider";
import { UserType } from "@/types/user";
import { Button, IconButton, Stack, Badge } from "@mui/material";
import { useRouter } from "next/router";
import LogoutIcon from "@mui/icons-material/Logout";
import NotificationIcon from "@mui/icons-material/Notifications";
import axios from "@/axios";
import { useQuery } from "react-query";

export default function UserMenu({ user }: { user: UserType }) {
    const router = useRouter();
    const { logout, isLoading } = useAuth();
    const {
        data,
        isLoading: isLoadingNotification,
        error,
    } = useQuery(["notification"], () => {
        return axios.get("/notification");
    },{
        refetchInterval: 10000,
    });
    return (
        <Stack direction={"row"} spacing={2}>
            <Button
                sx={{ color: "white" }}
                onClick={() => {
                    router.push("/auth/profile");
                }}
            >
                {user.email}
            </Button>

            <IconButton
                color="inherit"
                onClick={() => router.push("/notification")}
            >
                <Badge
                    invisible={
                        isLoadingNotification ||
                        isLoading ||
                        !data?.data ||
                        data?.data.length == 0
                    }
                    badgeContent={data?.data?.length ?? 0}
                    color="error"
                >
                    <NotificationIcon />
                </Badge>
            </IconButton>

            <IconButton color="inherit" onClick={logout}>
                <LogoutIcon />
            </IconButton>
        </Stack>
    );
}
