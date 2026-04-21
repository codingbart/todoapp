import { Outlet } from 'react-router-dom';

import { ProfileMenu } from '@/components/profile-menu';

export function Layout() {
    return (
        <div className='flex min-h-screen flex-col'>
            <header className='flex h-14 shrink-0 items-center justify-between border-b px-6'>
                <span className='font-semibold'>TodoApp</span>
                <ProfileMenu />
            </header>
            <main className='mx-auto w-full max-w-4xl flex-1 p-6'>
                <Outlet />
            </main>
        </div>
    );
}
