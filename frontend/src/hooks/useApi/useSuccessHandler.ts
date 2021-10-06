// import { AxiosResponse } from "axios";
import { useQueryClient } from "react-query";

import useAppContext from "../useAppContext";
// import { api_response } from "../../api/config";

export interface IError {
  code: number;
  message: string;
}

const useSuccessHandler = () => {
  const queryClient = useQueryClient();
  const { showFeedBack } = useAppContext();

  const mutationOnSuccess = (
    res: any,
    clearKeyName?: string | string[],
    setModalShow?: (a: boolean) => void
  ) => {
    showFeedBack("Submitted", "Submitted");
    if (Array.isArray(clearKeyName)) {
      clearKeyName.forEach((key) => {
        queryClient.invalidateQueries(key);
      });
    } else if (clearKeyName) {
      queryClient.invalidateQueries(clearKeyName);
    }
    if (setModalShow) setModalShow(false);
  };

  return { mutationOnSuccess };
};

export default useSuccessHandler;
