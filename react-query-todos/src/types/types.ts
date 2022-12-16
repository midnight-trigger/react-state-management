export interface Task {
    id: string
    title: string
    tag_id: number
    tag_name: string
    created_at: string
    updated_at: string
}

export interface EditTask {
    id: string
    title: string
    tag_id: number
}

export interface Tag {
    id: number
    name: string
}