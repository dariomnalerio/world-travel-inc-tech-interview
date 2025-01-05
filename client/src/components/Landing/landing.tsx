import MainContentLayout from "../Layout/main-content-layout";
import styles from "./landing.module.css";
import { useGetRandomDog } from "../../hooks/use-get-random-dog";
import { JSX, useCallback, useEffect, useState } from "react";
import { CardSection } from "./card-section";

const Landing = (): JSX.Element => {
  const { fetchNextDog, isFetching, currentUrl } = useGetRandomDog();
  const [likeCurrentDog, setLikeCurrentDog] = useState(false);

  const handleLike = () => {
    setLikeCurrentDog((prev) => !prev);
  };

  const handleFetchNextDog = useCallback(() => {
    fetchNextDog();
    setLikeCurrentDog(false);
  }, [fetchNextDog]);

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === "Enter" || e.key === "Space" || e.key === "ArrowRight") {
      handleFetchNextDog();
    }
  };

  useEffect(() => {
    // fetches twice in development
    if (!currentUrl) {
      handleFetchNextDog();
    }
  }, [currentUrl, handleFetchNextDog]);

  return (
    <MainContentLayout onKeyDown={handleKeyDown}>
      <section className={styles.section}>
        <div className={styles.sectionHeader}>
          <h1 className={styles.title}>Random Dog</h1>
          <button className={styles.styledBtn} data-variant="accent">
            Show Liked Dogs
          </button>
        </div>

        <CardSection
          {...{
            currentUrl,
            fetchNextDog: handleFetchNextDog,
            handleLike,
            isFetching,
            likeCurrentDog,
          }}
        />
      </section>
    </MainContentLayout>
  );
};

export default Landing;
