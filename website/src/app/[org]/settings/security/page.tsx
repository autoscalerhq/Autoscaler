import Link from "next/link";
import {Button} from "~/components/ui/button";
import {Label} from "~/components/ui/label";


export default function HomePage() {
    return (
        <div className="p-4">
            <div className=" pb-4">
                <div className="flex items-center justify-between">
                    <h2 className="text-2xl font-bold">Security</h2>
                    <div className="spacer"></div>
                </div>
                <div className="mt-4">
                    <div className="p-4 border rounded-lg">
                        <div className="flex items-center justify-between">
                            <div>
                                <p className="text-lg font-semibold">Enable auto-join for your domain</p>
                                <p className="text-sm text-gray-500">
                                    When enabled, any user with an email from the domains provided below can join your workspace as a member.
                                </p>
                            </div>
                            <div>
                                <span className="px-2 py-1 bg-gray-200 rounded-full">Disabled</span>
                            </div>
                        </div>
                        <div className="mt-4">
                            <div className="relative">
                                <input
                                    className="border-gray-300 rounded-md px-3 py-2 bg-gray-100 cursor-not-allowed"
                                    disabled
                                    placeholder="Add domains for autojoin"
                                />
                                <div className="absolute inset-y-0 right-0 flex items-center pr-2">
                                    <svg viewBox="0 0 24 24" focusable="false" className="w-5 h-5 text-gray-500">
                                        <path fill="currentColor" d="M16.59 8.59L12 13.17 7.41 8.59 6 10l6 6 6-6z"></path>
                                    </svg>
                                </div>
                            </div>
                        </div>
                        <p className="mt-2 text-sm text-gray-500">
                            You can select any non-public domains that belong to your account members.
                        </p>
                    </div>
                </div>

                <div className="mt-4">
                    <div className="flex flex-row p-4 border rounded-lg">
                        <div className="flex flex-col items-start justify-start">
                            <div className={"flex flex-row items-start justify-start space-x-4"}>
                                <p className="text-lg font-semibold ">Enable SAML SSO</p>
                                <Label className=" px-2 py-2 bg-green-200 rounded-full">Enterprise</Label>
                            </div>
                            <p className="mt-2 text-sm text-gray-500">
                                Anyone using email addresses with your domain can log in via SAML SSO.
                            </p>
                        </div>
                        <Button className={"ml-auto mt-2"}>
                            <Link href="/bearbinary/settings/billing">
                                Upgrade to Enterprise
                            </Link>
                        </Button>
                    </div>
                </div>

                <div className="mt-4">
                    <div className="flex flex-row p-4 border rounded-lg">
                        <div className="flex flex-col items-start justify-start">
                            <div className={"flex flex-row items-start justify-start space-x-4"}>
                                <p className="text-lg font-semibold ">Enable directory sync</p>
                                <Label className=" px-2 py-2 bg-green-200 rounded-full">Enterprise</Label>
                            </div>
                            <p className="mt-2 text-sm text-gray-500">
                            Provision and deprovision your organization&apos;s users using an Identity Management Platform (IdP).
                            </p>
                        </div>
                        <Button className={"ml-auto mt-2"} >
                            <Link href="/bearbinary/settings/billing">
                                Upgrade to Enterprise
                            </Link>
                        </Button>
                    </div>
                </div>
            </div>
        </div>
    );
}
