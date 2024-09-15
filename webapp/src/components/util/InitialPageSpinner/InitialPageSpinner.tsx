import styles from './InitialPageSpinner.module.css';

export function InitialPageSpinner() {
  return (
    <div className={styles['loading-container']}>
      <div className={styles['initial-page-loader']}></div>
      <div className={styles['loading-text']}>Loading</div>
    </div>
  )
}