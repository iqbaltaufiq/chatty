import { login } from "../api.js";

document.addEventListener("DOMContentLoaded", () => {
    handleSubmit();
});

function handleSubmit(): void {
    const form: HTMLFormElement | null =
        document.querySelector("form");
    if (!form) return;
    form.addEventListener("submit", async (ev) => {
        ev.preventDefault();

        const formInputs = document.getElementsByTagName("input");
        const username =
            formInputs[0]?.name === "username" ?
            formInputs[0].value :
            "";
        const password =
            formInputs[1]?.name === "password" ?
            formInputs[1].value :
            "";
        const isSuccess = await login(username, password);
        if (isSuccess) {
            window.location.assign("/");
            return
        }
        if (!isSuccess) {
            document.querySelector(".error")?.remove();
            const err = document.createElement("p");
            err.className = "error";
            err.textContent =
                "Please check your username and password again.";
            form.prepend(err);
        }
        form.reset();
    });
}
