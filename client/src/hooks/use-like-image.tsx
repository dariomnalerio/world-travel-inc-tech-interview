import { useEffect, useOptimistic, useState, useTransition } from "react";
import {
  likeDogImage as likeDog,
  unlikeDogImage as unlikeDog,
} from "../api/liked-images";

type LikeState = {
  isLiked: boolean;
  isLoading?: boolean;
};

/**
 * Custom hook to handle liking and unliking an image.
 * It provides the current like state and functions to like and unlike an image.
 * It performs optimistic updates to provide a better user experience.
 */
const useLikeImage = (
  userId: string,
  imageUrl: string,
  initialLiked: boolean = false
) => {
  const [isPending, startTransition] = useTransition();
  const [likeState, setLikeState] = useState<LikeState>({
    isLiked: initialLiked,
    isLoading: false,
  });
  const [optimisticLikeState, addOptimisticLike] = useOptimistic(
    likeState,
    (state: LikeState, optimisticValue: boolean) => ({
      ...state,
      isLiked: optimisticValue,
      isLoading: true,
    })
  );

  const likeDogImage = async () => {
    try {
      startTransition(async () => {
        addOptimisticLike(true);
        const { error } = await likeDog(userId, imageUrl);
        if (error) {
          setLikeState((prev) => ({ ...prev, isLiked: false }));
          return;
        }
        setLikeState((prev) => ({ ...prev, isLiked: true }));
      });
      return;
    } catch (error) {
      setLikeState((prev) => ({ ...prev, isLiked: false }));
      throw error;
    }
  };

  const unlikeDogImage = async () => {
    try {
      startTransition(async () => {
        addOptimisticLike(false);
        const { error } = await unlikeDog(userId, imageUrl);
        if (error) {
          setLikeState((prev) => ({ ...prev, isLiked: true }));
          return;
        }
        setLikeState((prev) => ({ ...prev, isLiked: false }));
      });
      return;
    } catch (error) {
      console.error("Unlike dog image error:", error);
      setLikeState((prev) => ({ ...prev, isLiked: true }));
      throw error;
    }
  };

  useEffect(() => {
    setLikeState({
      isLiked: initialLiked,
      isLoading: false,
    });
  }, [imageUrl, initialLiked]);

  return {
    isLiked: optimisticLikeState.isLiked,
    isLoading: optimisticLikeState.isLoading || isPending,
    likeDogImage,
    unlikeDogImage,
  };
};

export { useLikeImage };
