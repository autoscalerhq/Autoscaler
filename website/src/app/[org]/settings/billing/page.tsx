import React from 'react';
import {Button} from "~/components/ui/button";
import Link from "next/link";

export default function BillingPanel() {
    return (
        <div className="p-4">
            <div className="container mx-auto p-4 space-y-4">

                <h1 className="text-2xl font-bold mb-4">Billing</h1>

                <div className="p-4 shadow-md border rounded-lg flex flex-row">
                    <div className="mb-4">
                        <p className="text-lg font-semibold">You are on the Developer plan</p>
                        <p className="text-sm text-gray-600">
                            The Developer plan is completely free, forever. Click “Upgrade” to learn about our paid
                            plans.
                        </p>
                    </div>
                    <div className="flex ml-auto mt-2 pr-2">
                        <Button className="shadow-md">
                            <Link href={"#"}>
                                Manage
                            </Link>
                        </Button>
                    </div>
                </div>

                <div className="p-6 shadow-md border rounded-lg">
                    <div className="mb-4">
                        <p className="text-lg font-semibold">Current month usage</p>
                        <p className="text-sm text-gray-600">
                            You have used <strong>0</strong> systems this month.
                            You have <strong>1 deployment </strong> with up
                            to <strong>5 services </strong> a month on your current plan.
                        </p>
                    </div>
                </div>

                <div className="p-6 shadow-md border rounded-lg ">
                    <div className="mb-4">
                        <p className="text-lg font-semibold">Billing contact email</p>
                        <div className={"flex flex-row"}>
                            <div className="mt-2 space-y-2">
                                <div className="flex justify-between">
                                    <p className="text-sm">Email</p>
                                    <p className="text-sm text-gray-600">zac16530@gmail.com</p>
                                </div>
                            </div>

                            <Button className="shadow-md ml-auto">
                                Update
                            </Button>
                        </div>
                    </div>
                </div>

                <div className="p-6 shadow-md border rounded-lg">
                    <p className="text-lg font-semibold">Payment method</p>
                    <div className="mb-4 flex flex-row">
                        <div className="mt-2 space-y-2">
                            <div className="flex justify-between">
                                <p className="text-sm">Card ending</p>
                            </div>
                        </div>

                        <Button className="shadow-md ml-auto">
                            Add
                        </Button>
                    </div>
                </div>
            </div>
        </div>
    );
}