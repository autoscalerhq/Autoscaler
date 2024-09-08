"use client"

import {usePathname} from "next/navigation";
import {Nav, NavLinks} from "~/components/navigation/vertical-nav";
import {ArrowRightLeft, BookDashed, Search, Settings, SquareDashedKanban} from "lucide-react";
import {CommandShortcut} from "~/components/ui/command";
import * as console from "node:console";

import { useCommand } from "~/components/navigation/command-context"

export default function OrgNav({org, env}: { org: string; env: string }) {
    const pathname = usePathname();

    const org_link = `/${org}/`;
    const base_link = `/${org}/${env}/`;

    const { open, setOpen } = useCommand();

    const links: NavLinks[] = ([
        {
            title: "Settings",
            href: org_link + "settings",
            icon: Settings,
            variant: "ghost",
        },
        {
            title: "Integrations",
            href: base_link + "integrations",
            icon: ArrowRightLeft,
            variant: "ghost",
        },
        {
            title: "Search",
            href: "",
            label: <CommandShortcut>⌘ + J</CommandShortcut>,
            icon: Search,
            variant: "ghost",
            type: "action",
            action: () => setOpen(true)
        },
    ].map(link => ({
        ...link,
        // Set variant based on condition
        variant: link.type === "action" ? "ghost" : (pathname === link.href) ? "default" : "ghost"
    })) as NavLinks[]);

    return (
        <Nav
            isCollapsed={false}
            links={links}
        />
    );
}