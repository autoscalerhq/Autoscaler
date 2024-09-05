// app/org/[org]/[env]/page.js

import { redirect } from 'next/navigation';

interface OverviewPageProps {
    params: {
        org: string;
        env: string;
    };
}

export default async function EnvPage({ params }: OverviewPageProps) {
    const { org, env } = params;

    // Example: Check if the environment exists
    const envExists = await checkEnvExists(org, env);

    if (!envExists) {
        // Redirect to the not-found page if the environment does not exist
        redirect(`/org/${org}/${env}/not-found`);
    }

    return (
        <div>
            <h1>{env} Overview</h1>
            {/* Page content */}
        </div>
    );
}

// Example function to check if the environment exists
async function checkEnvExists(org: string, env: string) {
    // Replace this with your actual logic to check if the environment exists
    // const response = await fetch(`https://api.example.com/orgs/${org}/envs/${env}`);
    return 200;
}