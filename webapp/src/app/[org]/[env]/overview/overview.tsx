import {Heading} from "lucide-react";
import {Button} from "~/components/ui/button";
import {Badge} from "~/components/ui/badge";
import {Popover, PopoverContent, PopoverTrigger} from "~/components/ui/popover";
import Link from "next/link";
import Flow from "~/app/[org]/[env]/overview/Steps";
import React from "react";

export default function Overview(){
    return (
        <div className="space-y-4">
            <div className="space-y-2">
                <Heading className="text-lg font-semibold">
                    Information
                </Heading>
                <div>
                    <Button variant="default">Manage workflow</Button>
                </div>
            </div>
            <div className="space-y-2">
                <div className="flex justify-between">
                    <p>Status</p>
                    <Button variant="default">Active</Button>
                </div>
                <div className="flex justify-between">
                    <p>Name</p>
                    <p>test</p>
                </div>
                <div className="flex justify-between">
                    <p>Description</p>
                    <p></p>
                </div>
                <div className="flex justify-between">
                    <p>Key</p>
                    <Badge >test</Badge>
                </div>
                <div className="flex justify-between">
                    <p>Categories</p>
                    <p>-</p>
                </div>
                <div className="flex justify-between">
                    <p>Created at</p>
                    <Popover>
                        <PopoverTrigger>
                            <span>Sep 4 2024, 12:21:43</span>
                        </PopoverTrigger>
                        <PopoverContent>
                            <p>Popover content</p>
                        </PopoverContent>
                    </Popover>
                </div>
                <div className="flex justify-between">
                    <p>Updated at</p>
                    <Popover>
                        <PopoverTrigger>
                            <span>Sep 7 2024, 13:23:06</span>
                        </PopoverTrigger>
                        <PopoverContent>
                            <p>Popover content</p>
                        </PopoverContent>
                    </Popover>
                </div>
            </div>
            <div className="space-y-2">
                <Heading className="text-lg font-semibold">
                    Steps
                </Heading>
                <div>
                    <Button variant="default">
                        <Link href="/bearbd/development/workflows/test/steps">
                            Edit steps
                        </Link>
                    </Button>
                </div>
                <Flow/>
            </div>
        </div>
    )
}