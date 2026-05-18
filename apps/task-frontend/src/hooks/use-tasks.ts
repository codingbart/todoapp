import { useEffect } from 'react';
import { type Task } from '@/types/task';
import { useAuth } from '@/hooks/use-auth';
import { useAppDispatch } from '@/hooks/use-app-dispatch';
import { useAppSelector } from '@/hooks/use-app-selector';
import { selectAllTasks, selectTasksStatus } from '@/state/tasks/selectors';
import { createTask, deleteTask, fetchTasks, updateTask } from '@/state/tasks/thunks';

export function useTasks() {
    const { user } = useAuth();
    const tasks = useAppSelector(selectAllTasks);
    const status = useAppSelector(selectTasksStatus);
    const dispatch = useAppDispatch();

    useEffect(() => {
        dispatch(fetchTasks(user.id));
    }, [dispatch, user.id]);

    const countTasksByStatus = (status: Task['status']) =>
        tasks.filter(t => t.status === status).length;

    return {
        tasks,
        status,
        addTask: (task: Task) => dispatch(createTask({ userId: user.id, ...task })),
        updateTask: (task: Task) => dispatch(updateTask({ userId: user.id, ...task })),
        deleteTask: (id: string) => dispatch(deleteTask({ userId: user.id, id })),
        countTasksByStatus
    };
}
