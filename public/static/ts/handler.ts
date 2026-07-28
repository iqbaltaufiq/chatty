import { send as wsSend } from "./websocket.js";
import type { AppState, Message } from "./types.ts";
import { logout } from "./api.js";

export function registerHandler(state: AppState): void {
    send(state);
    handleLogout(state);
}

function send(state: AppState) {
    const form: HTMLElement | null = document.getElementById("chat-form");
    form?.addEventListener("submit", event => {
        event.preventDefault();
        const message = document.querySelector("textarea")?.value.trim();
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
    });
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
