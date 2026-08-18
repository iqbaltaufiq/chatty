import { addFriend, getChats, getUser, removeFriend } from "./api.js";
import type { AppState, Chat, Friend, Message, RecentChat, User } from "./types.ts";
import { padTime } from "./utility.js";

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

export function createChatBubble(state: AppState, chat: Message): Node | null {
    let time: Date;
    let timeFormat: string = "";
    if (chat.created_at) {
        time = new Date(chat.created_at);
        if (time) timeFormat =
                `${padTime(time.getHours())}:${padTime(time.getMinutes())}`;
    }
    const direction =
        chat.username === state.user?.username ? "outbound" : "inbound";
    const el =
    `<div class="bubble ${direction}">` +
        `<p class="message"></p>` +
        `<p class="time">${timeFormat}</p>` +
    `</div>`;
    const parser: DOMParser = new DOMParser();
    const parsedEl: Document = parser.parseFromString(el, "text/html");

    const messageEl: HTMLElement | null =
        parsedEl.body.querySelector(".message");
    if (messageEl) messageEl.textContent = chat.message;

    return parsedEl.body.firstChild;
}

export function createFriend(state: AppState, friend: Friend): Node | null {
    const el = `
        <div class="friend-card" data-id="${friend.relationship_id}">
            <div class="image-wrapper image-circle friend-photo">
                <img src="/static/assets/${friend.photo}" alt="">
            </div>
            <p class="friend-name"></p>
            <button 
                type="button"
                data-id="${friend.relationship_id}"
                class="btn btn-borderless action-btn remove-btn"
            >&times;</button>
        </div>
    `;
    const parser: DOMParser = new DOMParser();
    const doc: Document = parser.parseFromString(el, "text/html");

    const friendCard = doc.body.querySelector(".friend-card");
    const actionBtn = doc.body.querySelector(".action-btn");

    friendCard?.addEventListener("click", (e) => {
        if (e.target === actionBtn) return;
        setActiveChatroom(state, friend);
        document.getElementById("friend-list")?.togglePopover();
    });

    const friendName: HTMLParagraphElement | null =
        doc.body.querySelector(".friend-name");
    if (friendName) friendName.textContent = friend.fullname;

    actionBtn?.addEventListener("click", async (_) => {
        const success = await removeFriend(friend.relationship_id);
        if (success) return window.location.reload();
    });

    return doc.body.firstChild;
}

export function createUserSearch(user: User): Node | null {
    const el = `
        <div class="friend-card">
            <div class="image-wrapper image-circle friend-photo">
                <img src="/static/assets/${user.photo}" alt="">
            </div>
            <p class="friend-name"></p>
            <button 
                type="button"
                data-username="${user.username}"
                class="btn btn-borderless action-btn add-btn"
            >&plus;</button>
        </div>
    `;
    const parser = new DOMParser();
    const doc = parser.parseFromString(el.trim(), "text/html");
    const friendName = doc.body.querySelector("p");
    if (friendName) friendName.textContent = user.fullname;

    const addBtn = doc.body.querySelector("button");
    addBtn?.addEventListener("click", async () => {
        try {
            const success = addFriend(user.username);
            if (!success) throw new Error("failed adding user as friend");
            return window.location.reload();
        } catch (error) {
            console.error(error);
            return
        }
    });
    return doc.body.firstChild;
}

async function setActiveChatroom(state: AppState, chat: RecentChat | Friend): Promise<void> {
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
    const photo: HTMLImageElement | null =
        document.querySelector(".chat-header > img");
    if (photo) photo.src = `/static/assets/${chat.photo}`;

    const imageWrapper: HTMLElement | null =
        document.querySelector(".chat-header-photo");
    if (imageWrapper) imageWrapper.style.display = "block";

    const fullname: HTMLElement | null =
        document.querySelector(".chat-header-name");
    if (fullname) {
        fullname.textContent = chat.fullname.trim();
        fullname.style.display = "block";
    }
    const chatArea = document.querySelector(".chat-area");
    if (chatArea) chatArea.innerHTML = "";

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
        const newBubble = createChatBubble(state, message)
        if (newBubble) chatArea?.appendChild(newBubble);
    }
    if (chatArea) chatArea.scrollTop = chatArea.scrollHeight;

    const chatInput: HTMLElement | null =
        document.querySelector(".chat-input");
    if (!chatInput) return;
    chatInput.style.display = "block";

    const textarea = chatInput.querySelector("textarea");
    if (textarea) textarea.disabled = false;

    const submitBtn = chatInput.querySelector("button");
    if (submitBtn) submitBtn.disabled = false;
}
