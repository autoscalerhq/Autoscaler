"use client"
import {ReactNode, useState} from "react";
import { useRouter } from "next/navigation";
import { ChevronLeftIcon, ChevronRightIcon } from "lucide-react";
import { Button } from "~/components/ui/button";
import { Input } from "~/components/ui/input";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "~/components/ui/select";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "~/components/ui/table";
import { Badge } from "~/components/ui/badge";

interface IActions {
    id: string;
    name: string;
    type: string;
    service: string;
    createdAt: string;
    status: string;
    response?: string; // New property for responses
}

export default function Manager({params}: { params: { org: string, env: string } }) {

    const integrations: IActions[] = [
        {
            id: "0",
            name: "Increase network requests",
            type: "Scale up",
            service: "web",
            createdAt: "2023-01-01",
            status: "Active",
            response: "Scaling up process started"
        },
        {
            id: "1",
            name: "Increase ",
            type: "Scale up",
            service: "worker",
            createdAt: "2023-01-01",
            status: "Active",
            response: "Scaling up process started"
        },

    ];

    const [searchTerm, setSearchTerm] = useState("");
    const [statusFilter, setStatusFilter] = useState("All");
    const [currentPage, setCurrentPage] = useState(1);
    const itemsPerPage = 5;
    const router = useRouter();

    const filteredIntegrations = integrations.filter(
        (integration) =>
            (integration.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
                integration.createdAt.includes(searchTerm) ||
                integration.status.toLowerCase().includes(searchTerm.toLowerCase()) ||
                (integration.response && integration.response.toLowerCase().includes(searchTerm.toLowerCase()))) &&
            (statusFilter === "All" || integration.status === statusFilter)
    );

    const totalPages = Math.ceil(filteredIntegrations.length / itemsPerPage);
    const startIndex = (currentPage - 1) * itemsPerPage;
    const endIndex = startIndex + itemsPerPage;
    const currentIntegrations = filteredIntegrations.slice(startIndex, endIndex);

    const statusTypes = ["All", "Active", "Inactive"];

    return (
        <div className="p-4">
            <div className="container mx-auto p-4 space-y-4">
                <h1 className="text-2xl font-bold mb-4">Integrations</h1>
                <div className="flex flex-col sm:flex-row gap-4 mb-4">
                    <Input
                        placeholder="Search integrations..."
                        value={searchTerm}
                        onChange={(e) => setSearchTerm(e.target.value)}
                        className="w-full sm:w-64"
                    />
                    <Select value={statusFilter} onValueChange={setStatusFilter}>
                        <SelectTrigger className="w-full sm:w-40">
                            <SelectValue placeholder="Filter by status" />
                        </SelectTrigger>
                        <SelectContent>
                            {statusTypes.map((status) => (
                                <SelectItem key={status} value={status}>
                                    {status}
                                </SelectItem>
                            ))}
                        </SelectContent>
                    </Select>
                </div>
                <div className="rounded-md border">
                    <Table>
                        <TableHeader>
                            <TableRow>
                                <TableHead>Name</TableHead>
                                <TableHead>Created At</TableHead>
                                <TableHead>Response</TableHead>
                                <TableHead>Status</TableHead>
                                <TableHead>Options</TableHead>
                            </TableRow>
                        </TableHeader>
                        <TableBody>
                            {currentIntegrations.map((integration) => (
                                <TableRow key={integration.id}>
                                    <TableCell>{integration.name}</TableCell>
                                    <TableCell>{integration.createdAt}</TableCell>
                                    <TableCell>{integration.response}</TableCell>
                                    <TableCell>
                                        <Badge variant={integration.status === "Active" ? "default" : "secondary"}>{integration.status}</Badge>
                                    </TableCell>
                                    <TableCell>
                                        <Button
                                            variant="ghost"
                                            onClick={() =>
                                                router.push(`/${params.org}/${params.env}/actions/${integration.id}`)
                                            }
                                        >
                                            Manage
                                        </Button>
                                    </TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </div>
                <div className="flex justify-between items-center mt-4">
                    <div className="text-sm text-gray-500">
                        Showing {startIndex + 1} to {Math.min(endIndex, filteredIntegrations.length)} of {filteredIntegrations.length} entries
                    </div>
                    <div className="flex gap-2">
                        <Button
                            variant="outline"
                            size="icon"
                            onClick={() => setCurrentPage((prev) => Math.max(prev - 1, 1))}
                            disabled={currentPage === 1}
                        >
                            <ChevronLeftIcon className="h-4 w-4" />
                        </Button>
                        <Button
                            variant="outline"
                            size="icon"
                            onClick={() => setCurrentPage((prev) => Math.min(prev + 1, totalPages))}
                            disabled={currentPage === totalPages}
                        >
                            <ChevronRightIcon className="h-4 w-4" />
                        </Button>
                    </div>
                </div>
            </div>
        </div>
    );
}