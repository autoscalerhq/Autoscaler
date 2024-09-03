"use client"

import {usePathname} from "next/navigation";
import {Nav, NavLinks} from "~/components/navigation/vertical-nav";
import {ArrowRightLeft, BookDashed, Search, Settings, SquareDashedKanban} from "lucide-react";

export default function OrgNav({org, env}: { org: string; env: string }) {
    const pathname = usePathname();

    const org_link = `/${org}/`;
    const base_link = `/${org}/${env}/`;

    const links: NavLinks[] = ([
        {
            title: "Settings",
            href: org_link + "settings",
            icon: Settings ,
            variant: "ghost",
        },
        {
            title: "Integrations",
            href: base_link + "integrations",
            icon: ArrowRightLeft ,
            variant: "ghost",
        },
        {
            title: "System Templates",
            href: base_link + "system-templates",
            icon: BookDashed,
            variant: "ghost",
        },
        {
            title: "Service Templates",
            href: base_link + "service-templates",
            icon: SquareDashedKanban,
            variant: "ghost",
        },
        {
            title: "Search",
            href: "",
            label: "CMD + K",
            icon: Search,
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