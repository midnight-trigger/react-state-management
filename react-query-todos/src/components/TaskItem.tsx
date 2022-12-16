import { FC, memo } from 'react'
import { useAppDispatch } from '../app/hooks'
import { useMutateTask } from '../hooks/useMutateTask'
import { Task } from '../types/types'
import { PencilAltIcon, TrashIcon } from '@heroicons/react/solid'
import { setEditedTask } from '../slices/todoSlice'

interface Props {
    task: Task
}

const TaskItem: FC<Props> = ({ task }) => {
    const dispatch = useAppDispatch()
    const { deleteTaskMutation } = useMutateTask()
    console.log('rendered TaskItem')
    if (deleteTaskMutation.isLoading) return <p>Deleting...</p>
    return (
        <li className="my-3">
            <span className="font-bold">{task.title}</span>
            <span>
                {' : '}
                {task.tag_name}
                <div className="flex float-right ml-20">
                    <PencilAltIcon
                        className="h-5 w-5 mx-1 text-blue-500 cursor-pointer"
                        onClick={() => dispatch(setEditedTask({
                            id: task.id,
                            title: task.title,
                            tag_id: task.tag_id,
                        }))}
                        />
                    <TrashIcon
                        className="h-5 w-5 mx-1 text-blue-500 cursor-pointer"
                        onClick={() => deleteTaskMutation.mutate(task.id)}
                    />
                </div>
            </span>
        </li>
    )
}

export const TaskItemMemo = memo(TaskItem)