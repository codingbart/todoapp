import { Trash2 } from 'lucide-react';
import { type VariantProps } from 'class-variance-authority';
import { type Task, type TaskPriority, type TaskStatus } from '@/types/task';
import { Badge, badgeVariants } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';

type BadgeVariant = NonNullable<VariantProps<typeof badgeVariants>['variant']>;

const statusLabels: Record<TaskStatus, string> = {
    TODO: 'Do zrobienia',
    IN_PROGRESS: 'W toku',
    DONE: 'Ukończone',
};

const priorityLabels: Record<TaskPriority, string> = {
    LOW: 'Niski',
    MEDIUM: 'Średni',
    HIGH: 'Wysoki',
};

const priorityVariants: Record<TaskPriority, BadgeVariant> = {
    HIGH: 'destructive',
    MEDIUM: 'warning',
    LOW: 'outline',
};

const statusVariants: Record<TaskStatus, BadgeVariant> = {
    DONE: 'success',
    IN_PROGRESS: 'warning',
    TODO: 'secondary',
};

type TaskItemProps = {
    task: Task;
    onDelete: (id: string) => void;
};

export function TaskItem({ task, onDelete }: TaskItemProps) {
    return (
        <div className='flex items-center gap-4 px-4 py-3'>
            <div className='min-w-0 flex-1'>
                <p className='truncate text-sm font-medium'>{task.title}</p>
                {task.description && (
                    <p className='text-muted-foreground mt-0.5 truncate text-xs'>
                        {task.description}
                    </p>
                )}
            </div>
            <div className='flex shrink-0 items-center gap-2'>
                <Badge variant={priorityVariants[task.priority]}>
                    {priorityLabels[task.priority]}
                </Badge>
                <Badge variant={statusVariants[task.status]}>
                    {statusLabels[task.status]}
                </Badge>
                <span className='text-muted-foreground w-24 text-right text-xs'>{task.dueDate}</span>
                <Button variant='ghost' size='icon-sm' onClick={() => onDelete(task.id)}>
                    <Trash2 />
                </Button>
            </div>
        </div>
    );
}
