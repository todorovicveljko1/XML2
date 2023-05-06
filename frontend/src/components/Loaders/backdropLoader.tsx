import { useDebounce } from "@/helpers/debounce";
import { Backdrop, CircularProgress, Stack, Typography } from "@mui/material";
import { useEffect, useState } from "react";

export interface BackdropLoaderProps {
    text?: string;
    wait?: number
}

export default function BackdropLoader(props: BackdropLoaderProps) {
    const [isOpen, setIsOpen] = useState(false) 
    const debouncedIsOpen = useDebounce(isOpen, props.wait ?? 300);

    useEffect(()=>{
        if(!isOpen) setIsOpen(true)
    }, [isOpen, setIsOpen])

    return (
        <Backdrop
            sx={{zIndex: (theme) => theme.zIndex.drawer + 1}}
            open={debouncedIsOpen}
        >
            <Stack spacing={1} alignItems="center">
                <CircularProgress color="primary" />
                {props.text && (
                    <Typography variant="caption" color="primary">{props.text}</Typography>
                )}
            </Stack>
        </Backdrop>
    );
}
