"use client";

import styles from "../page.module.css";

import {useApiOnClient, useNextJsApiOnClient} from '@/app/api-client/client-hooks';

export const CallAPIButton = () => {
  const api = useApiOnClient();
  const nextJsApi = useNextJsApiOnClient();
  const fetchUserData = async () => {

    const userInfoResponse = await nextJsApi.getUser();
    const anotherResponse = await api.getComment()

    alert(JSON.stringify(await userInfoResponse.json()));
    alert(JSON.stringify(await anotherResponse.text()));
  };

  return (
    <div onClick={fetchUserData} className={styles.sessionButton}>
      Call API
    </div>
  );
};
