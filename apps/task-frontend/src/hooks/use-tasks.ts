import { useState } from 'react';
import { type Task } from '@/types/task';

export function useTasks() {
    const [tasks, setTasks] = useState<Task[]>([]);

    const addTask = (task: Task) => setTasks(prev => [task, ...prev]);

    const deleteTask = (id: string) => setTasks(prev => prev.filter(t => t.id !== id));

    const countByStatus = (status: Task['status']) => tasks.filter(t => t.status === status).length;

    return { tasks, addTask, deleteTask, countByStatus };
}
