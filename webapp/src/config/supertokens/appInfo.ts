import {AppInfo} from 'supertokens-node/lib/build/types';
import {GoBackendUrl} from '~/config';

export const appInfo = {
  appName: "AutoScaler",
  apiDomain: GoBackendUrl,
  websiteDomain: "http://localhost:3000",
  apiBasePath: "/auth",
  websiteBasePath: "/auth",
} as const satisfies AppInfo;
