import MainContentLayout from "../Layout/main-content-layout";
import styles from "./landing.module.css";
import { useGetRandomDog } from "../../hooks/use-get-random-dog";
import { JSX, useCallback, useEffect } from "react";
import { CardSection } from "./card-section";
import { useAuth } from "../../hooks/use-auth";
import { useLikeImage } from "../../hooks/use-like-image";

const Landing = (): JSX.Element => {
  const { userId } = useAuth();
  const { fetchNextDog, isFetching, currentUrl } = useGetRandomDog();
  // TODO: handle unlogged user case
  const { isLiked, likeDogImage, unlikeDogImage } = useLikeImage(
    userId!,
    currentUrl
  );

  const handleLikeClick = async () => {
    try {
      if (!userId) {
        return;
      }
      if (isLiked) {
        await unlikeDogImage();
      } else {
        await likeDogImage();
      }
    } catch (error) {
      console.error("Like dog image error:", error);
    }
  };

  const handleFetchNextDog = useCallback(() => {
    fetchNextDog();
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
            handleLike: handleLikeClick,
            isFetching,
            likeCurrentDog: isLiked,
          }}
        />
      </section>
    </MainContentLayout>
  );
};

export default Landing;
