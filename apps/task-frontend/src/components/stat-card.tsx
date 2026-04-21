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
            <CardHeader>
                <CardTitle>{label}</CardTitle>
            </CardHeader>
            <CardContent>
                <p className={cn('text-3xl font-bold', variantStyles[variant])}>{value}</p>
            </CardContent>
        </Card>
    );
}
