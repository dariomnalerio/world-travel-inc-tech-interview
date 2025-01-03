import React, { JSX } from "react";
import styles from "./button.module.css";
type ButtonProps = React.ComponentPropsWithRef<"button">;

const Button = (props: ButtonProps): JSX.Element => {
  return <button className={styles.button} {...props} />;
};

export { Button };
