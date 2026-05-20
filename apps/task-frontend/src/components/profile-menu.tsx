import { LogOut } from 'lucide-react';
import { useAuth } from '@/hooks/use-auth';
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuGroup,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuTrigger
} from '@/components/ui/dropdown-menu';

function initials(name: string): string {
    return name
        .split(' ')
        .slice(0, 2)
        .map(w => w[0])
        .join('')
        .toUpperCase();
}

export function ProfileMenu() {
    const { user, logout } = useAuth();

    return (
        <DropdownMenu>
            <DropdownMenuTrigger className='hover:bg-muted flex items-center gap-2.5 rounded-lg px-2 py-1.5 text-sm transition-colors outline-none'>
                <span className='bg-primary text-primary-foreground flex size-7 shrink-0 items-center justify-center rounded-full text-xs font-semibold'>
                    {initials(user.name)}
                </span>
                <div className='hidden text-left sm:block'>
                    <p className='text-sm leading-tight font-medium'>{user.name}</p>
                    <p className='text-muted-foreground text-xs leading-tight'>{user.email}</p>
                </div>
            </DropdownMenuTrigger>
            <DropdownMenuContent align='end'>
                <DropdownMenuGroup>
                    <DropdownMenuLabel className='text-muted-foreground text-xs'>
                        {user.email}
                    </DropdownMenuLabel>
                </DropdownMenuGroup>
                <DropdownMenuSeparator />
                <DropdownMenuGroup>
                    <DropdownMenuItem variant='destructive' onClick={logout}>
                        <LogOut />
                        Wyloguj
                    </DropdownMenuItem>
                </DropdownMenuGroup>
            </DropdownMenuContent>
        </DropdownMenu>
    );
}
