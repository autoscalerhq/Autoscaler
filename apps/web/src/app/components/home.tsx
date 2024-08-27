"use client"
import styles from "../page.module.css";
import Image from "next/image";
import { CelebrateIcon, SeparatorLine } from "../../assets/images";
import { CallAPIButton } from "./callApiButton";
import { LinksComponent } from "./linksComponent";
import React from 'react';
import {useAuthContext} from '@/app/config/supertokens/components/AuthContext';


export function HomeComponent() {

  const {accessTokenPayload} = useAuthContext();
  return (
      <div className={styles.homeContainer}>
        <div className={styles.mainContainer}>
          <div
            className={`${styles.topBand} ${styles.successTitle} ${styles.bold500}`}
          >
            <Image
              src={CelebrateIcon}
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
          src={SeparatorLine}
          alt="separator"
        />
      </div>
  );
}
