import MainContentLayout from "../Layout/main-content-layout";
import styles from "./landing.module.css";
import { useGetRandomDog } from "../../hooks/use-get-random-dog";
import { JSX, useCallback, useEffect } from "react";
import { CardSection } from "./card-section";
import { useAuth } from "../../hooks/use-auth";
import { useLikeImage } from "../../hooks/use-like-image";
import { useView } from "../../hooks/use-view";
import { Tooltip } from "../ui/tooltip/tooltip";

const Landing = (): JSX.Element => {
  const { changeView } = useView();
  const { userId } = useAuth();
  const { fetchNextDog, isFetching, currentUrl } = useGetRandomDog();
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

  useEffect(() => {
    // fetches twice in development
    if (!currentUrl) {
      handleFetchNextDog();
    }
  }, [currentUrl, handleFetchNextDog]);

  const handleGoToProfile = () => {
    if (!userId) {
      return;
    }
    changeView("profile");
  };

  const tooltipText = !userId ? "You must be logged in" : "";

  return (
    <MainContentLayout>
      <section className={styles.section}>
        <div className={styles.sectionHeader}>
          <h1 className={styles.title}>Random Dog</h1>
          <Tooltip text={tooltipText}>
            <button
              className={styles.styledBtn}
              data-variant="accent"
              onClick={handleGoToProfile}
            >
              Show Liked Dogs
            </button>
          </Tooltip>
        </div>

        <CardSection
          currentUrl={currentUrl}
          fetchNextDog={handleFetchNextDog}
          handleLike={handleLikeClick}
          isFetching={isFetching}
          likeCurrentDog={isLiked}
        />
      </section>
    </MainContentLayout>
  );
};

export default Landing;
