import {HomeComponent} from './components/home';
import styles from "./page.module.css";
import {AuthWrapper} from '@/app/config/supertokens/components/SuperTokensAuthWrapper';

export default function Home() {
  return (
    <AuthWrapper>
      <main className={styles.main}>
        <HomeComponent />
      </main>
    </AuthWrapper>
  );
}
