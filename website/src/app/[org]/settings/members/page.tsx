"use client"

import { useState } from "react";
import { Button } from "~/components/ui/button";
import {EllipsisVertical, Plus, Send} from "lucide-react";
import { Table, TableBody, TableHead, TableHeader, TableRow, TableCell } from "~/components/ui/table";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from "~/components/ui/dropdown-menu";
import { Dialog, DialogContent, DialogHeader, DialogFooter } from "~/components/ui/dialog";
import {Label} from "~/components/ui/label";
import {Input} from "~/components/ui/input";
import {Select, SelectContent, SelectGroup, SelectItem} from "~/components/ui/select";
import {Textarea} from "~/components/ui/textarea";

interface IMembers {
    name: string;
    email:  string;
    role:  string;
    addedAt:  string;
}

interface IPendingMembers {
    email:  string;
    role:  string;
    createdAt:  string;
}

export default function membersManager() {
    const members: IMembers[] = [
        {
            email: "zac16530@gmail.com",
            name: "Zac Clifton",
            role: "Owener",
            addedAt: "2021-12-17T03:24:00",
        },
        {
            email: "astiner21@gmail.com",
            name: "Alex Clifton",
            role: "Admin",
            addedAt: "2023-12-17T03:24:00",
        },
    ];

    const pendingInvites: IPendingMembers[] = [
        {
            email: "jason@gmail.com",
            role: "Admin",
            createdAt: "2023-12-17T03:24:00",
        },
    ]

    const [isInviteDialogOpen, setInviteDialogOpen] = useState(false);
    const [isPendingInvitesDialogOpen, setPendingInvitesDialogOpen] = useState(false);

    return (
        <div className="relative">
            <div className="flex flex-col">
                <div className="flex justify-between items-center">
                    <h2 className="text-lg font-bold">Environments</h2>
                    <div>
                        <Button variant={"ghost"} onClick={() => setPendingInvitesDialogOpen(true)}>
                            Pending Invites
                        </Button>
                        <Button variant={"ghost"} onClick={() => setInviteDialogOpen(true)}>
                            <Send className="mr-2" />
                            Invite Member
                        </Button>
                    </div>
                </div>

                <Table className="m-4">
                    <TableHeader>
                        <TableRow>
                            <TableHead>Email</TableHead>
                            <TableHead>Name</TableHead>
                            <TableHead>Role</TableHead>
                            <TableHead>Added At</TableHead>
                            <TableHead/>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {members.map((member) => (
                            <TableRow key={member.email}>
                                <TableCell>
                                    <div className={"flex flex-row items-center"}>
                                        {member.email}
                                    </div>
                                </TableCell>
                                <TableCell>
                                    {member.name}
                                </TableCell>
                                <TableCell>
                                    {member.role}
                                </TableCell>
                                <TableCell>
                                    {member.addedAt}
                                </TableCell>
                                <TableCell>
                                    <DropdownMenu>
                                        <DropdownMenuTrigger asChild>
                                            <Button variant={"ghost"}>
                                                <EllipsisVertical />
                                            </Button>
                                        </DropdownMenuTrigger>
                                        <DropdownMenuContent align="end">
                                            <DropdownMenuItem>Change Role</DropdownMenuItem>
                                            <DropdownMenuSeparator />
                                            <DropdownMenuItem disabled>Delete Member</DropdownMenuItem>
                                        </DropdownMenuContent>
                                    </DropdownMenu>
                                </TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </div>

            <Dialog open={isInviteDialogOpen} onOpenChange={(isOpen) => setInviteDialogOpen(isOpen)}>
                <DialogContent>
                    <DialogHeader>
                        <h2 className="text-lg font-bold">Invite member</h2>
                        <p className="text-sm text-gray-600">Invite a new member into your account</p>
                    </DialogHeader>
                    <form action="#" autoComplete="off">
                        <div className="mb-4">
                            <Label htmlFor="email" className="block text-sm font-medium text-gray-700">Email address</Label>
                            <Input
                                type="email"
                                id="email"
                                name="email"
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                                placeholder="jhammond@ingen.net"
                                defaultValue="astiner21@gmail.com"
                            />
                        </div>
                        <div className="mb-4">
                            <Label htmlFor="role" className="block text-sm font-medium text-gray-700">Role</Label>
                            <Select
                                name="role">
                                <SelectContent>
                                    <SelectGroup>
                                        <SelectItem value="OWNER">Owner</SelectItem>
                                        <SelectItem value="ADMIN">Admin</SelectItem>
                                        <SelectItem value="MEMBER">Member</SelectItem>
                                        <SelectItem value="SUPPORT">Support</SelectItem>
                                        <SelectItem value="BILLING">Billing</SelectItem>
                                    </SelectGroup>
                                </SelectContent>
                            </Select>
                            <p className="text-sm text-gray-500 mt-2">
                                <a target="_blank" rel="noopener noreferrer" className="text-indigo-600 underline" href="https://docs.knock.app/manage-your-account/roles-and-permissions?_gl=1*2y2vwj*_gcl_aw*R0NMLjE3MjIyMTE1NTMuRUFJYUlRb2JDaE1JenZDVV92bktod01WT2tuX0FSM3MtaTkyRUFBWUFTQUFFZ0tTbFBEX0J3RQ..*_gcl_au*NjEzMzg3NzIxLjE3MjIyMTE1NTM.*_ga*MTYyNjEzMDM5MS4xNzIyMjExNTUz*_ga_GJR2JW7XCV*MTcyNTMxMjkyMi4xNC4wLjE3MjUzMTI5MjIuNjAuMC4w">
                                    Learn more
                                </a>{" "}about roles within Knock.
                            </p>
                        </div>
                        <div className="mb-4">
                            <Label htmlFor="message" className="block text-sm font-medium text-gray-700">
                                Add a message <span className="text-gray-500">(Optional)</span>
                            </Label>
                            <Textarea
                                id="message"
                                name="message"
                                rows={3}
                                placeholder="Write a note to the invited recipient..."
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">
                            </Textarea>
                        </div>
                        <div className="mb-4">
                            <p className="text-sm text-gray-600">&#x1F31F; <strong>Tip:</strong> Enable auto-join for your domain members</p>
                            <p className="text-sm text-gray-500">
                                You can enable auto-join by domain under the <a href="/bearbinary/settings" className="text-indigo-600 underline">General
                                tab</a> of account settings. This helps your team members find your account instead of creating their own.
                            </p>
                        </div>
                        <DialogFooter>
                            <Button type="button" variant="secondary" onClick={() => setInviteDialogOpen(false)}>Cancel</Button>
                            <Button type="submit" variant="default">Invite</Button>
                        </DialogFooter>
                    </form>
                </DialogContent>
            </Dialog>


            <Dialog open={isPendingInvitesDialogOpen} onOpenChange={(isOpen) => setPendingInvitesDialogOpen(isOpen)}>
                <DialogContent>
                    <DialogHeader>
                        <h2 className="text-lg font-bold">Pending Invites</h2>
                    </DialogHeader>
                    <div className="overflow-auto">
                        <Table className="m-4">
                            <TableHeader>
                                <TableRow>
                                    <TableHead>Email</TableHead>
                                    <TableHead>Role</TableHead>
                                    <TableHead>Added At</TableHead>
                                    <TableHead/>
                                </TableRow>
                            </TableHeader>
                            <TableBody>
                                {pendingInvites.map((invite) => (
                                    <TableRow key={invite.email}>
                                        <TableCell>
                                            <div className="flex flex-row items-center">
                                                {invite.email}
                                            </div>
                                        </TableCell>
                                        <TableCell>
                                            {invite.role}
                                        </TableCell>
                                        <TableCell>
                                            {invite.createdAt}
                                        </TableCell>
                                        <TableCell>
                                            <DropdownMenu>
                                                <DropdownMenuTrigger asChild>
                                                    <Button variant={"ghost"}>
                                                        <EllipsisVertical />
                                                    </Button>
                                                </DropdownMenuTrigger>
                                                <DropdownMenuContent align="end">
                                                    <DropdownMenuItem>Resend</DropdownMenuItem>
                                                    <DropdownMenuSeparator />
                                                    <DropdownMenuItem disabled>Cancel</DropdownMenuItem>
                                                </DropdownMenuContent>
                                            </DropdownMenu>
                                        </TableCell>
                                    </TableRow>
                                ))}
                            </TableBody>
                        </Table>
                    </div>
                    <DialogFooter>
                        <Button type="button" variant="secondary" onClick={() => setPendingInvitesDialogOpen(false)}>
                            Close
                        </Button>
                    </DialogFooter>
                </DialogContent>
            </Dialog>

        </div>
    );
}