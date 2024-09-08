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
            <CommandDialog open={open} onOpenChange={setOpen}>
                <CommandInput placeholder="Type a command or search..." />
                <CommandList>

                    <CommandEmpty>No results found.</CommandEmpty>

                    <CommandGroup heading="Settings">
                        <CommandItem>
                            <User className="mr-2 h-4 w-4" />
                            <span>Profile</span>
                            <CommandShortcut>⌘ P</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <CreditCard className="mr-2 h-4 w-4" />
                            <span>Billing</span>
                            <CommandShortcut>⌘ B</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <Settings className="mr-2 h-4 w-4" />
                            <span>Environment</span>
                            <CommandShortcut>⌘ E</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <Settings className="mr-2 h-4 w-4" />
                            <span>Audit Log</span>
                            <CommandShortcut>⌘ L</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <Settings className="mr-2 h-4 w-4" />
                            <span>Access Tokens</span>
                            <CommandShortcut>⌘ T</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <Settings className="mr-2 h-4 w-4" />
                            <span>Members</span>
                            <CommandShortcut>⌘ M</CommandShortcut>
                        </CommandItem>
                    </CommandGroup>
                    <CommandSeparator />
                    <CommandGroup heading="Env Links">
                        <CommandItem>
                            <Boxes className="mr-2 h-4 w-4" />
                            <span>Overview</span>
                            <CommandShortcut>⌘ 1</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <GitPullRequestDraft className="mr-2 h-4 w-4" />
                            <span>Pull</span>
                            <CommandShortcut>⌘ 2</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <ChevronsLeftRight className="mr-2 h-4 w-4" />
                            <span>Streaming</span>
                            <CommandShortcut>⌘ 3</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <BookUp className="mr-2 h-4 w-4" />
                            <CommandShortcut>Push</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <PencilRuler className="mr-2 h-4 w-4" />
                            <CommandShortcut>Polices</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <ArrowDownUp className="mr-2 h-4 w-4" />
                            <CommandShortcut>Scalers</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <SquareLibrary className="mr-2 h-4 w-4" />
                            <CommandShortcut>Monitoring</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <Bell className="mr-2 h-4 w-4" />
                            <CommandShortcut>Alerts</CommandShortcut>
                        </CommandItem>
                    </CommandGroup>
                </CommandList>
            </CommandDialog>
        </>
    );
}