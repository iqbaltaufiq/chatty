export const MAX_FILE_SIZE = 2 * 1000 * 1000; // 2MB

export enum MessageType {
    Chat,
    System,
    AuthEvent,
}

export interface AppState {
    user: User | null;
    room: ActiveChatroom | null;
    recent_chats: RecentChat[];
    friends: Friend[];
}

export interface User {
    uuid: string;
    username: string;
    fullname: string;
    email: string;
    photo: string;
    created_at?: Date;
    updated_at?: Date;
    deleted_at?: Date | null;
}

export interface Chat {
    id: number | null;
    username: string;
    message: string;
    room_id: string;
    created_at?: Date | null;
    updated_at?: Date | null;
    deleted_at?: Date | null;
}

export interface Message extends Chat {
    type: MessageType;
    audiences: string[];
}

export interface Room {
    id: string;
    name: string;
    is_private: boolean;
    chats?: Chat[];
    created_at?: Date;
}

export interface ActiveChatroom extends Room {
    username: string;
    fullname: string;
    photo: string;
}

export interface Friend {
    relationship_id: number;
    uuid: string;
    username: string;
    fullname: string;
    photo: string;
    room_id?: string;
}

export interface RecentChat {
    id: number;
    username: string;
    fullname: string;
    photo: string;
    room_id: string;
    message: string;
    created_at: Date;
}

export interface Credential {
    username: string;
    password: string;
}
