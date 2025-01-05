import { JSX } from "react";
import { Card } from "../ui/card/card";
import { Heart, Loader2, RefreshCcw } from "lucide-react";
import styles from "./landing.module.css";
import { Button } from "../ui/button/button";
import { Tooltip } from "../ui/tooltip/tooltip";
import { useAuth } from "../../hooks/use-auth";

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
  const { userId } = useAuth();
  const tooltipText = !userId ? "Login to like dogs" : "";

  return (
    <Card data-testid="card" className={styles.card}>
      <Card.Content className={styles.cardContent}>
        {isFetching && !currentUrl ? (
          <div className={styles.imgLoading}>
            <Loader2 />
          </div>
        ) : (
          <img className={styles.img} src={currentUrl} alt="Random Dog" />
        )}
      </Card.Content>
      <Card.Footer className={styles.cardFooter}>
        <Tooltip text={tooltipText}>
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
        </Tooltip>
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
