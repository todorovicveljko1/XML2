import MainLayout from "@/components/Layout/MainLayout";
import { AuthWraper } from "@/providers/authProvider";
import { Container } from "@mui/material";
import { useRouter } from "next/router";


export default function AddRating() {
    const router = useRouter();
    const { id } = router.query;
    


    return <MainLayout>
        <AuthWraper roles={['G']}>
            <Container maxWidth='xs' sx={{py:3}}>
                
            </Container>
        </AuthWraper>
    </MainLayout>

}