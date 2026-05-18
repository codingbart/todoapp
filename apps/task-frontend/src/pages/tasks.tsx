import { useState } from 'react';
import { Plus } from 'lucide-react';
import { type Task } from '@/types/task';
import { useTasks } from '@/hooks/use-tasks';
import { TaskDialog } from '@/components/task-dialog';
import { StatCard } from '@/components/stat-card';
import { TaskItem } from '@/components/task-item';
import { Button } from '@/components/ui/button';

export function TasksPage() {
    const { tasks, addTask, updateTask, deleteTask, countTasksByStatus } = useTasks();
    const [dialogOpen, setDialogOpen] = useState(false);
    const [editingTask, setEditingTask] = useState<Task | null>(null);

    function handleOpenChange(open: boolean) {
        setDialogOpen(open);
        if (!open) setEditingTask(null);
    }

    function handleEdit(task: Task) {
        setEditingTask(task);
        setDialogOpen(true);
    }

    function handleSubmit(task: Task) {
        if (editingTask) updateTask(task);
        else addTask(task);
    }

    return (
        <div className='space-y-6'>
            <div className='grid grid-cols-3 gap-4'>
                <StatCard label='Do zrobienia' value={countTasksByStatus('todo')} />
                <StatCard
                    label='W toku'
                    value={countTasksByStatus('in_progress')}
                    variant='warning'
                />
                <StatCard label='Ukończone' value={countTasksByStatus('done')} variant='success' />
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
                        <TaskItem
                            key={task.id}
                            task={task}
                            onEdit={handleEdit}
                            onDelete={deleteTask}
                        />
                    ))}
                </div>
            </div>
            <TaskDialog
                open={dialogOpen}
                onOpenChange={handleOpenChange}
                onSubmit={handleSubmit}
                task={editingTask}
            />
        </div>
    );
}
