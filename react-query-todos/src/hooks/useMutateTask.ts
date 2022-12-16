import { useMutation, useQueryClient } from "@tanstack/react-query"
import axios from "axios"
import { useAppDispatch } from "../app/hooks"
import { resetEditedTask } from "../slices/todoSlice"
import { EditTask, Task } from "../types/types"

export const useMutateTask = () => {
    const dispatch = useAppDispatch()
    const queryClient = useQueryClient()

    const createTaskMutation = useMutation(
        (task: Omit<EditTask, 'id'>) => axios.post<Task>(`${process.env.REACT_APP_REST_URL}/tasks`, task),
        {
            onSuccess: (res) => {
                const previousTodos = queryClient.getQueryData<Task[]>(['tasks'])
                if (previousTodos) {
                    queryClient.setQueryData<Task[]>(['tasks'], [...previousTodos, res.data])
                }
                dispatch(resetEditedTask())
            },
        },
    )

    const updateTaskMutation = useMutation(
        (task: EditTask) => axios.put<Task>(`${process.env.REACT_APP_REST_URL}/tasks/${task.id}`, task),
        {
            onSuccess: (res, variables) => {
                const previousTodos = queryClient.getQueryData<Task[]>(['tasks'])
                if (previousTodos) {
                    queryClient.setQueryData<Task[]>(['tasks'], previousTodos.map((task) => task.id === variables.id ? res.data : task))
                }
                dispatch(resetEditedTask())
            },
        },
    )

    const deleteTaskMutation = useMutation(
        (id: string) => axios.delete(`${process.env.REACT_APP_REST_URL}/tasks/${id}`),
        {
            onSuccess: (_, variables) => {
                const previousTodos = queryClient.getQueryData<Task[]>(['tasks'])
                if (previousTodos) {
                    queryClient.setQueryData<Task[]>(['tasks'], previousTodos.filter((task) => task.id !== variables))
                }
                dispatch(resetEditedTask())
            },
        },
    )

    return { createTaskMutation, updateTaskMutation, deleteTaskMutation }
}