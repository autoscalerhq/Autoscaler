import Link from "next/link";
import {Separator} from "~/components/ui/separator";
import {Copy} from "lucide-react";
import TimezonePicker from "~/components/ui/timezone-picker";


export default function HomePage() {
    return (
        <div className="p-6">
            <div className="space-y-6">
                <div>
                    <h2 className="font-bold text-2xl mb-2">Overview</h2>
                    <Separator/>
                </div>

                <div className="space-y-4">
                    <div className="flex ">
                        <div>
                            <p className="text-gray-400 w-40">Account name</p>
                        </div>
                        <div>
                            <p className="text-gray-400 ">bearbinary</p>
                        </div>
                    </div>
                    <div className="flex ">
                        <div>
                            <p className="text-gray-400 w-40">Account slug</p>
                        </div>
                        <div>
                            <p className="text-gray-400 ">bearbinary</p>
                        </div>
                    </div>
                    <div className="flex ">
                        <div>
                            <p className="text-gray-400 w-40">Account ID</p>
                        </div>
                        <div className="flex items-center bg-gray-200 w-auto rounded">
                            <p className=" text-gray-600 font-mono rounded px-2 text-sm">
                                6b03f8e5-8a03-44c7-9bed-b748ddaad124
                            </p>
                            <button className="ml-2 p-1 rounded bg-gray-300">
                                <Copy className={"w-4 h-4"}/>
                            </button>
                        </div>
                    </div>
                    <div className="flex ">
                        <div>
                            <p className="text-gray-400 w-40">Created at</p>
                        </div>
                        <div>
                            <p className="text-gray-400">Jul 26 2024, 14:49:18</p>
                        </div>
                    </div>
                </div>
            </div>

            <div className="mt-8 space-y-6">
                <div>
                    <h2 className="font-bold text-2xl mb-2">Feature settings</h2>
                    <Separator/>
                </div>

                <div className="space-y-4 border b-gray-600 rounded p-4">
                    <div className="space-y-2 flex flex-row">
                        <div>
                            <p className="text-gray-800 text-lg font-bold">Account timezone</p>
                            <p className="text-gray-500">
                                The default timezone for your account.{' '}
                                <a
                                    target="_blank"
                                    className="text-blue-500 underline"
                                    href="#"
                                >
                                    Learn more about timezone support.
                                </a>
                            </p>
                        </div>
                        <div className="flex ml-auto">
                            <TimezonePicker/>
                        </div>
                    </div>

                </div>
            </div>
        </div>
    );
}
