"use client"

import {usePathname} from "next/navigation";
import {Nav, NavLinks} from "~/components/navigation/vertical-nav";

export default function SettingsNav({org}: { org: string }) {
    const pathname = usePathname();
    const setting_link = `/${org}/settings/`;

    const links: NavLinks[] = ([
        {
            title: "General",
            href: setting_link,
        },
        {
            title: "Environments",
            href: setting_link + "environments",
        },
        {
            title: "Members",
            href: setting_link + "members",
        },
        {
            title: "Security",
            href: setting_link + "security",
        },
        {
            title: "Audit Log",
            href: setting_link + "audit-log",
        },
        {
            title: "Access Tokens",
            href: setting_link + "access-tokens",
        },
        {
            title: "Billing",
            href: setting_link + "billing",
        },
    ].map(link => ({
        ...link,
        variant: (pathname === link.href) ? "default" : "ghost",
    })));

    return (
        <Nav
            className="ml-10"
            isCollapsed={false}
            links={links}
        />
    );
}