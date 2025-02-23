import axios, {
  type AxiosResponse,
  type InternalAxiosRequestConfig,
} from "axios";

const API_URL = `${import.meta.env.VITE_API_URL}`;

export const ApiInstance = axios.create({
  baseURL: API_URL,
  withCredentials: true,
});

export class API {}

export const sessionKeyRequestInterceptor = (
  config: InternalAxiosRequestConfig<any>,
) => {
  return config;
};

export const fulFilledResponseInterceptor = (response: AxiosResponse) =>
  response;
export const rejectedResponseInterceptor = async (e: any) => {
  console.error(e);
  // if (error.response && error.response.status == 401 && error.config) {
  //   try {
  //   } catch (e) {}
  // } else if (error.response && error.response.status == 404) {
  //   console.error("Проверьте url");
  // } else {
  //   console.error(error.toString());
  // }
  throw e;
};

ApiInstance.interceptors.request.use(sessionKeyRequestInterceptor);

ApiInstance.interceptors.response.use(
  fulFilledResponseInterceptor,
  rejectedResponseInterceptor,
);

export const getTextError = (error: any) => {
  if (axios.isAxiosError(error)) {
    //TODO узнать как будет на беке описываться
    return error.response?.data.error_description || "";
  }
};
