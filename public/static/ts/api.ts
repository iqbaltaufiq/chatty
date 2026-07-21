import type { User } from "./types.ts";

export async function getUser(): Promise<User | null> {
    const response = await fetch("/api/user/get");
    const json = await response.json();

    if (!response.ok) {
        window.location.href = "/login";
        console.log("we cannot find your record");
        return null
    }
    const user: User = json.data?.user;
    return user;
}
