"use client"

import {usePathname} from "next/navigation";
import {Nav, NavLinks} from "~/components/navigation/vertical-nav";
import {
    Bell,
    BookUp,
    Boxes,
    ChevronsLeftRight,
    GitPullRequestDraft,
    LucideLayoutDashboard,
    PencilRuler
} from "lucide-react";

export default function EnvNav({org, env}: { org: string; env: string }) {
    const pathname = usePathname();

    const base_link = `/${org}/${env}/`;

    const links: NavLinks[] = ([
        {
            title: "Overview",
            href: base_link + "overview",
            icon: Boxes,
            variant: "ghost",
        },
        {
            title: "Pull",
            href: base_link + "pull",
            icon: GitPullRequestDraft,
            variant: "ghost",
        },
        {
            title: "Streaming",
            href: base_link + "stream",
            icon: ChevronsLeftRight,
            variant: "ghost",
        },
        {
            title: "Push",
            href: base_link + "push",
            icon: BookUp,
            variant: "ghost",
        },
        {
            title: "Polices",
            href: base_link + "policies",
            icon: PencilRuler,
            variant: "ghost",
        },
        {
            title: "Analytics",
            href: base_link + "analytics",
            icon: LucideLayoutDashboard,
            variant: "ghost",
        },
        {
            title: "Alerts",
            href: base_link + "alerts",
            icon: Bell,
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