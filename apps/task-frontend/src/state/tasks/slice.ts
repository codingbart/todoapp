import { createEntityAdapter, createSlice, isAnyOf } from '@reduxjs/toolkit';
import { type Task } from '@/types/task';
import { type FetchStatus } from '@/types/status';
import { createTask, deleteTask, fetchTasks, updateTask } from './thunks';

export const tasksAdapter = createEntityAdapter<Task>();

const initialState = tasksAdapter.getInitialState<{ status: FetchStatus }>({
    status: 'idle'
});

const thunks = [fetchTasks, createTask, updateTask, deleteTask];

export const tasksSlice = createSlice({
    name: 'tasks',
    initialState,
    reducers: {},
    extraReducers: builder => {
        builder
            .addCase(fetchTasks.fulfilled, (state, action) => {
                tasksAdapter.setAll(state, action.payload);
            })
            .addCase(createTask.fulfilled, (state, action) => {
                tasksAdapter.addOne(state, action.payload);
            })
            .addCase(updateTask.fulfilled, (state, action) => {
                tasksAdapter.upsertOne(state, action.payload);
            })
            .addCase(deleteTask.fulfilled, (state, action) => {
                tasksAdapter.removeOne(state, action.payload);
            })
            .addMatcher(isAnyOf(...thunks.map(t => t.pending)), state => {
                state.status = 'loading';
            })
            .addMatcher(isAnyOf(...thunks.map(t => t.fulfilled)), state => {
                state.status = 'idle';
            })
            .addMatcher(isAnyOf(...thunks.map(t => t.rejected)), state => {
                state.status = 'error';
            });
    }
});

export default tasksSlice.reducer;
