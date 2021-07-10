import axios, { AxiosResponse } from "axios";
import { RequestBody, ResponseBody } from "./models/urlModels";

let uri = process.env.REACT_APP_HOST;
const port = process.env.REACT_APP_PORT;

if (port) {
  uri = `${uri}:${port}`;
}

export const fetchPageAnalysisData = (
  data: RequestBody
): Promise<AxiosResponse<ResponseBody>> => {
  return axios.post(`http://${uri}/url-analyze`, data);
};
