import { ReactNode } from "react";
import Header from "~/components/navigation/header";
import SettingsBack from "~/app/[org]/settings/settings-back";
import SettingsNav from "~/app/[org]/settings/settings-nav";
import {CommandProvider} from "~/components/navigation/command-context";
import {Command} from "~/components/navigation/command-dialog";

export default function Layout({ children, params }: { children: ReactNode, params: { org: string, env: string } }) {

    return (
        <CommandProvider org={params.org} env={params.env}>
            <div className="flex">
                <nav className="w-64 bg-gray-100 py-4 px-2 h-screen overflow-auto flex flex-col">
                    <div className="flex flex-row justify-center">
                        <SettingsBack org={params.org}  />
                        <p className="flex items-center text-2xl mr-auto">Settings</p>
                    </div>

                   <SettingsNav org={params.org}/>
                </nav>

                <div className="flex flex-col w-full">
                    <Header />

                    <div className="flex-1 p-4 overflow-auto min-w-full">
                        {children}
                    </div>
                </div>
            </div>
            <Command/>
        </CommandProvider>
    );
}