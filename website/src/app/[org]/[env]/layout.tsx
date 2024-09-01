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
    PencilRuler,
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


export default function Layout({children, params}: { children: ReactNode, params: { org: string, env: string } }) {

    const base_link = `/${params.org}/${params.env}/`

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

                        {/*<Separator title={"test"} className={"bg-gray-700 mt-1.5"}/>*/}

                        <Nav
                            isCollapsed={false}
                            links={[
                                {
                                    title: "Settings",
                                    href: base_link + "settings",
                                    icon: Settings,
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
                                    icon: SquareDashedKanban,
                                    variant: "ghost",
                                },
                            ]}
                        />
                    </div>

                    {/*<Separator title={"test"} className={"bg-gray-700"}/>*/}

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

                        <Nav
                            isCollapsed={false}
                            links={[
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
                            ]}
                        />
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
                <header
                    className="sticky top-0 z-30 flex h-16 items-center gap-4 border-b bg-background px-4 border-gray-200">

                    <Feedback className={"ml-auto"}/>

                    <DropdownMenu>
                        <DropdownMenuTrigger asChild>
                            <Button
                                variant="outline"
                                size="icon"
                                className="overflow-hidden rounded-full"
                            >
                                <Avatar>
                                    <AvatarImage src="https://github.com/shadcn.png"/>
                                    <AvatarFallback>CN</AvatarFallback>
                                </Avatar>
                            </Button>
                        </DropdownMenuTrigger>
                        <DropdownMenuContent align="end">
                            <DropdownMenuLabel>My Account</DropdownMenuLabel>
                            <DropdownMenuSeparator/>
                            <DropdownMenuItem>Settings</DropdownMenuItem>
                            <DropdownMenuItem>Support</DropdownMenuItem>
                            <DropdownMenuSeparator/>
                            <DropdownMenuItem>Logout</DropdownMenuItem>
                        </DropdownMenuContent>
                    </DropdownMenu>
                </header>

                <main className="flex-1 p-4 overflow-auto min-w-full ">
                    {children}
                </main>
            </div>

        </div>
    );
}