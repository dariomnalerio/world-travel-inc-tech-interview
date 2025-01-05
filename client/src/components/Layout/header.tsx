import { CircleUserRound, LogOut, PawPrint } from "lucide-react";
import { JSX } from "react";
import styles from "./header.module.css";
import { Button } from "../ui/button/button";
import { useView } from "../../hooks/use-view";
import { useAuth } from "../../hooks/use-auth";

const LoggedInNav = (): JSX.Element => {
  const { logout } = useAuth();
  const { currentView, changeView } = useView();

  const handleLogout = () => {
    logout();
    if (currentView === "profile") {
      changeView("home");
    }
  };

  const handleProfile = () => {
    changeView("profile");
  };
  return (
    <nav>
      <Button
        data-testid="profileBtn"
        onClick={handleProfile}
        className={styles.linkBtn}
      >
        <CircleUserRound size={16} />
        Profile
      </Button>
      <Button
        data-testid="logoutBtn"
        data-variant="secondary"
        onClick={handleLogout}
        className={styles.linkBtn}
      >
        <LogOut size={16} />
        Logout
      </Button>
    </nav>
  );
};

const LoggedOutNav = (): JSX.Element => {
  const { changeView } = useView();
  return (
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
  );
};

interface HeaderProps {
  title?: string;
}

const Header = ({ title = "Default Title" }: HeaderProps): JSX.Element => {
  const { changeView } = useView();
  const { userId } = useAuth();
  return (
    <header className={styles.header}>
      <button
        className={styles.logoContainer}
        onClick={() => changeView("home")}
      >
        <PawPrint data-testid="logo" className={styles.logo} size={32} />
        <span className={styles.titleText}>{title}</span>
      </button>
      {userId ? <LoggedInNav /> : <LoggedOutNav />}
    </header>
  );
};

export default Header;
