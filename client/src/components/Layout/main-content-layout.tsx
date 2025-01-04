import { JSX } from "react";
import styles from "./main-content-layout.module.css";
type MainContentLayoutProps = {
  children: React.ReactNode;
};

const MainContentLayout = ({
  children,
}: MainContentLayoutProps): JSX.Element => {
  return <main className={styles.layout}>{children}</main>;
};

export default MainContentLayout;
