import {ReactNode} from "react";
import {Nav} from "~/components/navigation/vertical-nav";
import Header from "~/components/navigation/header";
import SettingsBack from "~/app/[org]/settings/settings-back";


export default function Layout({children, params}: { children: ReactNode, params: { org: string, env: string } }) {

    const setting_link = `/${params.org}/settings/`

    return (
        <div className="flex">
            <nav className="w-64 bg-gray-100 py-4 px-2 h-screen overflow-auto flex flex-col">
                <div className="flex flex-row justify-center">
                    <SettingsBack/>
                    <p className={" flex items-center text-2xl mr-auto"}> Settings</p>
                </div>

                <Nav
                    className="ml-10"
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
                            href: setting_link + "security",
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
                            title: "Service Tokens",
                            href: setting_link + "access-tokens",
                            variant: "ghost",
                        },
                    ]}
                />
            </nav>

            <div className="flex flex-col w-full">
                <Header/>

                <div className="flex-1 p-4 overflow-auto min-w-full ">
                    {children}
                </div>
            </div>

        </div>
    );
}