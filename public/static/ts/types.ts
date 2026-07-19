export interface AppState {
    user: User;
    recent_chats: RecentChat[];
    room: Room;
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
    created_at: Date | null;
    updated_at?: Date | null;
    deleted_at?: Date | null;
}

export interface Message extends Chat {
    audiences: string[];
}

export interface Room {
    id: string;
    name: string;
    is_private: boolean;
    chats?: Chat[];
    created_at?: Date;
}

export interface ActiveChatroom {
    username: string;
    fullname: string;
    photo: string;
    room_id: string;
    chats: Chat[];
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
