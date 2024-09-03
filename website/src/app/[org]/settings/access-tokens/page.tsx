"use client"
import Link from "next/link";
import {FilterIcon} from "lucide-react";
import {Badge} from "~/components/ui/badge";
import {Dialog, DialogContent} from "~/components/ui/dialog";
import {useState} from "react";
import {Button} from "~/components/ui/button";


export default function HomePage() {

    const [isDialogOpen, setDialogOpen] = useState(false);

    const openDialog = () => setDialogOpen(true);
    const closeDialog = () => setDialogOpen(false);

    {/*TODO: Look to make this search better*/}
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
                            <FilterIcon className="mr-2" />
                            Filter
                            {/*<Badge className="ml-2">2</Badge>*/}
                        </Button>
                    </div>
                </div>
            </div>
            <div className="mt-4">
                <div className="overflow-x-auto">
                    <table className="min-w-full bg-white border border-gray-200">
                        <thead>
                        <tr className="text-left">
                            <th className="px-4 py-2 font-medium text-gray-600 border-b">Member</th>
                            <th className="px-4 py-2 font-medium text-gray-600 border-b">Action</th>
                            <th className="px-4 py-2 font-medium text-gray-600 border-b">Location</th>
                            <th className="px-4 py-2 font-medium text-gray-600 border-b">Date</th>
                        </tr>
                        </thead>
                        <tbody>
                        {/* Example row, replace with your data */}
                        <tr>
                            <td className="px-4 py-2 border-b">No results found</td>
                            <td className="px-4 py-2 border-b">We couldn't find any items matching your filter parameters.</td>
                            <td className="px-4 py-2 border-b"></td>
                            <td className="px-4 py-2 border-b"></td>
                        </tr>
                        </tbody>
                    </table>
                </div>
            </div>

            <Dialog open={isDialogOpen} onOpenChange={closeDialog}>
                <DialogContent>

                </DialogContent>
            </Dialog>

        </div>
    );
}
