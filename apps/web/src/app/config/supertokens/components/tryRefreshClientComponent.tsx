"use client";
import {useRouter} from 'next/navigation';
import {useEffect, useState} from 'react';
import Session from 'supertokens-auth-react/recipe/session';
import SuperTokens from 'supertokens-auth-react';

export const TryRefreshClientComponent = () => {
  const router = useRouter();
  const [didError, setDidError] = useState(false);

  useEffect(() => {
    void Session.attemptRefreshingSession()
      .then((hasSession) => {
        if (hasSession) {
          router.refresh();
        } else {
          SuperTokens.redirectToAuth();
        }
      })
      .catch(() => {
        setDidError(true);
      });
  }, [router]);

  if (didError) {
    return <div>Something went wrong, please reload the page</div>;
  }

  return <div>Loading...</div>;
};