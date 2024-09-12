import styles from "./page.module.css";
import Link from 'next/link';

export default async function Home() {
  return (
   <main className={styles.main}>
     Home Page
     <Link href={'/app'}>Go to protected page</Link>
   </main>
  );
}

