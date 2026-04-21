import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom';
import { TooltipProvider } from '@/components/ui/tooltip';
import { Layout } from '@/components/layout';
import { DashboardPage } from '@/pages/dashboard';
import { TasksPage } from '@/pages/tasks';

function App() {
    return (
        <BrowserRouter>
            <TooltipProvider>
                <Routes>
                    <Route element={<Layout />}>
                        <Route path='/' element={<DashboardPage />} />
                        <Route path='/tasks' element={<TasksPage />} />
                        <Route path='*' element={<Navigate to='/' replace />} />
                    </Route>
                </Routes>
            </TooltipProvider>
        </BrowserRouter>
    );
}

export default App;
