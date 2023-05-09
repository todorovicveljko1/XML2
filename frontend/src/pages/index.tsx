import MainLayout from "@/components/Layout/MainLayout";
import { AuthShow } from "@/providers/authProvider";
import { Container } from "@mui/material";

export default function Home() {
    return (
        <MainLayout>
            <Container maxWidth="lg">
                <AuthShow roles={["G"]}>
                    <h1>Guest</h1>
                </AuthShow>
                <AuthShow roles={["H"]}>
                    <h1>Host</h1>
                </AuthShow>
            </Container>
        </MainLayout>
    );
}
