import type { ReactNode } from 'react';

export function FadeIn({ children }: { children: ReactNode }) {
    return <div className='animate-in fade-in duration-200'>{children}</div>;
}
