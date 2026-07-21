import type { AppState, Message, RecentChat } from "./types.ts";

export function createRecentChat(chat: RecentChat): Node | null {
    const el: string = `
    <div class="chat-card" data-roomID="${chat.room_id}">
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

    const cardName = parsedEl.body.querySelector(".card-name");
    if (cardName) cardName.textContent = chat.fullname;

    const cardMessage = parsedEl.body.querySelector(".card-message");
    if (cardMessage) cardMessage.textContent = chat.message;

    return parsedEl.body.firstChild;
}

export function createChatBubble(state: AppState, chat: Message): Node | null {
    const direction =
        chat.username === state.user?.username ? "outbound" : "inbound";
    const time =
        `${chat.created_at?.getHours()}:${chat.created_at?.getMinutes()}`;
    const el = `
    <div class="bubble ${direction}">
        <p class="message"></p>
        <p class="time">${time}</p>
    </div>
    `;
    const parser: DOMParser = new DOMParser();
    const parsedEl: Document = parser.parseFromString(el, "text/html");

    const messageEl: HTMLElement | null =
        parsedEl.body.querySelector(".message");
    if (messageEl) messageEl.textContent = chat.message;

    return parsedEl.body.firstChild;
}
