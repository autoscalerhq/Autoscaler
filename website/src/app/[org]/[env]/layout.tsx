import {Avatar, AvatarFallback, AvatarImage} from "~/components/ui/avatar";
import {Button} from "~/components/ui/button";
import {ReactNode} from "react";
import {
    Bell,
    BookDashed,
    BookUp, Boxes,
    ChevronsLeftRight,
    GitPullRequestDraft,
    LucideLayoutDashboard,
    PencilRuler, Search,
    Settings,
    SquareDashedKanban,
} from "lucide-react";
import {Nav} from "~/components/navigation/vertical-nav";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuTrigger
} from "~/components/ui/dropdown-menu";
import Feedback from "~/components/feedback/PageExperience";
import {ModelSwitcher} from "~/components/navigation/model-switcher";
import Header from "~/components/navigation/header";
import OrgNav from "~/app/[org]/[env]/org-nav";
import EnvNav from "~/app/[org]/[env]/env-nav";


export default function Layout({children, params}: { children: ReactNode, params: { org: string, env: string } }) {

    const base_link = `/${params.org}/${params.env}/`
    const org_link = `/${params.org}/`

    return (
        <div className="flex ">
            <nav className="w-64 bg-gray-100 py-4 px-2 h-screen overflow-auto flex flex-col justify-between">
                <div>
                    <div>
                        {/*Project*/}
                        <ModelSwitcher items={[
                            {
                                name: "Autoscaler",
                                label: "Autoscaler",
                                shortname: "prod"
                            },
                        ]} isCollapsed={false} ariaLabel={""} defaultItemName={"Autoscaler"}/>

                        <OrgNav org={params.org} env={params.env}/>
                    </div>


                    <div className="mt-4">
                        {/*Env*/}
                        <ModelSwitcher
                            defaultItemName={"Development"}
                            items={[
                                {
                                    name: "Production",
                                    label: "Prod",
                                    icon: <GitPullRequestDraft/>,
                                    shortname: "prod"
                                },
                                {
                                    name: "Development",
                                    label: "Dev",
                                    shortname: "dev"
                                },
                            ]}
                            isCollapsed={false}
                         ariaLabel={""}/>

                        <ModelSwitcher
                            defaultItemName={"Api"}
                            items={[
                                {
                                    name: "Api",
                                    label: "Api",
                                    shortname: "api"
                                },
                                {
                                    name: "Webhook",
                                    label: "Wh",
                                    shortname: "wh"
                                },
                            ]}
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
                        links={[
                            {
                                title: "Get started",
                                href: base_link + "get-started",
                                icon: Settings,
                                variant: "ghost",
                            },
                            {
                                title: "Docs",
                                href: "https://docs.autoscaler.dev",
                                icon: Settings,
                                variant: "ghost",
                            },
                        ]}
                    />
                </div>
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