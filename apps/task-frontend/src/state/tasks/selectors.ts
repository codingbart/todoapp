import { type RootState } from '@/lib/store';
import { tasksAdapter } from './slice';

export const { selectAll: selectAllTasks } = tasksAdapter.getSelectors<RootState>(
    state => state.tasks
);

export const selectTasksStatus = (state: RootState) => state.tasks.status;
