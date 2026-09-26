import { deleteChat } from "../api.js";
import type { AppState, Message } from "../types.js";
import { padTime } from "../utility.js";

export function createChatBubble(state: AppState, chat: Message): Node | null {
    if (!chat.created_at) return null;

    const time = new Date(chat.created_at);
    const hh = `${padTime(time.getHours())}`; // 24h format
    const mm = `${padTime(time.getMinutes())}`;
    const timeFormat = `${hh}:${mm}`;
    const direction =
        chat.username === state.user?.username ?
        "outbound" :
        "inbound";
    const chatOption = 
        `<button popovertarget="chat-popover-${chat.id}" class="btn btn-borderless chat-option">` +
            `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-three-dots" viewBox="0 0 16 16">
              <path d="M3 9.5a1.5 1.5 0 1 1 0-3 1.5 1.5 0 0 1 0 3m5 0a1.5 1.5 0 1 1 0-3 1.5 1.5 0 0 1 0 3m5 0a1.5 1.5 0 1 1 0-3 1.5 1.5 0 0 1 0 3"/>
            </svg>` +
        `</button>` +
        `<div id="chat-popover-${chat.id}" class="chat-option-menu" popover>` +
            `<button class="btn btn-borderless chat-option-item delete-chat">Delete</button>` +
        `</div>`;
    const menuButton = direction === "outbound" ? chatOption : "";
    const el =
    `<div class="bubble ${direction}" data-id="${chat.id}">` +
        `<p class="message"></p>` +
        `<p class="time">${timeFormat}</p>` +
        menuButton +
    `</div>`;

    const doc = new DOMParser().parseFromString(el.trim(), "text/html");

    const messageEl = doc.body.querySelector(".message") as HTMLParagraphElement;
    messageEl.textContent = chat.message;
    if (direction === "inbound") return doc.body.firstChild;

    const dropdownMenu = doc.body.querySelector(".chat-option-menu") as HTMLDivElement;
    const deleteBtn = dropdownMenu.querySelector("button.delete-chat") as HTMLButtonElement;
    const bubble = doc.body.querySelector(".bubble") as HTMLDivElement;

    deleteBtn.addEventListener("click", async () => {
        if (chat.id === null) return;
        const isDeleted = await deleteChat(chat.id);
        if (!isDeleted) return;

        bubble.remove();
    });

    return doc.body.firstChild;
}
