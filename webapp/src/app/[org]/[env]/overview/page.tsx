import React from "react";
import {Tabs, TabsContent, TabsList, TabsTrigger} from "~/components/ui/tabs";
import Overview from "~/app/[org]/[env]/overview/overview";


// Sample react component with Tailwind CSS
const YourComponent = () => {
    return (
        <Tabs defaultValue={"overview"}>
            <TabsList>
                <TabsTrigger value="overview">Overview</TabsTrigger>
            </TabsList>

            <TabsContent value={"overview"}>
              <Overview/>
            </TabsContent>
        </Tabs>
    );
};

export default YourComponent;
