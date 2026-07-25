import { initApp } from "./init.js"
import type { AppState } from "./types.ts";

const state: AppState = {
    user: null,
    room: null,
    friends: [],
    recent_chats: []
};

initApp(state);
