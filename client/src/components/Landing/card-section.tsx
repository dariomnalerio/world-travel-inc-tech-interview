import { JSX } from "react";
import { Card } from "../ui/card/card";
import { Heart, RefreshCcw } from "lucide-react";
import styles from "./landing.module.css";
import { Button } from "../ui/button/button";

export type CardSectionProps = {
  isFetching: boolean;
  currentUrl: string;
  likeCurrentDog: boolean;
  fetchNextDog: () => void;
  handleLike: () => void;
};

const CardSection = ({
  currentUrl,
  fetchNextDog,
  handleLike,
  isFetching,
  likeCurrentDog,
}: CardSectionProps): JSX.Element => {
  return (
    <Card data-testid="card" className={styles.card}>
      <Card.Content className={styles.cardContent}>
        {isFetching && !currentUrl ? (
          <div className={styles.imgLoading}></div>
        ) : (
          <img className={styles.img} src={currentUrl} alt="Random Dog" />
        )}
      </Card.Content>
      <Card.Footer className={styles.cardFooter}>
        <Button
          className={`${styles.styledBtn} ${styles.heartIcon}`}
          disabled={isFetching}
          data-variant="secondary"
          onClick={handleLike}
        >
          <Heart
            data-testid="heartIcon"
            fill="#ffffff"
            fillOpacity={likeCurrentDog ? 1 : 0}
            size={16}
          />
          Like
        </Button>
        <Button
          className={styles.styledBtn}
          disabled={isFetching}
          onClick={fetchNextDog}
        >
          <RefreshCcw
            className={styles.refreshIcon}
            data-testid="refreshIcon"
            data-spin={isFetching ? "true" : "false"}
            size={16}
          />
          Next
        </Button>
      </Card.Footer>
    </Card>
  );
};

export { CardSection };
