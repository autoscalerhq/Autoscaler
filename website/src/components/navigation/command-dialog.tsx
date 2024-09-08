"use client"

import * as React from "react"
import {
    CreditCard,
    Settings,
    User,
    ArrowDownUp,
    Bell,
    BookUp,
    Boxes,
    ChevronsLeftRight,
    GitPullRequestDraft,
    PencilRuler,
    SquareLibrary
} from "lucide-react"

import {
    CommandDialog,
    CommandEmpty,
    CommandGroup,
    CommandInput,
    CommandItem,
    CommandList,
    CommandSeparator,
    CommandShortcut,
} from "~/components/ui/command"

import { useCommand } from "~/components/navigation/command-context"

export function Command() {

    const { open, setOpen } = useCommand();

    return (
        <>
            <CommandDialog open={open} onOpenChange={setOpen} title={"test"}>

                <CommandInput placeholder="Type a command or search..." onClick={() => {console.log("input click")}}/>

                <CommandList>

                    <CommandEmpty>No results found.</CommandEmpty>

                    <CommandGroup heading="Settings">
                        <CommandItem>
                            <User className="mr-2 h-4 w-4" />
                            <span>Profile</span>
                            <CommandShortcut>⌘ + R</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <CreditCard className="mr-2 h-4 w-4" />
                            <span>Billing</span>
                            <CommandShortcut>⌘ + B</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <Settings className="mr-2 h-4 w-4" />
                            <span>Environment</span>
                            <CommandShortcut>⌘ + E</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <Settings className="mr-2 h-4 w-4" />
                            <span>Audit Log</span>
                            <CommandShortcut>⌘ + L</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <Settings className="mr-2 h-4 w-4" />
                            <span>Access Tokens</span>
                            <CommandShortcut>⌘ + T</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <Settings className="mr-2 h-4 w-4" />
                            <span>Members</span>
                            <CommandShortcut>⌘ + M</CommandShortcut>
                        </CommandItem>
                    </CommandGroup>
                    <CommandSeparator />
                    <CommandGroup heading="Env Links">
                        <CommandItem>
                            <Boxes className="mr-2 h-4 w-4" />
                            <span>Overview</span>
                            <CommandShortcut>⌘ + 1</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <GitPullRequestDraft className="mr-2 h-4 w-4" />
                            <span>Inputs</span>
                            <CommandShortcut>⌘ + 2</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <PencilRuler className="mr-2 h-4 w-4" />
                            <span>Polices</span>
                            <CommandShortcut>⌘ + 3</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <ArrowDownUp className="mr-2 h-4 w-4" />
                            <span>Scalers</span>
                            <CommandShortcut>⌘ + 4</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <Bell className="mr-2 h-4 w-4" />
                            <CommandShortcut>Actions</CommandShortcut>
                            <CommandShortcut>⌘ + 5</CommandShortcut>
                        </CommandItem>
                    </CommandGroup>
                </CommandList>
            </CommandDialog>
        </>
    );
}