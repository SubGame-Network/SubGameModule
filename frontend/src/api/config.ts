import axios from "axios";
import { Cookies } from "react-cookie";
import { AxiosRequestConfig, AxiosResponse } from "axios";

const cookie = new Cookies();
export interface api_response<T = any> {
  code: number;
  data?: T;
  message?: string;
}

export const headers = {
  "Content-Type": "application/json",
};

const axiosConfig = {
  baseURL: "/api/v1",
  transformRequest: (data: any) => JSON.stringify(data),
  timeout: 60000,
  headers,
};
// const axiosConfig_customized = {
//   baseURL: "/api/customized",
//   transformRequest: (data: any) => JSON.stringify(data),
//   timeout: 20000,
//   headers,
// };
// export const interceptors_request = [
//   (request: AxiosRequestConfig) => {
//     return {
//       ...request,
//       headers: {
//         ...request.headers,
//         Authorization: `Bearer ${cookie.get(`subgame_stake_token`)}`,
//       },
//     };
//   },
// ];

export const interceptors_response = [
  (response: AxiosResponse<any>) => {
    return response;
  },
  (error: any) => {
    throw error.response?.data;
  },
];
export const formatQueryString = (data: { [key: string]: any }) => {
  return Object.entries(data)
    .map(([key, value]) => `${key}=${value}`)
    .join("&");
};

export const getAxiosInstance = () => {
  const axiosInstance = axios.create(axiosConfig);
  axiosInstance.interceptors.response.use(...interceptors_response);
  axiosInstance.interceptors.request.use((request) => {
    request.headers["Address"] = cookie.get("subgame_module_address");
    return request;
  });

  return axiosInstance;
};

export const fromDataConfig: AxiosRequestConfig = {
  ...axiosConfig,
  transformRequest: (data: any) => {
    const formData = new FormData();
    for (const key in data) {
      formData.append(key, data[key]);
    }
    return formData;
  },
  headers: {
    ...headers,
    "Content-Type": "multipart/form-data",
  },
};
