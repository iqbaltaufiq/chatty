import { createChatBubble } from "./component.js";
import type { AppState, Message } from "./types.ts";

export function registerHandler(state: AppState): void {
    send(state);
}

function send(state: AppState) {
    const form: HTMLElement | null = document.getElementById("chat-form");
    form?.addEventListener("submit", event => {
        event.preventDefault();
        const message = document.querySelector("textarea")?.value;

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
        const chatArea = document.querySelector(".chat-area");
        const chatBubble: Node | null = createChatBubble(state, newChat);
        if (chatBubble) chatArea?.append(chatBubble);
    });
}
