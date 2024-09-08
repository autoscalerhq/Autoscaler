"use client"

import * as React from "react"
import {
    CreditCard,
    Settings,
    User,
    ArrowDownUp,
    Bell,
    Boxes,
    GitPullRequestDraft,
    PencilRuler,
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

import { useCommand } from "~/components/navigation/command-context";
import {DialogFooter} from "~/components/ui/dialog";
import {Kbd} from "~/components/ui/kbd";

export function Command() {

    const { open, setOpen, selected, setSelected } = useCommand();

    return (
        <>
            <CommandDialog open={open} onOpenChange={ (open) => {
                setOpen(open)
                if(!open) {
                    setSelected("")
                }
            }} title={"test"} >

                <CommandInput placeholder="Type a command or search..." onClick={() => {console.log("input click")}}/>

                <CommandList>

                    <CommandEmpty>No results found.</CommandEmpty>

                    <CommandGroup heading="Settings">
                        <CommandItem
                            onSelect={(value) => {setSelected("r") }}
                        >
                            <User className="mr-2 h-4 w-4" />
                            <span>Profile</span>
                            <CommandShortcut><Kbd>⌘</Kbd>  <Kbd>r</Kbd></CommandShortcut>
                        </CommandItem>
                        <CommandItem
                            onSelect={(value) => {setSelected("b") }}
                        >
                            <CreditCard className="mr-2 h-4 w-4" />
                            <span>Billing</span>
                            <CommandShortcut><Kbd>⌘</Kbd> <Kbd>b</Kbd></CommandShortcut>
                        </CommandItem>
                        <CommandItem
                            onSelect={(value) => {setSelected("e") }}
                        >
                            <Settings className="mr-2 h-4 w-4" />
                            <span>Environment</span>
                            <CommandShortcut><Kbd>⌘</Kbd> <Kbd>e</Kbd></CommandShortcut>
                        </CommandItem>
                        <CommandItem
                            onSelect={(value) => {setSelected("l") }}
                        >
                            <Settings className="mr-2 h-4 w-4" />
                            <span>Audit Log</span>
                            <CommandShortcut><Kbd>⌘</Kbd> <Kbd>l</Kbd></CommandShortcut>
                        </CommandItem>
                        <CommandItem
                            onSelect={(value) => {setSelected("t") }}
                        >
                            <Settings className="mr-2 h-4 w-4" />
                            <span>Access Tokens</span>
                            <CommandShortcut><Kbd>⌘</Kbd> <Kbd>t</Kbd></CommandShortcut>
                        </CommandItem>
                        <CommandItem
                            onSelect={(value) => {setSelected("m") }}
                        >
                            <Settings className="mr-2 h-4 w-4" />
                            <span>Members</span>
                            <CommandShortcut><Kbd>⌘</Kbd> <Kbd>m</Kbd></CommandShortcut>
                        </CommandItem>
                    </CommandGroup>
                    <CommandSeparator />
                    <CommandGroup heading="Env Links">
                        <CommandItem
                            onSelect={(value) => {setSelected("1") }}
                        >
                            <Boxes className="mr-2 h-4 w-4" />
                            <span>Overview</span>
                            <CommandShortcut><Kbd>⌘</Kbd> <Kbd>1</Kbd></CommandShortcut>
                        </CommandItem>
                        <CommandItem
                            onSelect={(value) => {setSelected("2") }}
                        >
                            <GitPullRequestDraft className="mr-2 h-4 w-4" />
                            <span>Inputs</span>
                            <CommandShortcut><Kbd>⌘</Kbd> <Kbd>2</Kbd></CommandShortcut>
                        </CommandItem>
                        <CommandItem
                            onSelect={(value) => {setSelected("3") }}
                        >
                            <PencilRuler className="mr-2 h-4 w-4" />
                            <span>Polices</span>
                            <CommandShortcut><Kbd>⌘</Kbd> <Kbd>3</Kbd></CommandShortcut>
                        </CommandItem>
                        <CommandItem
                            onSelect={(value) => {setSelected("4") }}
                        >
                            <ArrowDownUp className="mr-2 h-4 w-4" />
                            <span>Scalers</span>
                            <CommandShortcut><Kbd>⌘</Kbd> <Kbd>4</Kbd></CommandShortcut>
                        </CommandItem>
                        <CommandItem
                            onSelect={(value) => {setSelected("5") }}
                        >
                            <Bell className="mr-2 h-4 w-4" />
                            <CommandShortcut>Actions</CommandShortcut>
                            <CommandShortcut><Kbd>⌘</Kbd>  <Kbd>5</Kbd></CommandShortcut>
                        </CommandItem>
                    </CommandGroup>
                </CommandList>

                <DialogFooter className="flex flex-col text-sm border border-t-1 ">

                    <div className="flex items-center m-2 text-sm text-gray-500 mr-4">
                        <span>Go to</span>
                        <Kbd className={"ml-2 text-bold text-md"}>&#x23CE;</Kbd>
                    </div>

                </DialogFooter>

            </CommandDialog>
        </>
    );
}