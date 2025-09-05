const learnMoreBtn = document.getElementById("learnMoreBtn");
const loginBtn = document.getElementById("logInBtn");
const createAccountBtn = document.getElementById("createAccountBtn");

learnMoreBtn &&
  learnMoreBtn.addEventListener("click", () => {
    window.location.href = "/docs/accounts";
  });

[loginBtn, createAccountBtn].forEach((element) => {
  element &&
    element.addEventListener("click", () => {
      window.location.href = "/login";
    });
});
