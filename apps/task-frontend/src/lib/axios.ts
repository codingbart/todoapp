// https://orval.dev/docs/guides/custom-axios/

import Axios, { type AxiosRequestConfig, AxiosError } from 'axios';
import { env } from '@/lib/env';
import { keycloak } from '@/lib/keycloak';

const _axiosInstance = Axios.create({
    baseURL: env.VITE_API_URL
});

_axiosInstance.interceptors.request.use(
    config => {
        if (keycloak.token) {
            config.headers.Authorization = `Bearer ${keycloak.token}`;
        }

        return config;
    },
    error => Promise.reject(error)
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
