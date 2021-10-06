export interface response_apiGetUserInfo {
  nickName: string;
  country: string;
  address: string;
  email: string;
  countryCode: string;
  phone: string;
}
export interface recordData {
  ID: number;
  ModuleName: string;
  StakeSGB: string;
  PeriodOfUseMonth: number;
  StartTime: string;
  EndTime: string;
  NFTHash: string;
  TxHash: string;
  TxStatus: number;
  DoneAt: string;
  CreatedAt: string;
  statusString: string;
}

export interface response_apiGetUserRecord {
  count: number;
  list: recordData[];
  stakedAmount: string;
  withrawn: string;
}
