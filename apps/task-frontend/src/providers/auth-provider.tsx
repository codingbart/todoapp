import { useEffect, useState, type ReactNode } from 'react';
import { SpinnerScreen } from '@/components/spinner-screen';
import { z } from 'zod';
import { keycloak } from '@/lib/keycloak';
import { AuthContext, type AuthContextValue } from '@/context/auth-context';
import { FadeIn } from '@/components/fade-in';

const ClaimsSchema = z.object({
    sub: z.string(),
    name: z.string(),
    email: z.string()
});

export function AuthProvider({ children }: { children: ReactNode }) {
    const [isLoading, setIsLoading] = useState(true);

    useEffect(() => {
        if (keycloak.didInitialize) {
            return;
        }

        keycloak.onTokenExpired = () => {
            keycloak.updateToken(30).catch(() => keycloak.logout());
        };

        keycloak.init({ onLoad: 'login-required' }).then(() => setIsLoading(false));
    }, []);

    if (isLoading) {
        return <SpinnerScreen />;
    }

    const claims = ClaimsSchema.parse(keycloak.tokenParsed);

    const value: AuthContextValue = {
        user: {
            id: claims.sub,
            name: claims.name,
            email: claims.email
        },
        logout: () => keycloak.logout()
    };

    return (
        <AuthContext.Provider value={value}>
            <FadeIn>{children}</FadeIn>
        </AuthContext.Provider>
    );
}
