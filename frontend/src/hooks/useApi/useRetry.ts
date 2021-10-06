import useAppContext from "../../hooks/useAppContext";

const useRetry = () => {
  const { showFeedBack } = useAppContext();

  return (failureCount: number, err: any) => {
    if (err?.code === 5005) {
      return false;
    }

    if (failureCount >= 3) {
      showFeedBack(err?.message || JSON.stringify(err));
      return false;
    }
    return true;
  };
};

export default useRetry;
