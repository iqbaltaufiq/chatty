import type { AppState, Friend } from "./types.js";
import { getFriends, getRecentChats, getUser } from "./api.js";
import { registerHandler } from "./handler.js";
import { toggleConnection } from "./websocket.js";
import { createFriend, createRecentChat } from "./component.js";

export async function initApp(state: AppState): Promise<void> {
    await initUser(state);
    await initRecentChats(state);
    await initFriendList(state);
    registerHandler(state);
    toggleConnection(state);
}

export async function initUser(state: AppState): Promise<void> {
    try {
        const user = await getUser();
        if (!user) throw new Error("failed initializing user");
        state.user = user;
    } catch (error) {
        console.error(error);
        return
    }

    const photoExists = `/static/assets/${state.user?.photo}`;
    const photoAbsent = `/static/assets/default.png`;
    const photo: HTMLImageElement | null = document.querySelector(".user-photo > img");
    if (photo) photo.src = photoExists || photoAbsent;

    const userName: HTMLParagraphElement | null = document.querySelector(".user-name");
    if (userName) userName.textContent = state.user?.fullname || "unknown";
}

export async function initRecentChats(state: AppState): Promise<void> {
    try {
        const chats = await getRecentChats();
        if (!chats) return;
        state.recent_chats = chats;
    } catch (error) {
        console.error(error);
        return
    }

    const chatList: HTMLDivElement | null = document.querySelector(".chat-list");
    for (const chat of state.recent_chats) {
        if (!chatList) break;
        const chatNode = createRecentChat(state, chat);
        if (chatNode) chatList.appendChild(chatNode);
    }
}

async function initFriendList(state: AppState): Promise<void> {
    try {
        const friends = await getFriends();
        state.friends = friends ? friends : [];
    } catch (error) {
        console.error(error)
        return
    }
    
    const friendList = document.getElementById("friend-list");
    for (const friend of state.friends) {
        const newFriend = createFriend(state, friend);
        if (newFriend) friendList?.appendChild(newFriend);
    }
}
