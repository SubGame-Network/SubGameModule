import useAppContext from "../useAppContext";

export interface IError {
  code: number;
  message: string;
}

const useErrorHandler = () => {
  const { showFeedBack } = useAppContext();

  const onError = (error: IError) => {};

  const mutationOnError = (error: IError) => {
    if (error?.code === 5005 || error?.code === 4002) {
    } else {
      showFeedBack("Failed", error?.code?.toString() || "500");
    }
  };

  return { onError, mutationOnError };
};

export default useErrorHandler;
