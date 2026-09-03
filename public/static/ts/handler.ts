import { send as wsSend } from "./websocket.js";
import type { AppState, Message, User } from "./types.ts";
import { getUser, logout, updateFullname, updatePhoto } from "./api.js";
import { createFriend, createUserSearch } from "./component.js";

export function registerHandler(state: AppState): void {
    handleSend(state);
    handleLogout(state);
    handleFriendListInput(state);
    handleUpdatePhoto(state);
    handleUpdateFullname(state);
}

function handleSend(state: AppState): void {
    const form = document.getElementById("chat-form") as HTMLFormElement;
    const textarea = form.querySelector("textarea") as HTMLTextAreaElement;
    form.addEventListener("submit", event => {
        event.preventDefault();
        send(state, textarea);
    });
    textarea.addEventListener("keydown", event => {
        if (event.key === "Enter" && event.shiftKey === false) {
            event.preventDefault();
            send(state, textarea);
        }
    });
}

function send(state: AppState, textarea: HTMLTextAreaElement) {
    const message = textarea.value.trim();
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
    wsSend(newChat);
    textarea.value = "";
}

function handleLogout(state: AppState): void {
    const logoutBtn = document.getElementById("logout-btn");
    logoutBtn?.addEventListener("click", async (ev) => {
        ev.preventDefault();

        const isSuccess = await logout();
        if (isSuccess) {
            state.user = null;
            state.room = null;
            state.friends = [];
            state.recent_chats = [];
            window.location.assign("/login");
        }
    });
}

function handleFriendListInput(state: AppState): void {
    const searchField =
        document.querySelector("#friend-list > input") as HTMLInputElement;
    const friendList =
        document.querySelector("#friend-list-area") as HTMLDivElement;

    searchField.addEventListener("keyup", event => {
        const searchString = searchField.value.trim();
        if (!searchString) {
            // reset friend list
            friendList.innerHTML = "";
            for (const friend of state.friends) {
                const node = createFriend(state, friend);
                if (node) friendList.appendChild(node);
            }
            return
        }
        if (searchString[0] === "@") {
            if (event.key === "Enter")
                searchUser(state, searchString, friendList);
        } else {
            filterFriends(state, searchString, friendList);
        }
        return
    });
}

async function searchUser(state: AppState, searchString: string, friendList: HTMLDivElement): Promise<void> {
    let found = false;
    const username = searchString.split("@")[1];
    if (username === state.user?.username) return;

    if (state.friends.length > 0) {
        for (let i = 0; i < state.friends.length; i++) {
            if (state.friends[i]?.username === username) {
                found = true;
                break;
            }
        }
        if (found) return;
    }
    const user = await getUser(username) satisfies User | null;
    if (!user) {
        friendList.innerHTML = "";
        return
    }
    const node = createUserSearch(user);
    friendList.innerHTML = "";
    if (node) friendList.appendChild(node);
}

function filterFriends(state: AppState, searchString: string, friendList: HTMLDivElement) {
    if (!state.friends.length) return;

    const pattern = new RegExp(searchString, "i");
    let filteredFriends = state.friends.filter(f => pattern.test(f.fullname));

    friendList.innerHTML = "";
    for (const friend of filteredFriends) {
        const node = createFriend(state, friend);
        if (node) friendList.appendChild(node);
    }
}

function handleUpdatePhoto(state: AppState): void {
    const form = document.querySelector("#profile-form") as HTMLFormElement;
    const fileInput = form.querySelector(`input[type="file"]`) as HTMLInputElement;

    fileInput.addEventListener("change", async () => {
        let file: File | null = null;
        if (fileInput.files && fileInput.files[0]) {
            file = fileInput.files[0];
        }
        const formData = new FormData();
        if (file) {
            formData.append("photo", file);
            formData.append("old_photo", state.user?.photo || "");
        }
        const success = await updatePhoto(formData);
        if (success) window.location.reload();
    });
}

function handleUpdateFullname(state: AppState): void {
    const inputField = document.querySelector(`#profile input[type="text"]`) as HTMLInputElement;
    const saveButton = document.querySelector("#profile button") as HTMLButtonElement;

    saveButton.addEventListener("click", async () => {
        const fullname = inputField.value.trim();
        const success = await updateFullname(fullname);
        if (!success) return;
        if (state.user) state.user.fullname = fullname;

        // update the user's name in user card
        const userCard = document.querySelector(".user-card .user-name");
        if (userCard) userCard.textContent = fullname;
    });
}
