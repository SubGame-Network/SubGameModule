import { getAxiosInstance, api_response } from "./config";
import {
  response_apiGetModules,
  response_apiGetModuleDetail,
} from "./types/module";
const sgbmodule = getAxiosInstance();

export const apiGetModules = async () => {
  return sgbmodule.get<api_response<response_apiGetModules[]>>(`/module`);
};

export const apiGetModuleDetail = async (id: number) => {
  return sgbmodule.get<api_response<response_apiGetModuleDetail>>(
    `/module/${id}`
  );
};
