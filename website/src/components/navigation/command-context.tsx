"use client"
import React, { useState, useEffect, createContext, ReactNode, useContext } from "react";
import {useRouter} from "next/navigation";

interface CommandContextProps {
    open: boolean;
    setOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

const CommandContext = createContext<CommandContextProps | undefined>(undefined);

interface ICommanderProvider {
    children: ReactNode;
    org: string;
    env: string;
}

export const CommandProvider = ({ children, org, env }: ICommanderProvider ) => {
    const [open, setOpen] = useState(false);

    const router = useRouter();

    const org_link = `/${org}/`;
    const base_link = `/${org}/${env}/`;
    const settings_link = `/${org}/settings/`;

    useEffect(() => {
        const down = (e: KeyboardEvent) => {
            if (e.metaKey || e.ctrlKey) {
                switch (e.key.toLowerCase()) {
                    case "j":
                        e.preventDefault();
                        setOpen((open) => !open);
                        break;
                    case "r":
                        e.preventDefault();
                        // Add your logic to open Profile
                        console.log("Open Profile");
                        // router.push('/profile');
                        setOpen(false);
                        break;
                    case "b":
                        e.preventDefault();
                        router.push(settings_link+'billing');
                        setOpen(false);
                        break;
                    case "e":
                        e.preventDefault();
                        router.push(settings_link+'environments');
                        setOpen(false);
                        break;
                    case "l":
                        e.preventDefault();
                        router.push(settings_link+'audit-logs');
                        setOpen(false);
                        break;
                    case "t":
                        e.preventDefault();
                        router.push(settings_link+'access-tokens');
                        setOpen(false);
                        break;
                    case "m":
                        e.preventDefault();
                        router.push(settings_link+'members');
                        setOpen(false);
                        break;
                    case "delete":
                    case "backspace":
                        e.preventDefault();
                        router.back();
                        break;
                }
            }
        };

        document.addEventListener("keydown", down);
        return () => document.removeEventListener("keydown", down);
    }, []);

    return (
        <CommandContext.Provider value={{ open, setOpen }}>
            {children}
        </CommandContext.Provider>
    );
};

export const useCommand = () => {
    const context = useContext(CommandContext);
    if (!context) {
        throw new Error("useCommand must be used within a CommandProvider");
    }
    return context;
};