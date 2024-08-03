"use client";

import styles from "../page.module.css";

export const CallAPIButton = () => {
  const fetchUserData = async () => {
    const userInfoResponse = await fetch("http://localhost:3000/api/user");
    const anotherResponse = await fetch("http://localhost:4000/comment", {
      headers: {

      }
    });

    alert(JSON.stringify(await userInfoResponse.json()));
    alert(JSON.stringify(await anotherResponse.text()));
  };

  return (
    <div onClick={fetchUserData} className={styles.sessionButton}>
      Call API
    </div>
  );
};
