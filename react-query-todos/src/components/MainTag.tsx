import { ChevronDoubleLeftIcon } from '@heroicons/react/solid'
import { FC } from 'react'
import { useNavigate } from 'react-router-dom'
import { TagEditMemo } from './TagEdit'
import { TagListMemo } from './TagList'

export const MainTag: FC = () => {
    const navigate = useNavigate()
    console.log('rendered MainTag')
    return (
        <>
            <p className="mb-10 text-xl font-bold">Tags</p>
            <div className="grid grid-cols-2 gap-40">
                <TagListMemo />
                <TagEditMemo />
            </div>
            <ChevronDoubleLeftIcon
                className="h-5 w-5 mt-2 text-blue-500 cursor-pointer"
                onClick={() => navigate('/')}
            />
        </>
    )
}
