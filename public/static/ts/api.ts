import type { Chat, Friend, RecentChat, User } from "./types.ts";

export async function getUser(): Promise<User | null> {
    const response = await fetch("/api/user/get");
    const json = await response.json();

    if (!response.ok) {
        window.location.href = "/login";
        console.log("we cannot find your record");
        return null
    }
    const user: User = json.data?.user;
    return user;
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

export async function login(username: string, password: string): Promise<boolean> {
    let isSuccess: boolean;
    try {
        const response = await fetch("/api/login", {
            method: "POST",
            headers: [["Content-Type", "application/json"]],
            body: JSON.stringify({username, password}),
        });
        isSuccess = response.ok;
    } catch (error) {
        console.error("Login failed: "+error);
        isSuccess = false;
    }
    return isSuccess
}
