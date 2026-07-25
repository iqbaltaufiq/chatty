import { getChats } from "./api.js";
import type { AppState, Chat, Message, RecentChat } from "./types.ts";

export function createRecentChat(state: AppState, chat: RecentChat): Node | null {
    const el: string = `
    <div class="chat-card" data-room_id="${chat.room_id}">
        <div class="image-wrapper image-circle card-image">
            <img src="/static/assets/${chat.photo}" alt="">
        </div>
        <div class="card-info">
            <p class="card-name"></p>
            <p class="card-message"></p>
        </div>
    </div>
    `;
    const parser: DOMParser = new DOMParser();
    const parsedEl = parser.parseFromString(el, "text/html");

    const chatCard: HTMLDivElement | null =
        parsedEl.body.querySelector(".chat-card");
    if (chatCard) {
        chatCard.addEventListener("click", () => setActiveChatroom(state, chat));
    }

    const cardName = parsedEl.body.querySelector(".card-name");
    if (cardName) cardName.textContent = chat.fullname;

    const cardMessage = parsedEl.body.querySelector(".card-message");
    if (cardMessage) cardMessage.textContent = chat.message;

    return parsedEl.body.firstChild;
}

export function createChatBubble(state: AppState, chat: Message): Node | null {
    let time: Date;
    let timeFormat: string = "";
    if (chat.created_at) {
        time = new Date(chat.created_at);
        timeFormat = `${time?.getHours()}:${time?.getMinutes()}`;
    }
    const direction =
        chat.username === state.user?.username ? "outbound" : "inbound";
    const el = `
    <div class="bubble ${direction}">
        <p class="message"></p>
        <p class="time">${timeFormat}</p>
    </div>
    `;
    const parser: DOMParser = new DOMParser();
    const parsedEl: Document = parser.parseFromString(el, "text/html");

    const messageEl: HTMLElement | null =
        parsedEl.body.querySelector(".message");
    if (messageEl) messageEl.textContent = chat.message;

    return parsedEl.body.firstChild;
}

async function setActiveChatroom(state: AppState, chat: RecentChat): Promise<void> {
    // avoid sending request for chatroom that is already open.
    if (state.room?.id === chat.room_id) return;

    let chats: Chat[] = [];
    const resp = await getChats(chat.room_id);
    chats = resp ? resp : [];

    state.room = {
        id: chat.room_id,
        username: chat.username,
        fullname: chat.fullname,
        photo: chat.photo,
        is_private: true,
        name: chat.fullname,
        chats: chats
    }
    const photo: HTMLImageElement | null =
        document.querySelector(".chat-header-photo > img");
    if (photo) photo.src = `/static/assets/${chat.photo}`;

    const fullname: HTMLParagraphElement | null =
        document.querySelector(".chat-header-name");
    if (fullname) fullname.textContent = chat.fullname.trim();

    const chatArea = document.querySelector(".chat-area");
    if (chatArea) chatArea.innerHTML = "";

    for (const chat of chats) {
        const message: Message = {
            id: chat.id,
            username: chat.username,
            room_id: chat.room_id,
            message: chat.message,
            type: "chat",
            audiences: [
                state.user?.username || "",
                state.room?.username || "",
            ]
        }
        const newBubble = createChatBubble(state, message)
        if (newBubble) chatArea?.appendChild(newBubble);
    }
}
