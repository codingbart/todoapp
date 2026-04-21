import { useState } from 'react';
import { Plus } from 'lucide-react';
import { useTasks } from '@/hooks/use-tasks';
import { TaskDialog } from '@/components/task-dialog';
import { StatCard } from '@/components/stat-card';
import { TaskItem } from '@/components/task-item';
import { Button } from '@/components/ui/button';

export function TasksPage() {
    const { tasks, addTask, deleteTask, countByStatus } = useTasks();
    const [dialogOpen, setDialogOpen] = useState(false);

    return (
        <div className='space-y-6'>
            <div className='grid grid-cols-3 gap-4'>
                <StatCard label='Do zrobienia' value={countByStatus('TODO')} />
                <StatCard label='W toku' value={countByStatus('IN_PROGRESS')} variant='warning' />
                <StatCard label='Ukończone' value={countByStatus('DONE')} variant='success' />
            </div>
            <div className='space-y-4'>
                <div className='flex items-center justify-between'>
                    <h1 className='text-2xl font-semibold'>Zadania</h1>
                    <Button onClick={() => setDialogOpen(true)}>
                        <Plus />
                        Nowe zadanie
                    </Button>
                </div>
                <div className='divide-y rounded-lg border'>
                    {tasks.length === 0 && (
                        <p className='text-muted-foreground px-4 py-8 text-center text-sm'>
                            Brak zadań
                        </p>
                    )}
                    {tasks.map(task => (
                        <TaskItem key={task.id} task={task} onDelete={deleteTask} />
                    ))}
                </div>
            </div>
            <TaskDialog open={dialogOpen} onOpenChange={setDialogOpen} onSubmit={addTask} />
        </div>
    );
}
