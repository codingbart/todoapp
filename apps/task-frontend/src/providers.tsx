import { type ReactNode } from 'react';
import { Provider as StoreProvider } from 'react-redux';
import { BrowserRouter } from 'react-router-dom';
import { TooltipProvider } from '@/components/ui/tooltip';
import { AuthProvider } from '@/providers/auth-provider';
import { store } from '@/lib/store';

type ProvidersProps = {
    children: ReactNode;
};

export function Providers({ children }: ProvidersProps) {
    return (
        <AuthProvider>
            <StoreProvider store={store}>
                <BrowserRouter>
                    <TooltipProvider>{children}</TooltipProvider>
                </BrowserRouter>
            </StoreProvider>
        </AuthProvider>
    );
}
