import { useMutation } from "react-query";
import {
  apiGetSendVerifyCode,
  apiGetUserJoin,
  apiGetUserInfo,
  apiEditName,
  apiGetUserRecord,
  apiConcat,
} from "../../api/user";
import useErrorHandler from "./useErrorHandler";
import useSuccessHandler from "./useSuccessHandler";
import { useQuery } from "react-query";

import useRetry from "./useRetry";
import { usePolkadotJS } from "@polkadot/api-provider";
export const formatQueryKey = (data: { [key: string]: any }) => {
  return Object.entries(data).map(([key, value]) => `${key}=${value}`);
};

export const useApiConcat = () => {
  const { mutationOnError: onError } = useErrorHandler();
  const { mutationOnSuccess } = useSuccessHandler();

  return useMutation(apiConcat, {
    onError,
    onSuccess: (res) => mutationOnSuccess(res, "concat"),
  });
};

export const useApiGetUserRecord = (
  data: {
    row: string;
    page: string;
    nftId: string;
    status: string;
    periodOfUse: string;
  },
  enabled?: boolean
) =>
  useQuery(
    ["userRecord", ...formatQueryKey(data)],
    () => apiGetUserRecord(data),
    { enabled }
  );

export const useApiEditName = () => {
  const { mutationOnError: onError } = useErrorHandler();
  const { mutationOnSuccess } = useSuccessHandler();

  return useMutation(apiEditName, {
    onError,
    onSuccess: (res) => mutationOnSuccess(res, "editname"),
  });
};

export const useApiGetUserInfo = (enabled?: boolean) => {
  const {
    state: { currentAccount },
  } = usePolkadotJS();
  const retry = useRetry();

  return useQuery(
    ["userinfo", currentAccount?.address],
    () => apiGetUserInfo(),
    {
      retry,
      onError: () => {},
      enabled,
    }
  );
};

export const useApiSendVerifyCode = () => {
  const { mutationOnError: onError } = useErrorHandler();
  const { mutationOnSuccess } = useSuccessHandler();

  return useMutation(apiGetSendVerifyCode, {
    onError,
    onSuccess: (res) => mutationOnSuccess(res, "sendverifycode"),
  });
};

export const useApiUserJoin = () => {
  const { mutationOnError: onError } = useErrorHandler();
  const { mutationOnSuccess } = useSuccessHandler();

  return useMutation(apiGetUserJoin, {
    onError,
    onSuccess: (res) => mutationOnSuccess(res, "userjoin"),
  });
};
