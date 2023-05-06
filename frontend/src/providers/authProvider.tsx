import { UserType } from "@/types/user";
import BackdropLoader from "@components/Loaders/backdropLoader";
import axios from "@/axios";
import Router from "next/router";
import { enqueueSnackbar } from "notistack";
import {
    createContext,
    ReactNode,
    useCallback,
    useContext,
    useEffect,
    useState,
} from "react";
import { useToken } from "./tokenProvider";
import { useQuery } from "react-query";
import { useDebounce } from "@helpers/debounce";

export type AuthContextType = {
    isLoading: boolean;
    user: UserType | undefined;
    isAuth: boolean;
    isGuest: boolean;
    login: (token: string) => void;
    logout: () => void;
    hasAccess: (role: string[]) => boolean;
    token: string | null;
};

export const AuthContext = createContext<AuthContextType>({
    isLoading: false,
    user: undefined,
    isAuth: false,
    isGuest: true,
    login: () => {},
    logout: () => {},
    hasAccess: () => false,
    token: null,
});
export interface AuthProviderProps {
    children: JSX.Element;
}
export default function AuthProvider(props: AuthProviderProps) {
    const {
        setToken,
        token,
        removeToken,
        hasToken,
        isLoading: isLoadingToken,
    } = useToken();
    const {
        data: user,
        isLoading: isLoadingUser,
        remove,
        status,
    } = useQuery<UserType>("/auth", { enabled: hasToken });
    const isAuth = !!user;
    const isGuest = !user;

    const hasAccess = useCallback(
        (roles: string[]) =>
            !!(user && roles.find((role) => role === user.role)),
        [user]
    );
    const logout = useCallback(() => {
        removeToken();
        remove();
    }, [removeToken, remove]);

    const login = useCallback(
        (token: string) => {
            setToken(token);
        },
        [setToken]
    );
    useEffect(() => {
        if (token && !user) {
            remove();
        }
    }, [token, remove, user]);

    const isLoading = useDebounce(isLoadingUser || isLoadingToken, 100);
    useEffect(() => {
        console.log(status);
        if (status === "error") {
            logout();
        }
    }, [status, logout]);
    return (
        <AuthContext.Provider
            value={{
                isLoading,
                user,
                isAuth,
                isGuest,
                login,
                logout,
                hasAccess,
                token,
            }}
        >
            {props.children}
        </AuthContext.Provider>
    );
}

export const useAuth = () => useContext(AuthContext);
type AuthWraperProps = {
    children?: ReactNode;
    loader?: ReactNode;
    roles?: string[];
    loose?: boolean;
};

const isGuestRedirect = () => {
    Router.push("/auth/login");
    return <BackdropLoader text="Redirecting..." />;
};
const noRoleRedirect = () => {
    Router.push("/");
    enqueueSnackbar("You don't have requred role.", {
        variant: "info",
        preventDuplicate: true,
    });
    return <BackdropLoader text="Redirecting..." />;
};

export function AuthWraper(props: AuthWraperProps) {
    const auth = useAuth();
    if (props.loose) {
        return (
            <>
                {auth.isLoading && (
                    <>
                        {props.loader ? (
                            props.loader
                        ) : (
                            <BackdropLoader text="Loading..." />
                        )}
                    </>
                )}
                {auth.isGuest && isGuestRedirect()}
                {props.roles &&
                    props.roles.length &&
                    !auth.hasAccess(props.roles) &&
                    noRoleRedirect()}
                {props.children}
            </>
        );
    }
    if (auth.isLoading) {
        return (
            <>
                {props.loader ? (
                    props.loader
                ) : (
                    <BackdropLoader text="Loading..." />
                )}
            </>
        );
    }
    if (auth.isGuest) {
        return isGuestRedirect();
    }
    if (props.roles && props.roles.length && !auth.hasAccess(props.roles)) {
        return noRoleRedirect();
    }
    return <>{props.children}</>;
}

type GuestWraperProps = Omit<AuthWraperProps, "permission">;

const isAuthRedirect = () => {
    Router.push("/");
    return <BackdropLoader text="Redirecting..." />;
};

export function GuestWraper(props: GuestWraperProps) {
    const auth = useAuth();
    if (auth.isLoading) {
        return (
            <>
                {props.loader ? (
                    props.loader
                ) : (
                    <BackdropLoader text="Loading..." />
                )}
            </>
        );
    }
    if (auth.isAuth) {
        return isAuthRedirect();
    }
    return <>{props.children}</>;
}

type AuthShowProps = {
    children: ReactNode;
    loader?: ReactNode;
    roles?: string[];
};
export function AuthShow(props: AuthShowProps) {
    const auth = useAuth();
    if (auth.isLoading) {
        return <>{props.loader ? props.loader : null}</>;
    }
    if (auth.isGuest) {
        return null;
    }
    if (props.roles && props.roles.length && !auth.hasAccess(props.roles)) {
        return null;
    }
    return <>{props.children}</>;
}
