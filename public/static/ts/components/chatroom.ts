import { getChats } from "../api.js";
import type { AppState, Chat, Friend, Message, RecentChat } from "../types.js";
import { createChatBubble } from "./chat.js";

export async function setActiveChatroom(state: AppState, chat: RecentChat | Friend): Promise<void> {
    // avoid sending request for chatroom that is already open.
    if (state.room?.id === chat.room_id) return;
    if (!chat.room_id) return;

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
    console.log({chats});
    const photo = document.querySelector(".chat-header-photo > img") as HTMLImageElement;
    photo.src = `/static/assets/${chat.photo}`;

    const imageWrapper = document.querySelector(".chat-header-photo") as HTMLDivElement;
    imageWrapper.style.display = "block";

    const fullname = document.querySelector(".chat-header-name") as HTMLParagraphElement;
    fullname.textContent = chat.fullname.trim();
    fullname.style.display = "block";
    const roomOption = document.querySelector(".room-option") as HTMLDivElement;
    roomOption.style.display = "block";

    const chatArea = document.querySelector(".chat-area") as HTMLDivElement;
    chatArea.innerHTML = "";

    for (const chat of chats) {
        const message: Message = {
            id: chat.id,
            username: chat.username,
            room_id: chat.room_id,
            message: chat.message,
            type: "chat",
            created_at: chat.created_at || null, 
            audiences: [
                state.user?.username || "",
                state.room?.username || "",
            ]
        }
        const newBubble = createChatBubble(state, message);
        if (newBubble) chatArea.appendChild(newBubble);
    }
    if (chatArea) chatArea.scrollTop = chatArea.scrollHeight;

    const chatInput = document.querySelector(".chat-input") as HTMLDivElement;
    chatInput.style.display = "block";

    const textarea = chatInput.querySelector("textarea") as HTMLTextAreaElement;
    textarea.disabled = false;

    const submitBtn = chatInput.querySelector("button") as HTMLButtonElement;
    submitBtn.disabled = false;
}
