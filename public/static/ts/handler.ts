import { send as wsSend } from "./websocket.js";
import type { AppState, Message } from "./types.ts";
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
    const form = document.getElementById("chat-form");
    if (!form) return;
    const textarea = form.querySelector("textarea");
    if (!textarea) return;
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

        let isSuccess = false;
        try {
            isSuccess = await logout();
        } catch (error) {
            console.error(error);
            return
        }
        if (isSuccess) {
            state.user = null;
            state.room = null;
            state.friends = [];
            state.recent_chats = [];
            window.location.href = "/login";
        }
    });
}

function handleFriendListInput(state: AppState): void {
    if (state.friends.length < 1) return;
    const searchField: HTMLInputElement | null =
        document.querySelector("#friend-list > input");
    const friendList = document.querySelector("#friend-list-area");
    if (!friendList) return;
    searchField?.addEventListener("keyup", event => {
        const searchString = searchField?.value;
        if (!searchString) {
            friendList.innerHTML = "";
            for (const friend of state.friends) {
                const node = createFriend(state, friend);
                if (node) friendList.appendChild(node);
            }
            return
        } else {
            if (searchString[0] === "@") {
                if (event.key === "Enter") {
                    let found = false;
                    const username = searchString.split("@")[1];
                    if (username === state.user?.username) return;
                    state.friends.map(f => {
                        if (f.username === username) {
                            found = true;
                            return
                        }
                    });
                    if (found) return;

                    getUser(username)
                    .then(user => {
                        if (!user) return;
                        const node = createUserSearch(user);
                        if (!node) return;
                        friendList.innerHTML = "";
                        friendList.appendChild(node);
                    })
                    .catch(err => console.error(err));
                }
            } else {
                const pattern = new RegExp(searchString, "i");
                let filteredFriends = state.friends.filter(f => pattern.test(f.fullname));
                friendList.innerHTML = "";
                for (const friend of filteredFriends) {
                    const node = createFriend(state, friend);
                    if (node) friendList.appendChild(node);
                }
            }
            return
        }
    });
}

function handleUpdatePhoto(state: AppState): void {
    const form: HTMLFormElement | null =
        document.querySelector("#profile-form");
    if (!form) return;
    const fileInput: HTMLInputElement | null =
        form.querySelector(`input[type="file"]`);
    if (!fileInput) return;

    fileInput.addEventListener("change", async (_) => {
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
    const inputField: HTMLInputElement | null =
        document.querySelector(`#profile input[type="text"]`);
    const saveButton = document.querySelector("#profile button");
    if (!inputField || !saveButton) return;

    saveButton.addEventListener("click", async () => {
        const fullname = inputField.value.trim();
        const success = await updateFullname(fullname);
        if (!success) return;
        console.log("sukses ganti nama");
        if (state.user) state.user.fullname = fullname;

        // update the user's name in user card
        const userCard = document.querySelector(".user-card .user-name");
        if (userCard) userCard.textContent = fullname;
    });
}
