// app/layout.js

import Link from 'next/link';
import {Avatar, AvatarFallback, AvatarImage} from "~/components/ui/avatar";
import {Button} from "~/components/ui/button";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel, DropdownMenuSeparator,
    DropdownMenuTrigger
} from "~/components/ui/dropdown-menu";
import {ReactNode} from "react";


export default function Layout({ children, params }: { children: ReactNode, params: {org: string, env: string}}) {

    const base_link = `/${params.org}/${params.env}/`

    return (
        <div className="flex " >
            <nav className="w-64 bg-gray-100 p-4 h-screen overflow-auto">
                <div className="flex flex-col justify-between h-full">
                    <div>
                        <div>
                            <DropdownMenu>
                                <DropdownMenuTrigger asChild>
                                    <Button
                                        className="flex items-center justify-between w-full"
                                        aria-haspopup="menu"
                                    >
                              <span className="flex items-center">
                                <Avatar>
                                  <AvatarImage src="https://github.com/shadcn.png"/>
                                  <AvatarFallback>CN</AvatarFallback>
                                </Avatar>
                                bearbinary
                              </span>
                                        <svg
                                            xmlns="http://www.w3.org/2000/svg"
                                            width="24"
                                            height="24"
                                            viewBox="0 0 24 24"
                                            fill="none"
                                            stroke="currentColor"
                                            strokeWidth="2"
                                            strokeLinecap="round"
                                            strokeLinejoin="round"
                                        >
                                            <path d="m6 9 6 6 6-6"></path>
                                        </svg>
                                    </Button>
                                </DropdownMenuTrigger>
                                <DropdownMenuContent>
                                    <DropdownMenuLabel>Accounts</DropdownMenuLabel>
                                    <DropdownMenuItem>bearbinary</DropdownMenuItem>
                                    <DropdownMenuSeparator/>
                                    <Link href="/join-account">
                                        <DropdownMenuItem asChild>
                                            <Button className="w-full justify-between">
                                                <span>Create or join account</span>
                                                <svg
                                                    xmlns="http://www.w3.org/2000/svg"
                                                    width="16"
                                                    height="16"
                                                    viewBox="0 0 24 24"
                                                    fill="none"
                                                    stroke="currentColor"
                                                    strokeWidth="2"
                                                    strokeLinecap="round"
                                                    strokeLinejoin="round"
                                                >
                                                    <path d="M256 112v288m144-144H112"></path>
                                                </svg>
                                            </Button>
                                        </DropdownMenuItem>
                                    </Link>
                                </DropdownMenuContent>
                            </DropdownMenu>
                            <Link href={base_link + "settings"} passHref>
                                <p>Settings</p>
                            </Link>
                            <Link href={base_link + "system-templates"} passHref>
                                <p>System Templates</p>
                            </Link>
                            <Link href={base_link + "service-templates"} passHref>
                                <p>Service Templates</p>
                            </Link>
                        </div>
                        <div className="mt-4 flex-grow">
                            <DropdownMenu>
                                <DropdownMenuTrigger asChild>
                                    <Button
                                        className="flex items-center justify-between w-full"
                                        aria-haspopup="menu"
                                    >
                              <span className="flex items-center">
                                <Avatar>
                                  <AvatarImage src="https://github.com/shadcn.png"/>
                                  <AvatarFallback>CN</AvatarFallback>
                                </Avatar>
                                bearbinary
                              </span>
                                        <svg
                                            xmlns="http://www.w3.org/2000/svg"
                                            width="24"
                                            height="24"
                                            viewBox="0 0 24 24"
                                            fill="none"
                                            stroke="currentColor"
                                            strokeWidth="2"
                                            strokeLinecap="round"
                                            strokeLinejoin="round"
                                        >
                                            <path d="m6 9 6 6 6-6"></path>
                                        </svg>
                                    </Button>
                                </DropdownMenuTrigger>
                                <DropdownMenuContent>
                                    <DropdownMenuLabel>Accounts</DropdownMenuLabel>
                                    <DropdownMenuItem>bearbinary</DropdownMenuItem>
                                    <DropdownMenuSeparator/>
                                    <Link href="/join-account">
                                        <DropdownMenuItem asChild>
                                            <Button className="w-full justify-between">
                                                <span>Create or join account</span>
                                                <svg
                                                    xmlns="http://www.w3.org/2000/svg"
                                                    width="16"
                                                    height="16"
                                                    viewBox="0 0 24 24"
                                                    fill="none"
                                                    stroke="currentColor"
                                                    strokeWidth="2"
                                                    strokeLinecap="round"
                                                    strokeLinejoin="round"
                                                >
                                                    <path d="M256 112v288m144-144H112"></path>
                                                </svg>
                                            </Button>
                                        </DropdownMenuItem>
                                    </Link>
                                </DropdownMenuContent>
                            </DropdownMenu>

                            <Link href={base_link + "pull"} passHref>
                                <p>Pull</p>
                            </Link>
                            <Link href={base_link + "stream"} passHref>
                                <p>Stream</p>
                            </Link>
                            <Link href={base_link + "push"} passHref>
                                <p>Push</p>
                            </Link>
                            <Link href={base_link + "policies"} passHref>
                                <p>Polices</p>
                            </Link>
                            <Link href={base_link + "scalers"} passHref>
                                <p>Scalers</p>
                            </Link>
                            <Link href={base_link + "analytics"} passHref>
                                <p>Analytics</p>
                            </Link>
                            <Link href={base_link + "alerts"} passHref>
                                <p>Alerts</p>
                            </Link>
                        </div>
                    </div>

                    <div className={"h-full"}>
                        <Link href={base_link + "get-started"} passHref>
                            <p>Get started</p>
                        </Link>
                        <Link href="https://docs.autoscaler.dev" passHref>
                            <p>Docs</p>
                        </Link>
                    </div>
                </div>
            </nav>
            <div className="flex flex-col w-full">
                <div className="w-full flex space-x-10 items-center justify-end pt-2 border-b border-gray-400">
                    <Button className="btn-primary">
                        Feedback?
                    </Button>

                    <Avatar>
                        <AvatarImage src="https://github.com/shadcn.png"/>
                        <AvatarFallback>CN</AvatarFallback>
                    </Avatar>
                </div>

                <main className="flex-1 p-4 overflow-auto min-w-full ">
                    {children}
                </main>
            </div>

        </div>
    );
}