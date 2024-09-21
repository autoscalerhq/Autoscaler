"use client"
import { ReactNode, useState } from "react";
import { useRouter } from "next/navigation";
import { ChevronLeftIcon, ChevronRightIcon } from "lucide-react";
import { Button } from "~/components/ui/button";
import { Input } from "~/components/ui/input";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "~/components/ui/select";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "~/components/ui/table";
import { Badge } from "~/components/ui/badge";

interface IScaler {
    id: string;
    displayName: string;
    integrationName: string;
    serviceName: string;
    options: ReactNode;
}

export default function ScalerManager({ params }: { params: { org: string, env: string } }) {
    const scalers: IScaler[] = [
        { id: "1", displayName: "ActiveMQ", integrationName: "activemq", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "2", displayName: "ActiveMQ Artemis", integrationName: "activemq", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "3", displayName: "Apache Kafka", integrationName: "apache", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "4", displayName: "Apache Kafka (Experimental)", integrationName: "apache", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "5", displayName: "Apache Pulsar", integrationName: "apache", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "6", displayName: "ArangoDB", integrationName: "arangodb", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "7", displayName: "AWS CloudWatch", integrationName: "aws", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "8", displayName: "AWS DynamoDB", integrationName: "aws", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "9", displayName: "AWS DynamoDB Streams", integrationName: "aws", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "10", displayName: "AWS Kinesis Stream", integrationName: "aws", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "11", displayName: "AWS SQS Queue", integrationName: "aws", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "12", displayName: "Azure Application Insights", integrationName: "azure", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "13", displayName: "Azure Blob Storage", integrationName: "azure", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "14", displayName: "Azure Data Explorer", integrationName: "azure", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "15", displayName: "Azure Event Hubs", integrationName: "azure", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "16", displayName: "Azure Log Analytics", integrationName: "azure", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "17", displayName: "Azure Monitor", integrationName: "azure", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "18", displayName: "Azure Pipelines", integrationName: "azure", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "19", displayName: "Azure Service Bus", integrationName: "azure", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "20", displayName: "Azure Storage Queue", integrationName: "azure", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "21", displayName: "Cassandra", integrationName: "apache", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "22", displayName: "CouchDB", integrationName: "apache", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "23", displayName: "CPU", integrationName: "system", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "24", displayName: "Cron", integrationName: "system", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "25", displayName: "Datadog", integrationName: "datadog", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "26", displayName: "Dynatrace", integrationName: "dynatrace", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "27", displayName: "Elasticsearch", integrationName: "elastic", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "28", displayName: "Etcd", integrationName: "coreos", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "29", displayName: "External", integrationName: "external", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "30", displayName: "External Push", integrationName: "external", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "31", displayName: "Github Runner Scaler", integrationName: "github", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "32", displayName: "Google Cloud Platform Cloud Tasks", integrationName: "gcp", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "33", displayName: "Google Cloud Platform Pub/Sub", integrationName: "gcp", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "34", displayName: "Google Cloud Platform Stackdriver", integrationName: "gcp", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "35", displayName: "Google Cloud Platform Storage", integrationName: "gcp", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "36", displayName: "Graphite", integrationName: "graphite", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "37", displayName: "Huawei Cloudeye", integrationName: "huawei", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "38", displayName: "IBM MQ", integrationName: "ibm", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "39", displayName: "InfluxDB", integrationName: "influxdb", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "40", displayName: "Kubernetes Workload", integrationName: "kubernetes", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "41", displayName: "Liiklus Topic", integrationName: "liiklus", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "42", displayName: "Loki", integrationName: "grafana", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "43", displayName: "Memory", integrationName: "system", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "44", displayName: "Metrics API", integrationName: "system", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "45", displayName: "MongoDB", integrationName: "mongodb", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "46", displayName: "MSSQL", integrationName: "mssql", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "47", displayName: "MySQL", integrationName: "mysql", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "48", displayName: "NATS JetStream", integrationName: "nats", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "49", displayName: "NATS Streaming", integrationName: "nats", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "50", displayName: "New Relic", integrationName: "newrelic", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "51", displayName: "OpenStack Metric", integrationName: "openstack", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "52", displayName: "OpenStack Swift", integrationName: "openstack", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "53", displayName: "PostgreSQL", integrationName: "postgresql", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "54", displayName: "Predictkube", integrationName: "predictkube", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "55", displayName: "Prometheus", integrationName: "prometheus", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "56", displayName: "RabbitMQ Queue", integrationName: "rabbitmq", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "57", displayName: "Redis Lists", integrationName: "redis", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "58", displayName: "Redis Lists (supports Redis Cluster)", integrationName: "redis", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "59", displayName: "Redis Lists (supports Redis Sentinel)", integrationName: "redis", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "60", displayName: "Redis Streams", integrationName: "redis", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "61", displayName: "Redis Streams (supports Redis Cluster)", integrationName: "redis", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "62", displayName: "Redis Streams (supports Redis Sentinel)", integrationName: "redis", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "63", displayName: "Selenium Grid Scaler", integrationName: "selenium", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "64", displayName: "Solace PubSub+ Event Broker", integrationName: "solace", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "65", displayName: "Solr", integrationName: "apache", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> },
        { id: "66", displayName: "Splunk", integrationName: "splunk", serviceName: "cfn-service", options: <Button variant="ghost">Manage</Button> }
    ];

    const [searchTerm, setSearchTerm] = useState("");
    const [currentPage, setCurrentPage] = useState(1);
    const itemsPerPage = 5;
    const router = useRouter();

    const filteredScalers = scalers.filter(
        (scaler) =>
            scaler.displayName.toLowerCase().includes(searchTerm.toLowerCase()) ||
            scaler.integrationName.toLowerCase().includes(searchTerm.toLowerCase()) ||
            scaler.serviceName.toLowerCase().includes(searchTerm.toLowerCase())
    );

    const totalPages = Math.ceil(filteredScalers.length / itemsPerPage);
    const startIndex = (currentPage - 1) * itemsPerPage;
    const endIndex = startIndex + itemsPerPage;
    const currentScalers = filteredScalers.slice(startIndex, endIndex);

    return (
        <div className="p-4">
            <div className="container mx-auto p-4 space-y-4">
                <h1 className="text-2xl font-bold mb-4">Scalers</h1>
                <div className="flex flex-col sm:flex-row gap-4 mb-4">
                    <Input
                        placeholder="Search scalers..."
                        value={searchTerm}
                        onChange={(e) => setSearchTerm(e.target.value)}
                        className="w-full sm:w-64"
                    />
                </div>
                <div className="rounded-md border">
                    <Table>
                        <TableHeader>
                            <TableRow>
                                <TableHead>Display Name</TableHead>
                                <TableHead>Integration Name</TableHead>
                                <TableHead>Service Name</TableHead>
                                <TableHead>Options</TableHead>
                            </TableRow>
                        </TableHeader>
                        <TableBody>
                            {currentScalers.map((scaler) => (
                                <TableRow key={scaler.id}>
                                    <TableCell>{scaler.displayName}</TableCell>
                                    <TableCell>{scaler.integrationName}</TableCell>
                                    <TableCell>{scaler.serviceName}</TableCell>
                                    <TableCell>
                                        <Button
                                            variant="ghost"
                                            onClick={() =>
                                                router.push(`/${params.org}/${params.env}/scalers/${scaler.id}`)
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
                        Showing {startIndex + 1} to {Math.min(endIndex, filteredScalers.length)} of {filteredScalers.length} entries
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