import {
  cancelDeleteBtn,
  confirmDeleteBtn,
  createApiKeyBtn,
  deleteBtns,
  keyModalBtn,
  keyText,
  popover,
  showKeyModal,
  submitKeyBtn,
  tableBody,
} from "./utils/constants";
import { createKeyRow } from "./utils/createKeyRow";
import { internalRequest } from "./utils/internalRequest";
import { addDeleteKeyHandler } from "./utils/removeKey";

createApiKeyBtn &&
  createApiKeyBtn.addEventListener("submit", async (e) => {
    e.preventDefault();

    submitKeyBtn && (submitKeyBtn.innerText = "Loading ...");

    const id =
      document.location.href.split("/")[
        document.location.href.split("/").length - 1
      ];

    try {
      const requestData = {
        projectId: id,
      };

      const { data } = await internalRequest.Post("/api/key", requestData);
      console.log("Data ==>", data);

      const newKeyRow = createKeyRow(data);

      showKeyModal && showKeyModal.showModal();
      keyText && (keyText.innerText = data.token);
      tableBody && tableBody.appendChild(newKeyRow);
    } catch (error) {
      console.log("Error ==>", error);
    } finally {
      submitKeyBtn && (submitKeyBtn.innerText = "Create API Key");
    }
  });

deleteBtns && deleteBtns.forEach((btn) => addDeleteKeyHandler(btn));

cancelDeleteBtn &&
  cancelDeleteBtn.addEventListener("click", () => {
    const popover = document.getElementById("delete-key-popover");
    popover && popover.hidePopover();
  });

confirmDeleteBtn &&
  confirmDeleteBtn.addEventListener("click", async function () {
    if (!popover) return;

    const keyId = popover.getAttribute("data-id");
    const rowToDelete = document.getElementById("tr-" + keyId);

    if (!rowToDelete) return;

    try {
      const requestUrl = `/api/key/${keyId}`;
      const response = await internalRequest.Delete(requestUrl);

      if (response.success) {
        rowToDelete.remove();
        popover.removeAttribute("data-id");
        popover.hidePopover();
      } else {
        throw new Error("Error in response");
      }
    } catch (error) {
      console.log("Error ==>", error);
    }
  });

keyModalBtn &&
  keyModalBtn.addEventListener("click", () => {
    showKeyModal.close();
  });
