import { from } from 'env-var';

const env = from(import.meta.env);

export const API_URL = env.get('VITE_API_URL').default('http://localhost:8080').asUrlString();
