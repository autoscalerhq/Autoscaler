"use client"
import Link from "next/link";
import {EllipsisVertical, FilterIcon, Plus} from "lucide-react";
import {Badge} from "~/components/ui/badge";
import {Dialog, DialogContent} from "~/components/ui/dialog";
import {useState} from "react";
import {Button} from "~/components/ui/button";
import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "~/components/ui/table";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuSeparator,
    DropdownMenuTrigger
} from "~/components/ui/dropdown-menu";

interface IAccessToken {
    name:string;
    lastUsed: string;
    createdOn: string;
}

export default function HomePage() {

    const [isDialogOpen, setDialogOpen] = useState(false);

    const openDialog = () => setDialogOpen(true);
    const closeDialog = () => setDialogOpen(false);

    const acessToken: IAccessToken[] = [
        {
            name: "Test",
            lastUsed: "2024-04-05T12:30−02:00",
            createdOn: "2007-04-05T12:30−02:00",
        }
    ]

    return (
        <div className="p-4">
            <div className="pb-4 border-b">
                <div className="flex items-center justify-between">
                    <h2 className="text-2xl font-bold">Access Tokens</h2>
                    <div className="relative">
                        <Button
                            onClick={openDialog}
                            type="button"
                            className="flex items-center px-4 py-2 text-white bg-blue-500 rounded-md">
                            <Plus className="mr-2" />
                            New Token
                        </Button>
                    </div>
                </div>
            </div>
            <div className="mt-4">
                <div className="overflow-x-auto">
                    <Table className="min-w-full bg-white border border-gray-200">
                        <TableHeader>
                            <TableRow>
                                <TableHead className="px-4 py-2 font-medium text-gray-600 border-b">Name</TableHead>
                                <TableHead className="px-4 py-2 font-medium text-gray-600 border-b">Last Used</TableHead>
                                <TableHead className="px-4 py-2 font-medium text-gray-600 border-b">Created</TableHead>
                                <TableHead/>
                            </TableRow>
                        </TableHeader>
                        <TableBody>
                            {acessToken.map((token) => (
                                <TableRow key={token.name}>
                                    <TableCell>
                                        <div className={"flex flex-row items-center"}>
                                            {token.name}
                                        </div>
                                    </TableCell>
                                    <TableCell>
                                        {token.lastUsed}
                                    </TableCell>
                                    <TableCell>
                                        {token.createdOn}
                                    </TableCell>
                                    <TableCell>
                                        <DropdownMenu>
                                            <DropdownMenuTrigger asChild>
                                                <Button variant={"ghost"}>
                                                    <EllipsisVertical />
                                                </Button>
                                            </DropdownMenuTrigger>
                                            <DropdownMenuContent align="end">
                                                <DropdownMenuItem>Edit Name</DropdownMenuItem>
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
            </div>

            <Dialog open={isDialogOpen} onOpenChange={closeDialog}>
                <DialogContent>

                </DialogContent>
            </Dialog>

        </div>
    );
}
