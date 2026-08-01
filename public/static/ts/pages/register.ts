import { register } from "../api.js";
import { MAX_FILE_SIZE } from "../types.js";

document.addEventListener("DOMContentLoaded", () => {
    handleRegister();
});

function handleRegister(): void {
    const form: HTMLFormElement | null =
        document.querySelector("form");
    if (!form) return;
    
    form.addEventListener("submit", async (ev) => {
        ev.preventDefault();

        const fileInput: HTMLInputElement | null =
            document.querySelector(`input[type="file"]`);
        const file = fileInput?.files ? fileInput.files[0] : null;
        if (file) {
            if (file.size > MAX_FILE_SIZE) {
                const message = "file exceeds maximum size allowed";
                displayError(message);
                throw new Error(message);
            }
        }
        const formData = new FormData(form);
        try {
            const ok = await register(formData);
            if (!ok) throw new Error("failed creating your account");
            form.reset();
            hideError();
        } catch (error) {
            displayError("Failed creating your account.");
            console.error(error);
            return
        }
    });
}

function displayError(message: string): void {
    const errMessage: HTMLElement | null = document.querySelector("p.error");
    if (!errMessage) return;
    errMessage.textContent = message;
    errMessage.style.display = "block";
}

function hideError(): void {
    const errMessage: HTMLElement | null = document.querySelector("p.error");
    if (!errMessage) return;
    errMessage.textContent = "";
    errMessage.style.display = "none";
}
