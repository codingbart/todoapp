import { createAsyncThunk } from '@reduxjs/toolkit';
import { type Task } from '@/types/task';

async function dummy_fetch() {
    return new Response('[]');
}

export const fetchTasks = createAsyncThunk<Task[], string>('tasks/fetchAll', async (_userId) => {
    const res = await dummy_fetch(); // TODO: GET /api/users/:userId/tasks

    if (!res.ok) {
        throw new Error('Failed to fetch tasks');
    }

    return res.json();
});

export const createTask = createAsyncThunk<Task, Task>('tasks/create', async (_task) => {
    const res = await dummy_fetch(); // TODO: POST /api/tasks with _task body

    if (!res.ok) {
        throw new Error('Failed to create task');
    }

    return _task; // TODO: return res.json()
});

export const deleteTask = createAsyncThunk<string, string>('tasks/delete', async (_id) => {
    const res = await dummy_fetch(); // TODO: DELETE /api/tasks/:_id

    if (!res.ok) {
        throw new Error('Failed to delete task');
    }

    return _id;
});
