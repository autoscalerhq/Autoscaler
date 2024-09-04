"use client"

import * as React from "react"
import {
    Calculator,
    Calendar,
    CreditCard,
    Settings,
    Smile,
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

export function Command() {
    const [open, setOpen] = React.useState(false)

    React.useEffect(() => {
        const down = (e: KeyboardEvent) => {
            if (e.key === "j" && (e.metaKey || e.ctrlKey)) {
                e.preventDefault()
                setOpen((open) => !open)
            }
        }

        document.addEventListener("keydown", down)
        return () => document.removeEventListener("keydown", down)
    }, [])

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
                            <CommandShortcut>⌘P</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <CreditCard className="mr-2 h-4 w-4" />
                            <span>Billing</span>
                            <CommandShortcut>⌘B</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <Settings className="mr-2 h-4 w-4" />
                            <span>Environment</span>
                            <CommandShortcut>⌘E</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <Settings className="mr-2 h-4 w-4" />
                            <span>Audit Log</span>
                            <CommandShortcut>⌘L</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <Settings className="mr-2 h-4 w-4" />
                            <span>Access Tokens</span>
                            <CommandShortcut>⌘T</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <Settings className="mr-2 h-4 w-4" />
                            <span>Members</span>
                            <CommandShortcut>⌘M</CommandShortcut>
                        </CommandItem>
                    </CommandGroup>
                    <CommandSeparator />
                    <CommandGroup heading="Env Links">
                        <CommandItem>
                            <Boxes className="mr-2 h-4 w-4" />
                            <span>Overview</span>
                            <CommandShortcut>⌘1</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <GitPullRequestDraft className="mr-2 h-4 w-4" />
                            <span>Pull</span>
                            <CommandShortcut>⌘2</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <ChevronsLeftRight className="mr-2 h-4 w-4" />
                            <span>Streaming</span>
                            <CommandShortcut>⌘3</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <BookUp className="mr-2 h-4 w-4" />
                            <span>Push</span>
                            <CommandShortcut>⌘4</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <PencilRuler className="mr-2 h-4 w-4" />
                            <span>Polices</span>
                            <CommandShortcut>⌘5</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <ArrowDownUp className="mr-2 h-4 w-4" />
                            <span>Scalers</span>
                            <CommandShortcut>⌘6</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <SquareLibrary className="mr-2 h-4 w-4" />
                            <span>Monitoring</span>
                            <CommandShortcut>⌘7</CommandShortcut>
                        </CommandItem>
                        <CommandItem>
                            <Bell className="mr-2 h-4 w-4" />
                            <span>Alerts</span>
                            <CommandShortcut>⌘8</CommandShortcut>
                        </CommandItem>
                    </CommandGroup>
                </CommandList>
            </CommandDialog>
        </>
    )
}