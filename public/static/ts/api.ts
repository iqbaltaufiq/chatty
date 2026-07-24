import type { RecentChat, User } from "./types.ts";

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
