import {HomeComponent} from './components/home';
import styles from "./page.module.css";
import {AuthWrapper} from '@/app/config/supertokens/components/SuperTokensAuthWrapper';
import {useApiOnServer} from '@/app/api-client/server-hooks';

export default async function Home() {
  const api = useApiOnServer();
  const comment = await api.getComment()
  return (
    <AuthWrapper>
      <main className={styles.main}>
        <span>{comment}</span>
        <HomeComponent />
      </main>
    </AuthWrapper>
  );
}
