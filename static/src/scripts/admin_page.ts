import {
  accountOptionInputs,
  planDescriptionText,
  planDescriptions,
  newProjectForm,
  projectTableBody,
  deleteProjectPopover,
  cancelDeleteProject,
  confirmDeleteBtn,
  confirmDeleteProject,
} from "./utils/constants";
import { createProjectRow } from "./utils/createProjectRow";
import { internalRequest } from "./utils/internalRequest";

accountOptionInputs.forEach((option) => {
  option.addEventListener("change", (event) => {
    const target = event.target as HTMLInputElement;
    const selectedValue = target.value;

    planDescriptionText &&
      (planDescriptionText.innerText =
        planDescriptions[selectedValue as keyof typeof planDescriptions]);
  });
});

newProjectForm &&
  newProjectForm.addEventListener("submit", async function (e) {
    e.preventDefault();

    const selected = Array.from(accountOptionInputs).find((option) => {
      return option.checked;
    }) as HTMLInputElement;

    const selectedPlan = selected ? selected.value : "basic";
    const formInput = newProjectForm["projectInput"];

    if (!formInput.value) {
      planDescriptionText &&
        (planDescriptionText.innerText = "Project name cannot be blank");

      setTimeout(() => {
        planDescriptionText &&
          (planDescriptionText.innerText =
            planDescriptions[selectedPlan as keyof typeof planDescriptions]);
      }, 1000);
      return;
    }

    try {
      const requestBody = {
        project: formInput.value,
        plan: selectedPlan,
      };

      const data = await internalRequest.Post("/api/project", requestBody);

      console.log("Data ==>", data);

      const { id, name } = data;
      const newTableRow = createProjectRow({ id, name, plan: selectedPlan });

      projectTableBody && projectTableBody.appendChild(newTableRow);
    } catch (error) {
      console.log("Error ==>", error);
    } finally {
      formInput.value = "";
    }
  });

projectTableBody?.addEventListener("click", (event) => {
  const target = event.target as HTMLElement | null;
  const btn = target?.closest("button");

  if (btn && deleteProjectPopover) {
    if (btn.classList.contains("delete__link")) {
      const projectId = btn.getAttribute("data-project-id");
      if (typeof projectId != "string") return;

      deleteProjectPopover.setAttribute("data-delete-id", projectId);
      deleteProjectPopover.showPopover();
    }
  }
});

deleteProjectPopover?.addEventListener("toggle", (event) => {
  if (event.newState == "closed") {
    deleteProjectPopover?.setAttribute("data-delete-id", "");
  }
});

cancelDeleteProject?.addEventListener("click", () => {
  deleteProjectPopover?.hidePopover();
});

confirmDeleteProject?.addEventListener("click", async (event) => {
  const projectDeleteId = deleteProjectPopover?.getAttribute("data-delete-id");

  if (!projectDeleteId || typeof projectDeleteId !== "string") return;

  try {
    await internalRequest.Delete(`/api/project/${projectDeleteId}`);

    const projectRow = document.getElementById("project_" + projectDeleteId);
    projectRow?.remove();

    deleteProjectPopover?.hidePopover();
  } catch (error) {
    console.log("Error =>", error);
  }
});
