import { PawPrint } from "lucide-react";
import { JSX } from "react";
import styles from "./header.module.css";
import { Button } from "../ui/button/button";
import { useView } from "../../hooks/use-view";

interface HeaderProps {
  title?: string;
}

const Header = ({ title = "Default Title" }: HeaderProps): JSX.Element => {
  const { changeView } = useView();
  return (
    <header className={styles.header}>
      <button
        className={styles.logoContainer}
        onClick={() => changeView("home")}
      >
        <PawPrint data-testid="logo" className={styles.logo} size={32} />
        <span>{title}</span>
      </button>
      <nav>
        <Button
          data-testid="loginBtn"
          onClick={() => changeView("login")}
          className={styles.linkBtn}
        >
          Login
        </Button>
        <Button
          data-testid="registerBtn"
          data-variant="secondary"
          onClick={() => changeView("register")}
          className={styles.linkBtn}
        >
          Register
        </Button>
      </nav>
    </header>
  );
};

export default Header;
