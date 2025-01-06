import { JSX } from "react";
import styles from "./error.module.css";

const Error = (): JSX.Element => {
  return (
    <div className={styles.errorContainer}>
      <h1 className={styles.errorTitle}>Something went wrong</h1>
      <p className={styles.errorMessage}>
        We're sorry, but an unexpected error has occurred.
      </p>
    </div>
  );
};

export default Error;
