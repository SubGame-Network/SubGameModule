import { useQuery } from "react-query";

import { apiGetModules, apiGetModuleDetail } from "../../api/module";
import useRetry from "./useRetry";

export const useApiGetModulesData = (enabled?: boolean) => {
  const retry = useRetry();
  return useQuery(["modules"], () => apiGetModules(), {
    retry,
    onError: () => {},
    enabled,
    keepPreviousData: true,
  });
};

export const useApiGetModuleDetail = (id: number, enabled?: boolean) => {
  const retry = useRetry();

  return useQuery(["moduledetail", id], () => apiGetModuleDetail(id), {
    retry,
    onError: () => {},
    enabled,
  });
};
