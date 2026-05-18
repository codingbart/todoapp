// https://orval.dev/docs/guides/custom-axios/

import Axios, { type AxiosRequestConfig, AxiosError } from 'axios';

const _axiosInstance = Axios.create({
    baseURL: import.meta.env.VITE_API_URL ?? 'http://localhost:8080'
});

_axiosInstance.interceptors.request.use(
    config => {
        // TODO: add jwt to header
        return config;
    },
    error => Promise.reject(error)
);

_axiosInstance.interceptors.response.use(
    response => response,
    error => {
        // TODO: refresh jwt
        return Promise.reject(error);
    }
);

export async function axiosInstance<T>(
    config: AxiosRequestConfig,
    options?: AxiosRequestConfig
): Promise<T> {
    const { data } = await _axiosInstance({ ...config, ...options });
    return data;
}

export type ErrorType<Error> = AxiosError<Error>;
export type BodyType<BodyData> = BodyData;
