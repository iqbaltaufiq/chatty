import { addFriend, removeFriend } from "../api.js";
import type { AppState, Friend, User } from "../types.js";
import { setActiveChatroom } from "./chatroom.js";

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
        const success = await addFriend(user.username);
        if (success) return window.location.reload();
    });
    return doc.body.firstChild;
}
