import { createChatBubble, updateRecentChats } from "./component.js";
import type { AppState, Message } from "./types.ts"

const wsUrl: string = "ws://localhost:8000/websocket";
let ws: WebSocket | null = null;

export function toggleConnection(state: AppState): void {
    if (!ws) {
        try {
            ws = new WebSocket(wsUrl);
            ws.onopen = _ => console.info("connected to the ws server.");
            ws.onmessage = message => receiveMessage(state, message);
            ws.onclose = _ => console.info("disconnected from the ws server.");
            ws.onerror = error => console.error(error);
        } catch (error) {
            console.error(error);
        }
    } else {
        if (ws.readyState === ws.OPEN) {
            ws.close();
            ws = null;
        }
    }
}

function receiveMessage(state: AppState, message: MessageEvent): void {
    const chat: Message = JSON.parse(message.data);
    const isTypeChat = chat.type === "chat";
    const isTheChatOpened = chat.room_id === state.room?.id;
    console.log({chat});

    // append new chat bubble into chat field.
    if (isTypeChat && isTheChatOpened) {
        const newBubble = createChatBubble(state, chat);
        if (newBubble) {
            const chatArea = document.querySelector(".chat-area");
            if (chatArea) {
                chatArea.append(newBubble);
                chatArea.scrollTop = chatArea.scrollHeight;
            }
        }
    }
    updateRecentChats(state, chat);
}

export function send(chat: Message): void {
    if (ws && ws.readyState === ws.OPEN) {
        ws.send(JSON.stringify(chat));
    }
}
