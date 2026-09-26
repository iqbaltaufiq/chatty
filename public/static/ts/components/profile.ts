import type { AppState } from "../types.js";

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
