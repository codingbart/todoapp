import { createAsyncThunk } from '@reduxjs/toolkit';
import {
    getUsersUserIdTasks,
    postUsersUserIdTasks,
    putUsersUserIdTasksId,
    deleteUsersUserIdTasksId,
    type TaskTaskResponse,
    type TaskCreateTaskRequest,
    type TaskUpdateTaskRequest
} from '@/generated/task-api';
import { type Task } from '@/types/task';
import dayjs from 'dayjs';

function toTask(res: TaskTaskResponse): Task {
    return {
        id: res.id,
        title: res.title,
        description: res.description ?? '',
        status: res.status as Task['status'],
        priority: res.priority as Task['priority'],
        dueDate: dayjs(res.due_date),
        createdAt: dayjs(res.created_at)
    };
}

export const fetchTasks = createAsyncThunk<Task[], string>('tasks/fetchAll', async userId => {
    const data = await getUsersUserIdTasks(userId);
    return data.map(toTask);
});

export const createTask = createAsyncThunk<Task, { userId: string } & Task>(
    'tasks/create',
    async ({ userId, ...task }) => {
        const req: TaskCreateTaskRequest = {
            title: task.title,
            description: task.description,
            status: task.status,
            priority: task.priority,
            due_date: task.dueDate.format('YYYY-MM-DD')
        };

        const data = await postUsersUserIdTasks(userId, req);
        return toTask(data);
    }
);

export const updateTask = createAsyncThunk<Task, { userId: string } & Task>(
    'tasks/update',
    async ({ userId, ...task }) => {
        const req: TaskUpdateTaskRequest = {
            title: task.title,
            description: task.description,
            status: task.status,
            priority: task.priority,
            due_date: task.dueDate.format('YYYY-MM-DD')
        };

        const data = await putUsersUserIdTasksId(userId, task.id, req);
        return toTask(data);
    }
);

export const deleteTask = createAsyncThunk<string, { userId: string; id: string }>(
    'tasks/delete',
    async ({ userId, id }) => {
        await deleteUsersUserIdTasksId(userId, id);
        return id;
    }
);
