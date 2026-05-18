import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { cn } from '@/lib/utils';

const variantStyles = {
    default: '',
    warning: 'text-yellow-600 dark:text-yellow-400',
    success: 'text-green-600 dark:text-green-400',
} as const;

type StatCardProps = {
    label: string;
    value: number;
    variant?: keyof typeof variantStyles;
};

export function StatCard({ label, value, variant = 'default' }: StatCardProps) {
    return (
        <Card>
            <CardHeader className='p-3 pb-1 sm:p-4 sm:pb-2'>
                <CardTitle className='text-xs sm:text-sm'>{label}</CardTitle>
            </CardHeader>
            <CardContent className='p-3 pt-0 sm:p-4 sm:pt-0'>
                <p className={cn('text-2xl font-bold sm:text-3xl', variantStyles[variant])}>
                    {value}
                </p>
            </CardContent>
        </Card>
    );
}
