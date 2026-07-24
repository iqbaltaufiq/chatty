import type { AppState } from "./types.js";
import { getUser } from "./api.js";
import { registerHandler } from "./handler.js";
import { toggleConnection } from "./websocket.js";

export async function initApp(state: AppState): Promise<void> {
    await initUser(state);
    registerHandler(state);
    toggleConnection(state);
}

export async function initUser(state: AppState): Promise<void> {
    try {
        const user = await getUser();
        if (!user) throw new Error("failed initializing user");
        state.user = user;
    } catch (error) {
        console.error(error);
    }

    const photoExists = `/static/assets/${state.user?.photo}`;
    const photoAbsent = `/static/assets/default.png`;
    const photo: HTMLImageElement | null = document.querySelector(".user-photo > img");
    if (photo) photo.src = photoExists || photoAbsent;

    const userName: HTMLParagraphElement | null = document.querySelector(".user-name");
    if (userName) userName.textContent = state.user?.fullname || "unknown";
}

