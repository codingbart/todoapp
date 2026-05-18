import { Pencil, Trash2 } from 'lucide-react';
import { type VariantProps } from 'class-variance-authority';
import { type Task, type TaskPriority, type TaskStatus } from '@/types/task';
import { Badge, badgeVariants } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';

type BadgeVariant = NonNullable<VariantProps<typeof badgeVariants>['variant']>;

const statusLabels: Record<TaskStatus, string> = {
    todo: 'Do zrobienia',
    in_progress: 'W toku',
    done: 'Ukończone',
};

const priorityLabels: Record<TaskPriority, string> = {
    low: 'Niski',
    medium: 'Średni',
    high: 'Wysoki',
};

const priorityVariants: Record<TaskPriority, BadgeVariant> = {
    high: 'destructive',
    medium: 'warning',
    low: 'outline',
};

const statusVariants: Record<TaskStatus, BadgeVariant> = {
    done: 'success',
    in_progress: 'warning',
    todo: 'secondary',
};

type TaskItemProps = {
    task: Task;
    onEdit: (task: Task) => void;
    onDelete: (id: string) => void;
};

export function TaskItem({ task, onEdit, onDelete }: TaskItemProps) {
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
                <span className='text-muted-foreground w-24 text-right text-xs'>
                    {task.dueDate || '—'}
                </span>
                <Button variant='ghost' size='icon-sm' onClick={() => onEdit(task)}>
                    <Pencil />
                </Button>
                <Button variant='ghost' size='icon-sm' onClick={() => onDelete(task.id)}>
                    <Trash2 />
                </Button>
            </div>
        </div>
    );
}
