import { JSX } from "react";
import { Card } from "../ui/card/card";
import { Button } from "../ui/button/button";
import styles from "./profile.module.css";
import { Heart } from "lucide-react";
import { useAuth } from "../../hooks/use-auth";
import { useLikeImage } from "../../hooks/use-like-image";

interface DogCardProps {
  url: string;
  liked: boolean;
}

const DogCard = ({ url, liked }: DogCardProps): JSX.Element => {
  const { userId } = useAuth();
  const { isLiked, likeDogImage, unlikeDogImage } = useLikeImage(
    userId!,
    url,
    liked
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

  return (
    <Card className={styles.dogCard} data-testid="dogCard">
      <Card.Content>
        <img src={url} alt="Dog" />
      </Card.Content>
      <Card.Footer className={styles.dogCardFooter}>
        <Button
          className={styles.styledBtn}
          data-variant="secondary"
          onClick={handleLikeClick}
        >
          <Heart
            data-testid="heartIcon"
            fill="#2f92e9"
            fillOpacity={isLiked ? 1 : 0}
            size={16}
          />
          Unlike
        </Button>
      </Card.Footer>
    </Card>
  );
};

export { DogCard };
