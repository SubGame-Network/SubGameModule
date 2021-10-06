export interface response_apiGetModules {
  id: number;
  name: string;
  depiction: string;
}
export interface porgameData {
  periodOfUse: number;
  amount: string;
  programID: number;
}
export interface response_apiGetModuleDetail {
  module: {
    id: number;
    name: string;
    depiction: string;
    readmeMdUrl: string;
  };
  program: porgameData[];
}
