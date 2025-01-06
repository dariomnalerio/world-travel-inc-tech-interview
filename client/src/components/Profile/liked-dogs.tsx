import { useEffect, useState, type JSX } from "react";
import { DogCard } from "./dog-card";
import { getLikedDogImages } from "../../api/liked-images";
import { useAuth } from "../../hooks/use-auth";
import { useView } from "../../hooks/use-view";
import styles from "./profile.module.css";

const LikedDogsSection = (): JSX.Element => {
  const { changeView } = useView();
  const { userId } = useAuth();
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [imgs, setImgs] = useState<string[]>([]);

  useEffect(() => {
    if (!userId) {
      changeView("home");
    }
    const fetchLikedImages = async () => {
      try {
        setIsLoading(true);
        const { data, error } = await getLikedDogImages(userId!);
        setIsLoading(false);
        if (error) {
          console.error("Fetch liked images error:", error);
          return;
        }
        setImgs(data.likedImages);
      } catch (error) {
        console.error("Fetch liked images error:", error);
      }
    };

    fetchLikedImages();
  }, [userId, changeView]);

  return (
    <div className={styles.dogCardContainer} data-testid="dogCardContainer">
      {!isLoading && (!imgs || imgs.length === 0) && (
        <h2 className={styles.noLikedDogs}>You have not liked a dog yet.</h2>
      )}
      {!isLoading &&
        imgs &&
        imgs.length > 0 &&
        imgs.map((img) => {
          return <DogCard key={img} url={img} liked={true} />;
        })}
    </div>
  );
};

export { LikedDogsSection };
