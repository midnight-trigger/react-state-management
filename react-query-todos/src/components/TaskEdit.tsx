import { FC, FormEvent, memo } from 'react'
import { useAppDispatch, useAppSelector } from '../app/hooks'
import { useMutateTask } from '../hooks/useMutateTask'
import { useQueryTags } from '../hooks/useQueryTags'
import { selectTask, setEditedTask } from '../slices/todoSlice'

const TaskEdit: FC = () => {
    const dispatch = useAppDispatch()
    const editedTask = useAppSelector(selectTask)
    const { data, status } = useQueryTags()
    const { createTaskMutation, updateTaskMutation } = useMutateTask()

    const submitHandler = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        if (!editedTask.id) createTaskMutation.mutate(editedTask)
        else updateTaskMutation.mutate(editedTask)
    }

    const tagOptions = data?.map((tag) => (
        <option key={tag.id} value={tag.id}>
            {tag.name}
        </option>
    ))
    console.log('rendered EditTask')
    if (status === 'loading') return <div>{"Loading..."}</div>
    if (status === 'error') return <div>{"Error"}</div>
    if (updateTaskMutation.isLoading) return <span>Updating...</span>
    if (createTaskMutation.isLoading) return <span>Creating...</span>
    return (
        <div>
            <form onSubmit={submitHandler}>
                <input
                    className="mb-3 px-3 py-2 border border-gray-300"
                    placeholder="new task ?"
                    type="text"
                    onChange={(e) => dispatch(setEditedTask({
                        ...editedTask,
                        title: e.target.value
                    }))}
                    value={editedTask.title}
                />
                <button
                    className="disabled:opacity-40 my-3 mx-3 py-2 px-3 text-white bg-indigo-600 hover:bg-indigo-700 rounded"
                    disabled={!editedTask.title || !editedTask.tag_id}
                >
                    {editedTask.id ? 'Update' : 'Create'}
                </button>
            </form>
            <select
                className="mb-3 px-3 py-2 border border-gray-200"
                value={editedTask.tag_id}
                onChange={(e) => dispatch(setEditedTask({
                    ...editedTask,
                    tag_id: Number(e.target.value),
                }))}
            >
                <option value={0}>Tag</option>
                {tagOptions}
            </select>
        </div>
    )
}

export const TaskEditMemo = memo(TaskEdit)