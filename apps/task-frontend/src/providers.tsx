import { type ReactNode } from 'react';
import { Provider } from 'react-redux';
import { BrowserRouter } from 'react-router-dom';
import { TooltipProvider } from '@/components/ui/tooltip';
import { store } from '@/lib/store';

type ProvidersProps = {
    children: ReactNode;
};

export function Providers({ children }: ProvidersProps) {
    return (
        <Provider store={store}>
            <BrowserRouter>
                <TooltipProvider>{children}</TooltipProvider>
            </BrowserRouter>
        </Provider>
    );
}
