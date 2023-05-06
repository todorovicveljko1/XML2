import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import { AuthShow, useAuth } from "@/providers/authProvider";
import UserMenu from "./UserMenu";
import Link from "next/link";
import { useRouter } from "next/router";
import { Stack } from "@mui/material";

export interface MainLayoutProps {
    children?: React.ReactNode;
}
export default function MainLayout({ children }: MainLayoutProps) {
    const router = useRouter();
    const { user, isLoading } = useAuth();
    return (
        <Box>
            <AppBar position="fixed" elevation={1} enableColorOnDark>
                <Toolbar>
                    <Stack sx={{ flexGrow: 1 }} direction="row" spacing={3}>
                        <Typography variant="h6">Accomodation App</Typography>
                    </Stack>
                    {isLoading && <span>Loading...</span>}
                    {!isLoading &&
                        (user ? (
                            <UserMenu user={user} />
                        ) : (
                            <>
                                <Link href="/auth/login">
                                    <Button sx={{ color: "white" }}>
                                        Login
                                    </Button>
                                </Link>
                                <Link href="/auth/register">
                                    <Button sx={{ color: "white" }}>
                                        Register
                                    </Button>
                                </Link>
                            </>
                        ))}
                </Toolbar>
            </AppBar>
            <Box component="main">
                <Toolbar />
                {children}
            </Box>
        </Box>
    );
}
