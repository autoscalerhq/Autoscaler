import ThirdPartyReact, {Google, Facebook} from 'supertokens-auth-react/recipe/thirdparty'
import SessionReact from 'supertokens-auth-react/recipe/session'
import { appInfo } from './appinfo'
import { SuperTokensConfig } from 'supertokens-auth-react/lib/build/types'
import { useRouter } from "next/navigation";

const routerInfo: { router?: ReturnType<typeof useRouter>; pathName?: string } =
    {};

export function setRouter(
    router: ReturnType<typeof useRouter>,
    pathName: string,
) {
    routerInfo.router = router;
    routerInfo.pathName = pathName;
}

export const frontendConfig = (): SuperTokensConfig => {
    return {
        appInfo,
        recipeList: [
            ThirdPartyReact.init({
                signInAndUpFeature: {
                    providers: [
                        ThirdPartyReact.Google.init(),
                        ThirdPartyReact.Facebook.init(),
                        ThirdPartyReact.Apple.init(),
                        ThirdPartyReact.Github.init(),
                    ],
                },
            }),
            SessionReact.init(),
        ],
        windowHandler: (original) => ({
            ...original,
            location: {
                ...original.location,
                getPathName: () => routerInfo.pathName!,
                assign: (url) => routerInfo.router!.push(url.toString()),
                setHref: (url) => routerInfo.router!.push(url.toString()),
            },
        }),
    }
}