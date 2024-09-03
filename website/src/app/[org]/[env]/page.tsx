// src/app/org/[env]/page.tsx

import { redirect } from 'next/navigation';

interface Params {
    env: string;
}

const OrgEnvPage = ({ params }: { params: Params }) => {
    const { env } = params;

    // Redirect to /org/env/overview
    redirect(`/org/${env}/overview`);

    return null;
};

export default OrgEnvPage;