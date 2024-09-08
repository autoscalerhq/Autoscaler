import {ReactNode} from "react";
import {CirclePlay, GitPullRequestDraft, MessageSquareText, Settings,} from "lucide-react";
import {Nav} from "~/components/navigation/vertical-nav";
import {ModelSwitcher} from "~/app/[org]/[env]/model-switcher";
import Header from "~/components/navigation/header";
import OrgNav from "~/app/[org]/[env]/org-nav";
import EnvNav from "~/app/[org]/[env]/env-nav";
import {Command} from "~/components/navigation/command-dialog";
import {CommandProvider} from "~/components/navigation/command-context";


export default function Layout({children, params}: { children: ReactNode, params: { org: string, env: string } }) {

    const base_link = `/${params.org}/${params.env}/`
    const org_link = `/${params.org}/`

    return (<div className="flex ">
            <CommandProvider>
                <nav className="w-64 bg-gray-100 py-4 px-2 h-screen overflow-auto flex flex-col justify-between">
                    <div>
                        <div>
                            {/*Project*/}
                            <ModelSwitcher items={[{
                                name: "Autoscaler", label: "Autoscaler", shortname: "prod"
                            },]} isCollapsed={false} ariaLabel={""} defaultItemName={"Autoscaler"}/>

                            <OrgNav org={params.org} env={params.env}/>
                        </div>


                        <div className="mt-4">
                            {/*Env*/}
                            <ModelSwitcher
                                defaultItemName={"Development"}
                                items={[{
                                    name: "Production", label: "Prod", icon: <GitPullRequestDraft/>, shortname: "prod"
                                }, {
                                    name: "Development", label: "Dev", shortname: "dev"
                                },]}
                                isCollapsed={false}
                                ariaLabel={""}/>

                            <ModelSwitcher
                                defaultItemName={"Api"}
                                items={[{
                                    name: "Api", label: "Api", shortname: "api"
                                }, {
                                    name: "Webhook", label: "Wh", shortname: "wh"
                                },]}
                                isCollapsed={false}
                                ariaLabel={""}
                                className={"mt-4"}
                            />

                            <EnvNav org={params.org} env={params.env}/>
                        </div>
                    </div>

                    <div className="mt-auto">
                        <Nav
                            isCollapsed={false}
                            links={[{
                                title: "Get started",
                                href: base_link + "get-started",
                                icon: CirclePlay,
                                variant: "ghost",
                            }, {
                                title: "Docs", href: "https://autoscaler.dev/docs", icon: Settings, variant: "ghost",
                            }, {
                                title: "Help", href: "#", icon: MessageSquareText, variant: "ghost",
                            },]}
                        />
                    </div>
                </nav>

                <div className="flex flex-col w-full">
                    <Header/>

                    <main className="flex-1 p-4 overflow-auto min-w-full ">
                        {children}
                    </main>
                </div>

                <Command/>

            </CommandProvider>
        </div>);
}