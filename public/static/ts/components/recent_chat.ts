import { getUser } from "../api.js";
import type { AppState, Message, RecentChat, User } from "../types.js";
import { setActiveChatroom } from "./chatroom.js";

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
    const doc = new DOMParser().parseFromString(el.trim(), "text/html");

    const chatCard = doc.body.querySelector(".chat-card") as HTMLDivElement;
    chatCard.addEventListener("click", () => setActiveChatroom(state, chat));
    chatCard.title = chat.fullname;

    const cardName = doc.body.querySelector(".card-name") as HTMLParagraphElement;
    cardName.textContent = chat.fullname;

    const cardMessage = doc.body.querySelector(".card-message") as HTMLParagraphElement;
    cardMessage.textContent = chat.message;

    return doc.body.firstChild;
}

export async function updateRecentChats(state: AppState, chat: Message): Promise<void> {
    let matchIndex: number | null = null;
    const rcLength = state.recent_chats.length;

    for (let i = 0; i < rcLength; i++) {
        if (chat.room_id === state.recent_chats[i]?.room_id) {
            matchIndex = i;
            break;
        }
    }
    if (matchIndex !== null) {
        const chatFound = state.recent_chats[matchIndex];
        if (!chatFound) return;

        chatFound.message = chat.message;
        if (matchIndex > 0) {
            state.recent_chats.splice(matchIndex, 1);
            state.recent_chats.unshift(chatFound);
        }
    } else {
        let user: User | null = null;
        const targetUser =
            chat.audiences[0] !== state.user?.username ?
            chat.audiences[0] :
            chat.audiences[1];
        try {
            user = await getUser(targetUser);
            if (!user) throw new Error("username not found");
        } catch (error) {
            console.error(error);
            return
        }
        state.recent_chats.unshift({
            id: chat.id || 0,
            username: chat.username,
            fullname: user.fullname,
            room_id: chat.room_id,
            message: chat.message,
            photo: user.photo,
            created_at: chat.created_at || new Date()
        });
    }
    const chatList = document.querySelector(".chat-list");
    if (!chatList) return;

    chatList.innerHTML = "";
    state.recent_chats.map(rc => {
        const node = createRecentChat(state, rc);
        if (node) chatList.appendChild(node);
    });
}
