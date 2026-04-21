import { cva, type VariantProps } from 'class-variance-authority';

import { cn } from '@/lib/utils';

const badgeVariants = cva('inline-flex items-center rounded-md px-2 py-0.5 text-xs font-medium', {
    variants: {
        variant: {
            default: 'bg-primary/10 text-primary',
            secondary: 'bg-secondary text-secondary-foreground',
            success: 'bg-green-100 text-green-700 dark:bg-green-900/30 dark:text-green-400',
            warning: 'bg-yellow-100 text-yellow-700 dark:bg-yellow-900/30 dark:text-yellow-400',
            destructive: 'bg-destructive/10 text-destructive',
            outline: 'border border-border text-foreground',
        },
    },
    defaultVariants: { variant: 'default' },
});

function Badge({
    className,
    variant,
    ...props
}: React.ComponentProps<'span'> & VariantProps<typeof badgeVariants>) {
    return <span className={cn(badgeVariants({ variant }), className)} {...props} />;
}

export { Badge, badgeVariants };
