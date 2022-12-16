import { FC, FormEvent, memo } from 'react'
import { useAppDispatch, useAppSelector } from '../app/hooks'
import { useMutateTag } from '../hooks/useMutateTag'
import { selectTag, setEditedTag } from '../slices/todoSlice'

const TagEdit: FC = () => {
    const dispatch = useAppDispatch()
    const editedTag = useAppSelector(selectTag)
    const { createTagMutation, updateTagMutation } = useMutateTag()

    const submitHandler = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        if (editedTag.id === 0) createTagMutation.mutate(editedTag)
        else updateTagMutation.mutate(editedTag)
    }

    console.log('rendered TagEdit')
    if (createTagMutation.isLoading) return <span>Creating...</span>
    if (updateTagMutation.isLoading) return <span>Updating...</span>

    return (
        <div>
            <form onSubmit={submitHandler}>
                <input
                    type="text"
                    className="mb-3 px-3 py-2 border border-gray-300"
                    placeholder="new tag ?"
                    value={editedTag.name}
                    onChange={(e) =>  dispatch(setEditedTag({
                        ...editedTag,
                        name: e.target.value,
                    }))}
                />
                <button
                    className="disabled:opacity-40 my-3 mx-3 py-2 px-3 text-white bg-indigo-600 hover:bg-indigo-700 rounded"
                    disabled={!editedTag.name}
                >
                    {editedTag.id ? 'Update' : 'Create'}
                </button>
            </form>
        </div>
    )
}

export const TagEditMemo = memo(TagEdit)