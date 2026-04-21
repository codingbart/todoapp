import { type Task, type TaskPriority, type TaskStatus } from '@/types/task';
import { useTaskForm } from '@/hooks/use-task-form';
import { Button } from '@/components/ui/button';
import { Dialog, DialogContent, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { FormField } from '@/components/ui/form-field';
import { Input } from '@/components/ui/input';
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue
} from '@/components/ui/select';
import { Textarea } from '@/components/ui/textarea';

const statusLabels: Record<TaskStatus, string> = {
    TODO: 'Do zrobienia',
    IN_PROGRESS: 'W toku',
    DONE: 'Ukończone'
};

const priorityLabels: Record<TaskPriority, string> = {
    LOW: 'Niski',
    MEDIUM: 'Średni',
    HIGH: 'Wysoki'
};

type TaskDialogProps = {
    open: boolean;
    onOpenChange: (open: boolean) => void;
    onSubmit: (task: Task) => void;
};

export function TaskDialog({ open, onOpenChange, onSubmit }: TaskDialogProps) {
    const { form, reset, setTitle, setDescription, setStatus, setPriority, setDueDate } =
        useTaskForm();

    function handleOpenChange(open: boolean) {
        if (!open) reset();
        onOpenChange(open);
    }

    function handleSubmit(e: React.SyntheticEvent<HTMLFormElement>) {
        e.preventDefault();
        onSubmit({
            id: String(Date.now()),
            ...form,
            createdAt: new Date().toISOString().split('T')[0]
        });
        onOpenChange(false);
    }

    return (
        <Dialog open={open} onOpenChange={handleOpenChange}>
            <DialogContent>
                <DialogHeader>
                    <DialogTitle>Nowe zadanie</DialogTitle>
                </DialogHeader>
                <form onSubmit={handleSubmit} className='space-y-4'>
                    <FormField label='Tytuł'>
                        <Input
                            required
                            value={form.title}
                            onChange={e => setTitle(e.target.value)}
                            placeholder='Tytuł zadania'
                        />
                    </FormField>
                    <FormField label='Opis'>
                        <Textarea
                            value={form.description}
                            onChange={e => setDescription(e.target.value)}
                            placeholder='Opis zadania'
                            rows={3}
                        />
                    </FormField>
                    <FormField label='Status'>
                        <Select value={form.status} onValueChange={v => setStatus(v as TaskStatus)}>
                            <SelectTrigger className='w-full'>
                                <SelectValue>{statusLabels[form.status]}</SelectValue>
                            </SelectTrigger>
                            <SelectContent>
                                <SelectItem value='TODO'>Do zrobienia</SelectItem>
                                <SelectItem value='IN_PROGRESS'>W toku</SelectItem>
                                <SelectItem value='DONE'>Ukończone</SelectItem>
                            </SelectContent>
                        </Select>
                    </FormField>
                    <FormField label='Priorytet'>
                        <Select
                            value={form.priority}
                            onValueChange={v => setPriority(v as TaskPriority)}
                        >
                            <SelectTrigger className='w-full'>
                                <SelectValue>{priorityLabels[form.priority]}</SelectValue>
                            </SelectTrigger>
                            <SelectContent>
                                <SelectItem value='LOW'>Niski</SelectItem>
                                <SelectItem value='MEDIUM'>Średni</SelectItem>
                                <SelectItem value='HIGH'>Wysoki</SelectItem>
                            </SelectContent>
                        </Select>
                    </FormField>
                    <FormField label='Termin'>
                        <Input
                            type='date'
                            value={form.dueDate}
                            onChange={e => setDueDate(e.target.value)}
                        />
                    </FormField>
                    <div className='flex gap-2 pt-2'>
                        <Button type='submit' className='flex-1'>
                            Utwórz zadanie
                        </Button>
                        <Button
                            type='button'
                            variant='outline'
                            onClick={() => handleOpenChange(false)}
                        >
                            Anuluj
                        </Button>
                    </div>
                </form>
            </DialogContent>
        </Dialog>
    );
}
