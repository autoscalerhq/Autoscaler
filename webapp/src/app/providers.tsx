'use client'
import posthog from 'posthog-js'
import { PostHogProvider } from 'posthog-js/react'
import {env} from "~/env";
import {ReactNode} from "react";

if (typeof window !== 'undefined' && env.NEXT_PUBLIC_POSTHOG_KEY ) {
    posthog.init(env.NEXT_PUBLIC_POSTHOG_KEY, {
        api_host: "/ingest",
        person_profiles: 'identified_only', // or 'always' to create profiles for anonymous users as well
    })
}
export function CSPostHogProvider({ children }: { children: ReactNode }) {
    return (
        <PostHogProvider client={posthog}>
            {children}
        </PostHogProvider>
    )
}



