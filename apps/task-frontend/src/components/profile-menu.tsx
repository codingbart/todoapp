import { LogOut, User } from 'lucide-react';
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';

const mockUser = {
    name: 'John Doe',
    email: 'johndoe@example.com',
    initials: 'JD',
};

export function ProfileMenu() {
    return (
        <DropdownMenu>
            <DropdownMenuTrigger className='hover:bg-muted flex items-center gap-2.5 rounded-lg px-2 py-1.5 text-sm transition-colors outline-none'>
                <span className='bg-primary text-primary-foreground flex size-7 shrink-0 items-center justify-center rounded-full text-xs font-semibold'>
                    {mockUser.initials}
                </span>
                <div className='text-left'>
                    <p className='text-sm leading-tight font-medium'>{mockUser.name}</p>
                    <p className='text-muted-foreground text-xs leading-tight'>{mockUser.email}</p>
                </div>
            </DropdownMenuTrigger>
            <DropdownMenuContent align='end'>
                <DropdownMenuLabel className='text-muted-foreground text-xs'>
                    {mockUser.email}
                </DropdownMenuLabel>
                <DropdownMenuSeparator />
                <DropdownMenuItem>
                    <User />
                    Profil
                </DropdownMenuItem>
                <DropdownMenuItem variant='destructive'>
                    <LogOut />
                    Wyloguj
                </DropdownMenuItem>
            </DropdownMenuContent>
        </DropdownMenu>
    );
}
