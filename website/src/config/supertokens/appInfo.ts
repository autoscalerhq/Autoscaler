import {AppInfo} from 'supertokens-node/lib/build/types';

export const appInfo = {
  appName: "AutoScaler",
  apiDomain: "http://localhost:4000",
  websiteDomain: "http://localhost:3000",
  apiBasePath: "/auth",
  websiteBasePath: "/auth",
} as const satisfies AppInfo;
