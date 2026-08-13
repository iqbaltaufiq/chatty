import { send as wsSend } from "./websocket.js";
import type { AppState, Message } from "./types.ts";
import { logout } from "./api.js";
import { createFriend } from "./component.js";

export function registerHandler(state: AppState): void {
    handleSend(state);
    handleLogout(state);
    handleFriendListInput(state);
}

function handleSend(state: AppState): void {
    const form = document.getElementById("chat-form");
    if (!form) return;
    const textarea = form.querySelector("textarea");
    if (!textarea) return;
    form.addEventListener("submit", event => {
        event.preventDefault();
        send(state, textarea);
    });
    textarea.addEventListener("keydown", event => {
        if (event.key === "Enter" && event.shiftKey === false) {
            event.preventDefault();
            send(state, textarea);
        }
    });
}

function send(state: AppState, textarea: HTMLTextAreaElement) {
    const message = textarea.value.trim();
    const newChat: Message = {
        id: null,
        username: state.user?.username || "",
        room_id: state.room?.id || "",
        type: "chat",
        message: message || "",
        created_at: new Date(),
        audiences: [
            state.user?.username || "",
            state.room?.username || "",
        ]
    }
    wsSend(newChat);
    textarea.value = "";
}

function handleLogout(state: AppState): void {
    const logoutBtn = document.getElementById("logout-btn");
    logoutBtn?.addEventListener("click", async (ev) => {
        ev.preventDefault();

        let isSuccess = false;
        try {
            isSuccess = await logout();
        } catch (error) {
            console.error(error);
            return
        }
        if (isSuccess) {
            state.user = null;
            state.room = null;
            state.friends = [];
            state.recent_chats = [];
            window.location.href = "/login";
        }
    });
}

function handleFriendListInput(state: AppState): void {
    if (state.friends.length < 1) return;
    const searchField: HTMLInputElement | null =
        document.querySelector("#friend-list > input");
    const friendList = document.querySelector("#friend-list-area");
    if (!friendList) return;
    searchField?.addEventListener("keyup", event => {
        const searchString = searchField?.value;
        if (!searchString) {
            friendList.innerHTML = "";
            for (const friend of state.friends) {
                const node = createFriend(state, friend);
                if (node) friendList.appendChild(node);
            }
            return
        } else {
            const pattern = new RegExp(searchString, "i");
            let filteredFriends = state.friends.filter(f => pattern.test(f.fullname));
            friendList.innerHTML = "";
            for (const friend of filteredFriends) {
                const node = createFriend(state, friend);
                if (node) friendList.appendChild(node);
            }
            return
        }
    });
}
