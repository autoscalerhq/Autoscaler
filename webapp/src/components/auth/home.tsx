"use client"
import styles from "~/app/(protected)/app/page.module.css";
import Image from "next/image";
import { CelebrateIcon, SeparatorLine } from "../../assets/images";
import { CallAPIButton } from "./callApiButton";
import { LinksComponent } from "./linksComponent";
import React from 'react';
import {useAuthContext} from '~/components/supertokens/AuthContext';

export function HomeComponent() {

  const {accessTokenPayload} = useAuthContext();
  return (
      <div className={styles.homeContainer}>
        <div className={styles.mainContainer}>
          <div
            className={`${styles.topBand} ${styles.successTitle} ${styles.bold500}`}
          >
            <Image
              src={"/images/celebrate-icon.svg"}
              width={31}
              height={31}
              alt="Login successful"
              className={styles.successIcon}
            />{" "}
            Login successful
          </div>
          <div className={styles.innerContent}>
            <div>Your userID is:</div>
            <div className={`${styles.truncate} ${styles.userId}`}>
              {accessTokenPayload.sub}
            </div>
            <CallAPIButton />
          </div>
        </div>
        <LinksComponent />
        <Image
          className={styles.separatorLine}
          src={"images/separator-line.svg"}
          width={530}
          height={1}
          alt="separator"
        />
      </div>
  );
}
