import { internalRequest } from "./internalRequest";
import { addDeleteKeyHandler } from "./removeKey";

type createKeyRow = {
  id: string;
  project: string;
  count: number;
  dateCreated: string;
  token: string;
};

export const createKeyRow = (props: createKeyRow) => {
  const tr = document.createElement("tr");
  const keyCell = document.createElement("td");
  const requestCountCell = document.createElement("td");
  const deleteCell = document.createElement("td");
  const deleteBtn = document.createElement("button");

  tr.setAttribute("id", "tr-" + props.id);

  let last4OfToken = props.token.slice(-4);
  last4OfToken = "..." + last4OfToken;

  keyCell.innerText = last4OfToken;
  requestCountCell.innerText = props.count.toString();

  deleteBtn.innerText = "Delete";
  deleteBtn.classList.add("delete__link");
  deleteBtn.setAttribute("data-id", props.id);

  deleteBtn.addEventListener("click", async () => {
    try {
      const popover = document.getElementById("delete-key-popover");
      if (!popover) return;
      popover.setAttribute("data-id", props.id);
      popover.showPopover();
    } catch (error) {
      console.log("Error here =>", error);
    }
  });

  deleteCell.appendChild(deleteBtn);
  tr.appendChild(keyCell);
  tr.appendChild(requestCountCell);
  tr.appendChild(deleteCell);

  return tr;
};
