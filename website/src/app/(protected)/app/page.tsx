import {HomeComponent} from '~/components/auth/home';
import styles from "./page.module.css";
import {useApiOnServer} from '~/api-client/server-hooks';

export default async function Home() {
  const api = useApiOnServer();
  const comment = await api.getComment()
  return (
      <main className={styles.main}>
        <span>{comment}</span>
        <HomeComponent />
      </main>
  );
}

