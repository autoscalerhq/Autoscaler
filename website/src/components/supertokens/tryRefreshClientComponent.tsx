'use client';
import {useRouter} from 'next/navigation';
import {useEffect, useState} from 'react';
import Session from 'supertokens-auth-react/recipe/session';
import SuperTokens from 'supertokens-auth-react';
import {reportException} from '~/lib/errors';

export const TryRefreshClientComponent = () => {
  const router = useRouter();
  const [didError, setDidError] = useState(false);

  useEffect(() => {
    async function attemptRefreshingSession() {
      try {
        const hasSession = await Session.attemptRefreshingSession();
        if (hasSession) {
          router.refresh();
        } else {
          await SuperTokens.redirectToAuth();
        }
      } catch (ex) {
        reportException(ex);
        setDidError(true);
      }
    }

    void attemptRefreshingSession();
  }, [router]);

  if (didError) {
    return <div>Something went wrong, please reload the page</div>;
  }

  return <div>Loading...</div>;
};