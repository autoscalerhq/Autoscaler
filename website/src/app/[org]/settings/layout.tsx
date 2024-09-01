import {ReactNode} from "react";
import {BookDashed, Search, SquareDashedKanban,} from "lucide-react";
import {Nav} from "~/components/navigation/vertical-nav";
import Header from "~/components/navigation/header";
import {Label} from "~/components/ui/label";


export default function Layout({children, params}: { children: ReactNode, params: { org: string, env: string } }) {

    const setting_link = `/${params.org}/settings/`

    return (
        <div className="flex ">
            <nav className="w-64 bg-gray-100 py-4 px-2 h-screen overflow-auto flex flex-col justify-between">

                <Label > Settings</Label>
                <Nav
                    isCollapsed={false}
                    links={[
                        {
                            title: "General",
                            href: setting_link,
                            variant: "ghost",
                        },
                        {
                            title: "Environments",
                            href: setting_link + "environments",
                            variant: "ghost",
                        },
                        {
                            title: "Members",
                            href: setting_link + "members",
                            variant: "ghost",
                        },
                        {
                            title: "Security",
                            href: setting_link + "environments",
                            variant: "ghost",
                        },
                        {
                            title: "Billing",
                            href: setting_link + "billing",
                            variant: "ghost",
                        },
                        {
                            title: "Audit Log",
                            href: setting_link + "audit-log",
                            variant: "ghost",
                        },
                        {
                            title: "Security",
                            href: setting_link + "security",
                            variant: "ghost",
                        },
                        {
                            title: "Service Tokens",
                            href: setting_link + "service-tokens",
                            variant: "ghost",
                        },
                    ]}
                />
            </nav>

            <div className="flex flex-col w-full">
                <Header/>

                <main className="flex-1 p-4 overflow-auto min-w-full ">
                    {children}
                </main>
            </div>

        </div>
    );
}