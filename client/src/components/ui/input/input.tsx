import type { JSX } from "react";
import styles from "./input.module.css";
type InputProps = React.ComponentPropsWithRef<"input">;
const Input = (props: InputProps): JSX.Element => {
  return <input className={styles.input} {...props} />;
};

export { Input };
