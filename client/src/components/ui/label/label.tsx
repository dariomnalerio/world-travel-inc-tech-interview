import React, { type JSX } from "react";
import styles from "./label.module.css";
type LabelProps = React.ComponentPropsWithRef<"label">;

const Label = (props: LabelProps): JSX.Element => {
  return <label className={styles.label} {...props} data-testid="label" />;
};

export { Label };
