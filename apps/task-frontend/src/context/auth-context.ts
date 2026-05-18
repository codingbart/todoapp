import { createContext } from 'react';

export type AuthUser = {
    id: string;
    name: string;
    email: string;
};

export type AuthContextValue = {
    user: AuthUser;
    logout: () => void;
};

export const AuthContext = createContext<AuthContextValue | null>(null);
