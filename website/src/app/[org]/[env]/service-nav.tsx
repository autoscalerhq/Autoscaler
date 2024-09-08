"use client"

import {usePathname} from "next/navigation";
import {Nav, NavLinks} from "~/components/navigation/vertical-nav";
import {
    ArrowDownUp,
    Bell,
    BookUp,
    Boxes,
    ChevronsLeftRight,
    GitPullRequestDraft,
    LucideLayoutDashboard,
    PencilRuler, SquareLibrary
} from "lucide-react";

export default function ServiceNav({org, env}: { org: string; env: string }) {
    const pathname = usePathname();

    const base_link = `/${org}/${env}/`;

    const links: NavLinks[] = ([
        {
            title: "Actions",
            href: base_link + "actions",
            icon: Bell,
            variant: "ghost",
        },
        {
            title: "Overview",
            href: base_link + "overview",
            icon: Boxes,
            variant: "ghost",
        },
        {
            title: "Analytics",
            href: base_link + "analytics",
            icon: SquareLibrary,
            variant: "ghost",
        },
        {
            title: "Inputs",
            href: base_link + "inputs",
            icon: GitPullRequestDraft,
            variant: "ghost",
        },
        {
            title: "Polices",
            href: base_link + "policies",
            icon: PencilRuler,
            variant: "ghost",
        },

        {
            title: "Scalers",
            href: base_link + "scalers",
            icon: ArrowDownUp,
            variant: "ghost",
        },


    ].map(link => ({
        ...link,
        variant: (pathname === link.href) ? "default" : "ghost",
    })));

    return (
        <Nav
            isCollapsed={false}
            links={links}
        />
    );
}