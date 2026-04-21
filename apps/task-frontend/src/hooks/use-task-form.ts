import { useState } from 'react';
import { type TaskPriority, type TaskStatus } from '@/types/task';

export type TaskFormData = {
    title: string;
    description: string;
    status: TaskStatus;
    priority: TaskPriority;
    dueDate: string;
};

const emptyForm: TaskFormData = {
    title: '',
    description: '',
    status: 'TODO',
    priority: 'MEDIUM',
    dueDate: ''
};

export function useTaskForm() {
    const [form, setForm] = useState<TaskFormData>(emptyForm);

    function reset() {
        setForm(emptyForm);
    }

    const setTitle = (title: string) => setForm(f => ({ ...f, title }));
    const setDescription = (description: string) => setForm(f => ({ ...f, description }));
    const setStatus = (status: TaskStatus) => setForm(f => ({ ...f, status }));
    const setPriority = (priority: TaskPriority) => setForm(f => ({ ...f, priority }));
    const setDueDate = (dueDate: string) => setForm(f => ({ ...f, dueDate }));

    return { form, reset, setTitle, setDescription, setStatus, setPriority, setDueDate };
}
