"use client";
import styles from "~/app/(protected)/app/page.module.css";
import { BlogsIcon, GuideIcon, SignOutIcon } from "../../assets/images";
import { recipeDetails } from "~/config/supertokens/frontend";
import Link from "next/link";
import Image from "next/image";
import Session from "supertokens-auth-react/recipe/session";
import SuperTokens from "supertokens-auth-react";

const SignOutLink = (props: { name: string; link: string; icon: string }) => {
  return (
    <div
      className={styles.linksContainerLink}
      onClick={async () => {
        await Session.signOut();
        void SuperTokens.redirectToAuth();
      }}
    >
      <Image className={styles.linkIcon} src={props.icon} alt={props.name} />
      <div role={"button"}>{props.name}</div>
    </div>
  );
};

export const LinksComponent = () => {
  const links: {
    name: string;
    link: string;
    icon: string;
  }[] = [
    {
      name: "Blogs",
      link: "https://supertokens.com/blog",
      icon: "/images/blogs-icon.svg",
    },
    {
      name: "Guides",
      link: recipeDetails.docsLink,
      icon: "/images/guide-icon.svg",
    },
    {
      name: "Sign Out",
      link: "",
      icon: "/images/sign-out-icon.svg",
    },
  ];

  return (
    <div className={styles.bottomLinksContainer}>
      {links.map((link) => {
        if (link.name === "Sign Out") {
          return (
            <SignOutLink
              name={link.name}
              link={link.link}
              icon={link.icon}
              key={link.name}
            />
          );
        }

        return (
          <Link
            href={link.link}
            className={styles.linksContainerLink}
            key={link.name}
            target="_blank"
          >
            <Image
              className={styles.linkIcon}
              src={link.icon}
              alt={link.name}
            />
            <div role={"button"}>{link.name}</div>
          </Link>
        );
      })}
    </div>
  );
};
