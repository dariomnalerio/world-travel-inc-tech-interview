import { JSX } from "react";
import MainContentLayout from "../Layout/main-content-layout";
import styles from "./profile.module.css";
import { LikedDogsSection } from "./liked-dogs";
import { useView } from "../../hooks/use-view";

const Profile = (): JSX.Element => {
  const { changeView } = useView();

  const handleGoToLanding = () => {
    changeView("home");
  };
  return (
    <MainContentLayout>
      <section className={styles.section}>
        <div className={styles.sectionHeader}>
          <h1 className={styles.title}>Liked Dogs</h1>
          <button
            className={styles.styledBtn}
            data-variant="accent"
            onClick={handleGoToLanding}
          >
            Show Random Dogs
          </button>
        </div>

        {/*  liked dogs */}
        <LikedDogsSection />
      </section>
    </MainContentLayout>
  );
};

export default Profile;
