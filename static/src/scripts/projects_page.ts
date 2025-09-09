import {
  accountOptionInputs,
  planDescriptionText,
  planDescriptions,
  newProjectForm,
  projectTableBody,
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

      const data = await internalRequest.Post("/project", requestBody);
      const { id, name } = data;
      const newTableRow = createProjectRow({ id, name, plan: selectedPlan });

      projectTableBody && projectTableBody.appendChild(newTableRow);
    } catch (error) {
      console.log("Error ==>", error);
    } finally {
      formInput.value = "";
    }
  });
