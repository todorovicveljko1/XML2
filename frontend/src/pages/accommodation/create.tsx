import { CreateAccomodationForm } from "@/components/Forms/CreateAccomodationForm";
import MainLayout from "@/components/Layout/MainLayout";
import { AuthWraper } from "@/providers/authProvider";
import { Container, Paper } from "@mui/material";

export default function CreateAccomodation() {
    return (
        <MainLayout>
            <AuthWraper roles={["H"]}>
                <Container maxWidth="md" sx={{ py: 3 }}>
                    <Paper>
                        <CreateAccomodationForm />
                    </Paper>
                </Container>
            </AuthWraper>
        </MainLayout>
    );
}
