import { Navigate, Route, Routes } from 'react-router-dom';
import { Layout } from '@/components/layout';
import { TasksPage } from '@/pages/tasks';

export function Router() {
    return (
        <Routes>
            <Route element={<Layout />}>
                <Route path='/' element={<TasksPage />} />
                <Route path='*' element={<Navigate to='/' replace />} />
            </Route>
        </Routes>
    );
}
