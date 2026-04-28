import { useEffect } from 'react';
import { type Task } from '@/types/task';
import { useAppDispatch } from '@/hooks/use-app-dispatch';
import { useAppSelector } from '@/hooks/use-app-selector';
import { selectAllTasks, selectTasksStatus } from '@/state/tasks/selectors';
import { createTask, deleteTask, fetchTasks } from '@/state/tasks/thunks';

export function useTasks() {
    const userId = '123'; // TODO: get user id - useAuth() or etc.
    const tasks = useAppSelector(selectAllTasks);
    const status = useAppSelector(selectTasksStatus);
    const dispatch = useAppDispatch();

    useEffect(() => {
        dispatch(fetchTasks(userId));
    }, [dispatch, userId]);

    const countTasksByStatus = (status: Task['status']) =>
        tasks.filter(t => t.status === status).length;

    return {
        tasks,
        status,
        addTask: (task: Task) => dispatch(createTask(task)),
        deleteTask: (id: string) => dispatch(deleteTask(id)),
        countTasksByStatus
    };
}
