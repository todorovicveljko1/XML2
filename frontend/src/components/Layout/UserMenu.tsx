import { AuthShow, useAuth } from "@/providers/authProvider";
import { UserType } from "@/types/user";
import { Button, Link, Typography } from "@mui/material";
import { useRouter } from "next/router";

export default function UserMenu({ user }: {user:UserType}) {
    const { logout } = useAuth()
    return (
        <>
            <Typography fontWeight={"medium"} onClick={()=>{
                console.log(user)
            }}>{user.email}</Typography>
           
            
            <Button variant="text" color="inherit" sx={{ marginLeft: 2 }} onClick={logout}>
                Sign out
            </Button>
        </>
    );
}
