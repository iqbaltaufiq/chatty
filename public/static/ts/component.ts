import type { AppState, Message } from "./types.ts";

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
