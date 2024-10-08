"use client";
import React from "react";
import { SuperTokensWrapper } from "supertokens-auth-react";
import {ensureFrontendSuperTokensInit, setRouter} from '~/config/supertokens/frontend';
import { usePathname, useRouter } from "next/navigation";

if (typeof window !== 'undefined') {
  // we only want to call this init function on the frontend, so we check typeof window !== 'undefined'
  ensureFrontendSuperTokensInit();
}

export const SuperTokensProvider: React.FC<React.PropsWithChildren<NonNullable<unknown>>> = ({
  children,
}) => {
  setRouter(useRouter(), usePathname() || window.location.pathname);

  return <SuperTokensWrapper>{children}</SuperTokensWrapper>;
};