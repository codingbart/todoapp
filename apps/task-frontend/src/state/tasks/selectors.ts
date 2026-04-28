import { type RootState } from '@/lib/store';
import { tasksAdapter } from './slice';

export const { selectAll: selectAllTasks } = tasksAdapter.getSelectors(
    (state: RootState) => state.tasks,
);

export const selectTasksStatus = (state: RootState) => state.tasks.status;
