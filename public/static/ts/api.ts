import type { Chat, Friend, RecentChat, User } from "./types.ts";

export async function getUser(identifier: string | null = null): Promise<User | null> {
    try {
        let url = "/api/user/get";
        if (identifier) url = `/api/user/${identifier}/get`;

        const response = await fetch(url);
        if (!response.ok) throw new Error("user not found");

        const json = await response.json();
        const user: User = json.data?.user;
        return user
    } catch (error) {
        console.error(error);
        return null
    }
}

export async function getRecentChats(): Promise<RecentChat[] | null> {
    try {
        const response = await fetch("/api/chat/recent");
        const json = await response.json();
        if (!response.ok) {
            throw new Error(json.data?.message);
        }
        const recentChats: RecentChat[] = json.data?.chats;
        return recentChats
    } catch (error) {
        console.log(error)
        return null
    }
}

export async function getChats(room_id: string): Promise<Chat[] | null> {
    try {
        const response = await fetch(`/api/chat/${room_id}`);
        const json = await response.json();

        if (!response.ok) throw new Error("Error: "+json.data?.message);
        const chats: Chat[] =
            json.data ? json.data.chats : null;
        return chats;
    } catch (error) {
        console.error(error);
        return null;
    }
}

export async function getFriends(): Promise<Friend[] | null> {
    try {
        const response = await fetch("/api/relationship/all");
        const json = await response.json();
        if (!response.ok) throw new Error("Error: "+json.data?.message);
        const friends: Friend[] = json.data?.friends;
        return friends
    } catch (error) {
        console.error(error);
        return null
    }
}

export async function addFriend(username: string): Promise<boolean> {
    try {
        const response = await fetch(`/api/relationship/${username}`, {
            method: "POST",
        });
        if (!response.ok) throw new Error("failed to add user as a friend");
        return true
    } catch (error) {
        console.error(error);
        return false
    }
}

export async function removeFriend(relationshipID: number): Promise<boolean> {
    try {
        const response = await fetch(`/api/relationship/id/${relationshipID}`, {
            method: "DELETE"
        });
        if (!response.ok) throw new Error("failed to remove a friend");
        return true
    } catch (error) {
        console.error(error);
        return false
    }
}

export async function updatePhoto(payload: FormData): Promise<boolean> {
    try {
        const response = await fetch("/api/user/photo", {
            method: "PATCH",
            body: payload
        });
        if (!response.ok) throw new Error("failed updating your photo");
        return true
    } catch (error) {
        console.error(error);
        return false
    }
}

export async function updateFullname(fullname: string): Promise<boolean> {
    try {
        const response = await fetch(`/api/user/fullname`, {
            method: "PATCH",
            headers: [["Content-Type", "application/json"]],
            body: JSON.stringify({"fullname": fullname})
        });
        if (!response.ok) throw new Error("failed updating your name");
        return true
    } catch (error) {
        console.error(error);
        return false
    }
}

export async function deleteChat(chatID: number): Promise<boolean> {
    let isSuccess = false;
    try {
        const opt = { method: "DELETE" };
        const response = await fetch(`/api/chat/${chatID}`, opt);
        isSuccess = response.ok;

        if (!response.ok) throw new Error(`failed deleting a chat with ID: ${chatID}`)
    } catch (error) {
        console.log((error as Error).message);
    }
    return isSuccess
}

export async function login(username: string, password: string): Promise<boolean> {
    let isSuccess: boolean;
    try {
        const response = await fetch("/api/login", {
            method: "POST",
            headers: [["Content-Type", "application/json"]],
            cache: "no-store",
            body: JSON.stringify({username, password}),
        });
        isSuccess = response.ok;
    } catch (error) {
        console.error("Login failed: "+error);
        isSuccess = false;
    }
    return isSuccess
}

export async function logout(): Promise<boolean> {
    let isSuccess = false;
    try {
        const response = await fetch("/api/logout", { method: "POST", cache: "no-store" });
        isSuccess = response.ok;
    } catch (error) {
        console.error(error);
        isSuccess = false;
    }
    return isSuccess
}

export async function register(data: FormData): Promise<boolean> {
    let isSuccess = false;
    try {
        const response = await fetch("/api/register", {
            method: "POST",
            body: data
        });
        isSuccess = response.ok;
    } catch (error) {
        console.error(error);
        isSuccess = false;
    }
    return isSuccess
}
