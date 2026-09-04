import { addFriend, deleteChat, getChats, getUser, removeFriend } from "./api.js";
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
        `<div class="chat-option">` +
            `<a tabindex="0" class="">` +
                `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-three-dots" viewBox="0 0 16 16">
                  <path d="M3 9.5a1.5 1.5 0 1 1 0-3 1.5 1.5 0 0 1 0 3m5 0a1.5 1.5 0 1 1 0-3 1.5 1.5 0 0 1 0 3m5 0a1.5 1.5 0 1 1 0-3 1.5 1.5 0 0 1 0 3"/>
                </svg>` +
            `</a>` +
            `<div class="chat-option-menu">` +
                `<button class="btn btn-borderless chat-option-item delete-chat">Delete</button>` +
            `</div>` +
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

    const tripleDot = doc.body.querySelector(".chat-option a") as HTMLAnchorElement;
    const dropdownMenu = doc.body.querySelector(".chat-option-menu") as HTMLDivElement;
    const deleteBtn = dropdownMenu.querySelector("button.delete-chat") as HTMLButtonElement;
    const bubble = doc.body.querySelector(".bubble") as HTMLDivElement;

    tripleDot.addEventListener("click", () => {
        const isMenuVisible = dropdownMenu.style.display === "block";
        dropdownMenu.style.display = isMenuVisible ? "none" : "block";
    });
    deleteBtn.addEventListener("click", async () => {
        if (chat.id === null) return;
        const isDeleted = await deleteChat(chat.id);
        if (!isDeleted) return;

        bubble.remove();
    });

    return doc.body.firstChild;
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
    const doc = new DOMParser().parseFromString(el.trim(), "text/html");

    const friendCard = doc.body.querySelector(".friend-card") as HTMLDivElement;
    const actionBtn = doc.body.querySelector(".action-btn") as HTMLButtonElement;

    friendCard.addEventListener("click", (e) => {
        if (e.target === actionBtn) return;
        setActiveChatroom(state, friend);
        document.getElementById("friend-list")?.togglePopover();
    });

    const friendName = doc.body.querySelector(".friend-name") as HTMLParagraphElement;
    friendName.textContent = friend.fullname;

    actionBtn.addEventListener("click", async () => {
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
    const doc = new DOMParser().parseFromString(el.trim(), "text/html");
    const friendName = doc.body.querySelector("p") as HTMLParagraphElement;
    friendName.textContent = user.fullname;

    const addBtn = doc.body.querySelector("button") as HTMLButtonElement;
    addBtn.addEventListener("click", async () => {
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
    console.log({chats});
    const photo = document.querySelector(".chat-header-photo > img") as HTMLImageElement;
    photo.src = `/static/assets/${chat.photo}`;

    const imageWrapper = document.querySelector(".chat-header-photo") as HTMLDivElement;
    imageWrapper.style.display = "block";

    const fullname = document.querySelector(".chat-header-name") as HTMLParagraphElement;
    fullname.textContent = chat.fullname.trim();
    fullname.style.display = "block";

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

export function createProfile(state: AppState): Node | null {
    if (!state.user) return null;
    const el = `
    <div id="profile" popover>
        <table>
            <tr>
                <td>UserID</td>
                <td>${state.user.uuid}</td>
            </tr>
            <tr>
                <td>Fullname</td>
                <td>
                    <div class="input-group">
                        <input type="text" name="fullname" maxlength="32">
                        <button type="button" class="btn btn-success">OK</button>
                    </div>
                </td>
            </tr>
            <tr>
                <td>Username</td>
                <td>${state.user.username}</td>
            </tr>
            <tr>
                <td>Email</td>
                <td>${state.user.email}</td>
            </tr>
            <tr>
                <td>Change Photo</td>
                <td>
                    <form id="profile-form" method="post" enctype="multipart/form-data">
                        <input type="file" name="update_photo">
                    </form>
                </td>
            </tr>
        </table>
    </div>
    `
    const doc = new DOMParser().parseFromString(el.trim(), "text/html");
    const inputText = doc.body.querySelector(`input[type="text"]`) as HTMLInputElement;
    inputText.value = state.user.fullname;
    return doc.body.firstChild;
}
