import "~/styles/globals.css";

import { GeistSans } from "geist/font/sans";
import { type Metadata } from "next";

import React from 'react';

import {SuperTokensProvider} from "~/components/supertokensProvider";
import {CSPostHogProvider} from "~/app/providers";

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
      <SuperTokensProvider>
        <CSPostHogProvider>
          <body className={"overflow-hidden"}>
              {children}
          </body>
        </CSPostHogProvider>
      </SuperTokensProvider>
    </html>
  );
}
