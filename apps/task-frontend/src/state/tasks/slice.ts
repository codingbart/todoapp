import { createEntityAdapter, createSlice } from '@reduxjs/toolkit';
import { type Task } from '@/types/task';
import { type FetchStatus } from '@/types/status';
import { createTask, deleteTask, fetchTasks } from './thunks';

export const tasksAdapter = createEntityAdapter<Task>();

const initialState = tasksAdapter.getInitialState<{ status: FetchStatus }>({
    status: 'idle'
});

export const tasksSlice = createSlice({
    name: 'tasks',
    initialState,
    reducers: {},
    extraReducers: builder => {
        builder
            .addCase(fetchTasks.pending, state => {
                state.status = 'loading';
            })
            .addCase(fetchTasks.fulfilled, (state, action) => {
                state.status = 'idle';
                tasksAdapter.setAll(state, action.payload);
            })
            .addCase(fetchTasks.rejected, state => {
                state.status = 'error';
            })
            .addCase(createTask.fulfilled, (state, action) => {
                tasksAdapter.addOne(state, action.payload);
            })
            .addCase(deleteTask.fulfilled, (state, action) => {
                tasksAdapter.removeOne(state, action.payload);
            });
    }
});

export default tasksSlice.reducer;
