import { AuthShow, useAuth } from "@/providers/authProvider";
import { UserType } from "@/types/user";
import { Button, Link, Typography } from "@mui/material";
import { useRouter } from "next/router";

export default function UserMenu({ user }: { user: UserType }) {
    const router = useRouter();
    const { logout } = useAuth();
    return (
        <>
            <Button
                sx={{ color: "white" }}
                onClick={() => {
                    router.push("/auth/profile");
                }}
            >
                {user.email}
            </Button>

            <Button
                variant="text"
                color="inherit"
                sx={{ marginLeft: 2 }}
                onClick={logout}
            >
                Sign out
            </Button>
        </>
    );
}
