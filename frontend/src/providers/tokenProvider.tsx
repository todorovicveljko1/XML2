import axios from "@/axios";
import {
    createContext,
    useCallback,
    useContext,
    useEffect,
    useState,
} from "react";

export type TokenContextType = {
    hasToken: boolean;
    isLoading: boolean;
    setToken: (token: string) => void;
    token: string | null;
    removeToken: () => void;
};

const TokenContext = createContext<TokenContextType>({
    hasToken: false,
    isLoading: false,
    setToken: (token: string) => {},
    token: null,
    removeToken: () => {},
});

export interface TokenProviderProps {
    children: JSX.Element;
}

export const TOKEN_KEY = "__TOKEN";

export default function TokenProvider(props: TokenProviderProps) {
    const [hasToken, setHasToken] = useState(false);
    const [isLoading, setIsLoading] = useState(true);
    const [token, _setToken] = useState<null | string>(null);
    const setToken = useCallback(
        (token: string) => {
            setHasToken(false)
            localStorage.setItem(TOKEN_KEY, token);
            _setToken(token);
        },
        [_setToken]
    );
    const removeToken = useCallback(() => {
        _setToken(null);
        setHasToken(false)
        localStorage.removeItem(TOKEN_KEY);
    }, [_setToken]);

    useEffect(() => {
        const t = localStorage.getItem(TOKEN_KEY);
        _setToken(t && t !== "null" ? t : null);
        setIsLoading(false);
    }, [_setToken, setIsLoading]);
    useEffect(() => {
        if (!token) return;
        if (!isLoading) setHasToken(true);
        const interceptor = axios.interceptors.request.use(
            function (config) {
                if (config.headers) {
                    config.headers.Authorization = `Bearer ${token}`;
                }
                return config;
            },
            (error) => {
                return Promise.reject(error);
            }
        );
        return () => {
            axios.interceptors.request.eject(interceptor);
        };
    }, [token, isLoading]);

    return (
        <TokenContext.Provider
            value={{ hasToken, isLoading, token, setToken, removeToken }}
        >
            {props.children}
        </TokenContext.Provider>
    );
}

export const useToken = () => useContext(TokenContext);
export const useTokenCheck = () => {
    const { hasToken } = useToken();
    return hasToken;
};
