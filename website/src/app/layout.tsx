import "~/styles/globals.css";

import { GeistSans } from "geist/font/sans";
import { type Metadata } from "next";

import React from 'react';

import SuperTokens, { SuperTokensWrapper } from "supertokens-auth-react";
import ThirdParty, {Github, Google, Facebook, Apple} from "supertokens-auth-react/recipe/thirdparty";
import Session from "supertokens-auth-react/recipe/session";
import {SuperTokensProvider} from "~/components/supertokensProvider";

// leverages client side context
// SuperTokens.init({
//   appInfo: {
//     appName: "autoscaler",
//     apiDomain: "http://localhost:8080",
//     websiteDomain: "http://localhost:3000",
//     apiBasePath: "/auth",
//     websiteBasePath: "/auth"
//   },
//   recipeList: [
//     ThirdParty.init({
//       signInAndUpFeature: {
//         providers: [
//           Github.init(),
//           Google.init(),
//           Facebook.init(),
//           Apple.init(),
//         ]
//       }
//     }),
//     Session.init()
//   ]
// });

export const metadata: Metadata = {
  title: "Autoscaler App",
  description: "Scale your app in any infrastructure",
  icons: [{ rel: "icon", url: "/favicon.ico" }],
};


export default function RootLayout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  return (
    <html lang="en" className={`${GeistSans.variable}`}>
      {/*<SuperTokensProvider>*/}
        <body>
            {children}
        </body>
      {/*</SuperTokensProvider>*/}
    </html>
  );
}
