"use client"
import React, { useState, useEffect, createContext, ReactNode, useContext } from "react";

interface CommandContextProps {
    open: boolean;
    setOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

const CommandContext = createContext<CommandContextProps | undefined>(undefined);

export const CommandProvider = ({ children }: { children: ReactNode }) => {
    const [open, setOpen] = useState(false);

    useEffect(() => {
        const down = (e: KeyboardEvent) => {
            if (e.key === "j" && (e.metaKey || e.ctrlKey)) {
                e.preventDefault();
                setOpen((open) => !open);
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