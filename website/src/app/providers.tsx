'use client'
import posthog from 'posthog-js'
import { PostHogProvider } from 'posthog-js/react'
import {env} from "~/env";
import {ReactNode} from "react";
import SuperTokens from "supertokens-auth-react";
import ThirdParty, {Apple, Facebook, Github, Google} from "supertokens-auth-react/recipe/thirdparty";
import Session from "supertokens-auth-react/recipe/session";

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

if (typeof window !== 'undefined'  ) {
    SuperTokens.init({
        appInfo: {
            appName: "autoscaler",
            apiDomain: "http://localhost:8080",
            websiteDomain: "http://localhost:3000",
            apiBasePath: "/auth",
            websiteBasePath: "/auth"
        },
        recipeList: [
            ThirdParty.init({
                signInAndUpFeature: {
                    providers: [
                        Github.init(),
                        Google.init(),
                        Facebook.init(),
                        Apple.init(),
                    ]
                }
            }),
            Session.init()
        ]
    });

}


