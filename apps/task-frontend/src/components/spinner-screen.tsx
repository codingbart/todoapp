import { Spinner } from '@/components/ui/spinner';
import { FadeIn } from './fade-in';

export function SpinnerScreen() {
    return (
        <FadeIn>
            <div className='flex h-screen items-center justify-center'>
                <Spinner className='size-12' />
            </div>
        </FadeIn>
    );
}
