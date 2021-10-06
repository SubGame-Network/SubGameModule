import { getAxiosInstance, api_response, formatQueryString } from "./config";
import {
  response_apiGetUserInfo,
  response_apiGetUserRecord,
} from "./types/user";
const sgbmodule = getAxiosInstance();
export const apiGetSendVerifyCode = async (data: { email: string }) => {
  return sgbmodule.post<api_response<any>>(`/user/email/send`, {
    email: data.email,
  });
};

export const apiConcat = async (data: { type: string; contact: string }) => {
  return sgbmodule.post<api_response<any>>(`/contact`, {
    type: data.type,
    contact: data.contact,
  });
};

export const apiGetUserRecord = async (data: {
  row: string;
  page: string;
  nftId: string;
  status: string;
  periodOfUse: string;
}) => {
  return sgbmodule.get<api_response<response_apiGetUserRecord>>(
    `/stake/record?${formatQueryString(data)}`
  );
};

export const apiGetUserJoin = async (data: {
  email: string;
  nickName: string;
  country: string;
  address: string;
  countryCode: string;
  phone: string;
  verifyCode: string;
}) => {
  return sgbmodule.post<api_response<any>>(`/user/join`, {
    email: data.email,
    nickName: data.nickName,
    country: data.country,
    address: data.address,
    countryCode: data.countryCode,
    phone: data.phone,
    verifyCode: data.verifyCode,
  });
};

export const apiGetUserInfo = async () => {
  return sgbmodule.get<api_response<response_apiGetUserInfo>>(`/user`);
};

export const apiEditName = async (data: {
  nickName: string;
  phone: string;
  countryCode: string;
  country: string;
}) => {
  return sgbmodule.patch<api_response<any>>(`/user/name`, {
    nickName: data.nickName,
    phone: data.phone,
    countryCode: data.countryCode,
    country: data.country,
  });
};
