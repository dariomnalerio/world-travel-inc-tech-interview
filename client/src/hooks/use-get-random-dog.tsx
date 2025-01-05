import { useCallback, useState } from "react";
import { getRandomDog } from "../api/random-dog";

const useGetRandomDog = () => {
  const [isFetching, setIsFetching] = useState(false);
  const [error, setError] = useState("");
  const [currentUrl, setCurrentUrl] = useState("");

  const fetchNextDog = useCallback(async () => {
    setIsFetching(true);
    const { data, error } = await getRandomDog();
    setIsFetching(false);
    if (error) {
      setError(error.message);
      return;
    }

    setCurrentUrl(data.image_url);
  }, []);

  return {
    isFetching,
    error,
    currentUrl,
    fetchNextDog,
  };
};

export { useGetRandomDog };
