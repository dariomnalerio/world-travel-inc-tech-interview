import { useCallback, useState } from "react";
import { getRandomDog } from "../api/random-dog";
import { useAuth } from "./use-auth";

const useGetRandomDog = () => {
  const { userId } = useAuth();
  const [isFetching, setIsFetching] = useState(false);
  const [error, setError] = useState("");
  const [currentUrl, setCurrentUrl] = useState("");
  const [liked, setLiked] = useState(false);

  const fetchNextDog = useCallback(async () => {
    setIsFetching(true);
    const { data, error } = await getRandomDog(userId!);
    setIsFetching(false);
    if (error) {
      setError(error.message);
      return;
    }

    setCurrentUrl(data.imageUrl);
    setLiked(data.liked);
  }, [userId]);

  return {
    isFetching,
    error,
    currentUrl,
    liked,
    fetchNextDog,
  };
};

export { useGetRandomDog };
