import {
  submittedLoginHeadingText,
  submittedLoginHeadingTextError,
  submittedLoginSupportingText,
  submittedLoginSupportingTextError,
  submitEmailLoginForm,
  loginHeading,
  loginSupport,
  submitLoginBtn,
} from "./utils/constants";

submitEmailLoginForm &&
  submitEmailLoginForm.addEventListener("submit", async (e) => {
    e.preventDefault();

    submitLoginBtn.disabled = true;
    submitLoginBtn.innerText = "Loading";

    const form = e.target as HTMLFormElement;
    const formData = new FormData(form);
    const email = formData.get("email") as string;

    try {
      const body = { email };
      const response = await fetch("/api/send_login_link", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(body),
      });

      const data = await response.json();

      if (!data.success && loginHeading && loginSupport) {
        loginHeading.innerText = submittedLoginHeadingTextError;
        loginSupport.innerText = submittedLoginSupportingTextError;
        return;
      }

      loginHeading && (loginHeading.innerText = submittedLoginHeadingText);
      loginSupport && (loginSupport.innerText = submittedLoginSupportingText);
      submitEmailLoginForm &&
        submitEmailLoginForm.classList.add("hide__visibility");
    } catch (error) {
      loginHeading && (loginHeading.innerText = submittedLoginHeadingTextError);
      loginSupport &&
        (loginSupport.innerText = submittedLoginSupportingTextError);
    }
  });
