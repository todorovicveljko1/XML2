import axios from "@/axios";
import createEmotionCache from "@/createEmotionCache";
import AuthProvider from "@providers/authProvider";
import TokenProvider from "@providers/tokenProvider";
import theme from "@/theme";
import { CacheProvider, EmotionCache, ThemeProvider } from "@emotion/react";
import { CssBaseline } from "@mui/material";
import { AppProps } from "next/app";
import Head from "next/head";
import { SnackbarProvider } from "notistack";
import { QueryClient, QueryClientProvider } from "react-query";
import { LocalizationProvider } from "@mui/x-date-pickers";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import 'dayjs/locale/en-gb';

const clientSideEmotionCache = createEmotionCache();

export interface MyAppProps extends AppProps {
    emotionCache?: EmotionCache;
}

const queryClient = new QueryClient({
    defaultOptions: {
        queries: {
            staleTime:5000,
            queryFn: async ({ queryKey }: any) =>
                (await axios.get(queryKey[0])).data,
        },
    },
});

export default function MyApp(props: MyAppProps) {
    const {
        Component,
        emotionCache = clientSideEmotionCache,
        pageProps,
    } = props;
    return (
        <CacheProvider value={emotionCache}>
            <Head>
                <meta
                    name="viewport"
                    content="initial-scale=1, width=device-width"
                />
            </Head>
            <ThemeProvider theme={theme}>
                <CssBaseline />

                <LocalizationProvider dateAdapter={AdapterDayjs} adapterLocale='en-gb'>
                    <QueryClientProvider client={queryClient}>
                        <SnackbarProvider>
                            <TokenProvider>
                                <AuthProvider>
                                    <Component {...pageProps} />
                                </AuthProvider>
                            </TokenProvider>
                        </SnackbarProvider>
                    </QueryClientProvider>
                </LocalizationProvider>
            </ThemeProvider>
        </CacheProvider>
    );
}
