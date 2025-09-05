export const addDeleteKeyHandler = (btn: Element) => {
  btn.addEventListener("click", () => {
    const popover = document.getElementById("delete-key-popover");
    const apiKeyId = btn.getAttribute("data-id");

    popover && apiKeyId && popover.setAttribute("data-id", apiKeyId);
    popover && popover.showPopover();
  });
};
