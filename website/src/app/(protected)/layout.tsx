import React from 'react';
import {AuthWrapper} from '~/components/supertokens/SuperTokensAuthWrapper';

export default function ProtectedLayout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  return (
    <AuthWrapper>
      {children}
    </AuthWrapper>
  );
}
