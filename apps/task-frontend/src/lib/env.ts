import { from } from 'env-var';

const e = from(import.meta.env);

export const env = {
    VITE_API_URL: e.get('VITE_API_URL').default('http://localhost:8080').asUrlString(),
    VITE_KEYCLOAK_URL: e.get('VITE_KEYCLOAK_URL').default('http://localhost:8080').asUrlString(),
    VITE_KEYCLOAK_REALM: e.get('VITE_KEYCLOAK_REALM').default('todoapp').asString(),
    VITE_KEYCLOAK_CLIENT_ID: e.get('VITE_KEYCLOAK_CLIENT_ID').default('task-frontend').asString()
};
