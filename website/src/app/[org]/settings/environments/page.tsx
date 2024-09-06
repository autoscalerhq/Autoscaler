"use client"

import { useState } from "react";
import { Button } from "~/components/ui/button";
import { EllipsisVertical, Plus } from "lucide-react";
import { Table, TableBody, TableHead, TableHeader, TableRow, TableCell } from "~/components/ui/table";
import { Badge } from "~/components/ui/badge";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from "~/components/ui/dropdown-menu";
import CopyBox from "~/components/ui/copy-box";
import { Dialog, DialogTrigger, DialogContent, DialogHeader, DialogFooter } from "~/components/ui/dialog";
import {ColorPicker} from "~/components/ui/color-picker";

interface IEnvironments {
    name: string;
    status:  string;
    id:  string;
    color:  string;
}

export default function EnvironmentManager() {
    const environments: IEnvironments[] = [
        {
            name: "Development",
            status: "Default",
            id: "575f67ca-e67a-454e-bee7-3252a2e25fa1",
            color: "#FF0000",
        },
        {
            name: "Production",
            status: "",
            id: "4c36f976-a771-4942-b911-bd74e54b9e44",
            color: "#00FF00",
        },
    ];

    const [selectedEnvironment, setSelectedEnvironment] = useState<IEnvironments| null>(null);
    const [isCreateDialogOpen, setCreateDialogOpen] = useState(false);

    return (
        <div className="relative">
            <div className="flex flex-col">
                <div className="flex justify-between items-center">
                    <h2 className="text-lg font-bold">Environments</h2>
                    <Button variant={"ghost"} onClick={() => setCreateDialogOpen(true)}>
                        <Plus className="mr-2" />
                        Create environment
                    </Button>
                </div>

                <Table className="m-4">
                    <TableHeader>
                        <TableRow>
                            <TableHead>Name</TableHead>
                            <TableHead>Color</TableHead>
                            <TableHead>Status</TableHead>
                            <TableHead>Options</TableHead>
                            <TableHead/>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {environments.map((env) => (
                            <TableRow key={env.id}>
                                <TableCell>
                                    <div className={"flex flex-row items-center"}>
                                        <div className={`w-3 h-3 mr-2 rounded-sm`} style={{ backgroundColor: env.color }} />
                                        {env.name}
                                    </div>
                                </TableCell>
                                <TableCell>
                                    {env.status && <Badge variant="default">{env.status}</Badge>}
                                </TableCell>
                                <TableCell>
                                    <CopyBox className={"w-auto"} content={env.id} />
                                </TableCell>
                                <TableCell>
                                    <DropdownMenu>
                                        <DropdownMenuTrigger asChild>
                                            <Button variant={"ghost"}>
                                                <EllipsisVertical />
                                            </Button>
                                        </DropdownMenuTrigger>
                                        <DropdownMenuContent align="end">
                                            <DropdownMenuItem>Switch to default</DropdownMenuItem>
                                            <DropdownMenuItem
                                                onSelect={() => setSelectedEnvironment(env)}
                                            >
                                                Edit environment
                                            </DropdownMenuItem>
                                            <DropdownMenuSeparator />
                                            <DropdownMenuItem disabled>Delete environment</DropdownMenuItem>
                                        </DropdownMenuContent>
                                    </DropdownMenu>
                                </TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </div>

            <Dialog open={!!selectedEnvironment} onOpenChange={(isOpen) => {
                if (!isOpen) setSelectedEnvironment(null);
            }}>

                <DialogContent>
                    <DialogHeader>
                        <h2 className="text-lg font-bold">Update environment</h2>
                    </DialogHeader>
                    <form action="#" >
                        <div className="mb-4">
                            <label htmlFor="color" className="block text-sm font-medium text-gray-700">
                                Color
                            </label>
                            <input
                                type="text"
                                id="color"
                                name="color"
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                                defaultValue={selectedEnvironment?.color}
                            />
                            <ColorPicker value={selectedEnvironment?.color ?? "#000000"} onChange={ (value) =>{ console.log(value) }} />
                            <p className="mt-2 text-sm text-gray-500">
                                The color to assign this environment. Used in sidebar navigation.
                            </p>
                        </div>
                        <DialogFooter>
                            <Button type="button" variant="secondary" onClick={() => setSelectedEnvironment(null)}>
                                Cancel
                            </Button>
                            <Button type="submit" variant="default">
                                Update
                            </Button>
                        </DialogFooter>
                    </form>
                </DialogContent>
            </Dialog>

            <Dialog open={isCreateDialogOpen} onOpenChange={(isOpen) => setCreateDialogOpen(isOpen)}>
                <DialogContent>
                    <DialogHeader>
                        <h2 className="text-lg font-bold">Create environment</h2>
                        <p className="text-sm text-gray-600">Please note: this environment will always be placed before the production environment.</p>
                    </DialogHeader>
                    <form action="#">
                        <div className="mb-4">
                            <label htmlFor="new-name" className="block text-sm font-medium text-gray-700">Name</label>
                            <input
                                type="text"
                                id="new-name"
                                name="name"
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                                placeholder="Staging"
                            />
                            <ColorPicker value={selectedEnvironment?.color ?? "#000000"} onChange={  (value) =>{ console.log(value) }}/>
                            <p className="mt-2 text-sm text-gray-500">A unique name for this environment.</p>
                        </div>
                        <div className="mb-4">
                            <label htmlFor="new-color" className="block text-sm font-medium text-gray-700">Color</label>
                            <input
                                type="text"
                                id="new-color"
                                name="color"
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                                defaultValue="#14B8A7"
                            />
                            <p className="mt-2 text-sm text-gray-500">The color to assign this environment. Used in sidebar navigation.</p>
                        </div>

                        <DialogFooter>
                            <Button type="button" variant="secondary" onClick={() => setCreateDialogOpen(false)}>Cancel</Button>
                            <Button type="submit" variant="default">Create</Button>
                        </DialogFooter>
                    </form>
                </DialogContent>
            </Dialog>
        </div>
    );
}