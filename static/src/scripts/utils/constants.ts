export const submitEmailLoginForm = document.getElementById("emailLoginForm");
export const loginHeading = document.getElementById("loginHeading");
export const loginSupport = document.getElementById("loginSupport");
export const submitLoginBtn = document.getElementById(
  "submitLoginBtn"
) as HTMLButtonElement;

export const submitEmailNotifyForm = document.getElementById(
  "submitEmailNotifyForm"
);
export const submitEmailHeading = document.getElementById("signUpHeading");
export const submitEmailSupport = document.getElementById("signUpSupport");

export const accountOptionInputs = document.querySelectorAll(
  'input[name="account__option"]'
) as NodeListOf<HTMLInputElement>;

export const dashboardProjects = document.getElementById("dashboardProjects");

export const planDescription = document.getElementById("plan_description");
export const projectTableBody = document.getElementById("projectTableBody");

export const planDescriptionText = document.getElementById(
  "planDescriptionText"
);

export const newProjectForm = document.getElementById(
  "newProjectForm"
) as HTMLFormElement;

export const cancelDeleteButton = document.getElementById("cancelDeleteButton");
export const deleteProjectButton = document.getElementById(
  "deleteProjectButton"
);

export const deleteProjectPopover = document.getElementById(
  "delete-project-popover"
);

export const cancelDeleteProject = document.getElementById(
  "cancelDeleteProject"
);

export const confirmDeleteProject = document.getElementById(
  "confirmDeleteProject"
);

export const createApiKeyBtn = document.getElementById("createApiKey");
export const tableBody = document.getElementById("keyTable");
export const deleteBtns = document.querySelectorAll(".delete__link");
export const cancelDeleteBtn = document.getElementById("cancelDeleteKey");
export const confirmDeleteBtn = document.getElementById("confirmDeleteBtn");
export const popover = document.getElementById("delete-key-popover");
export const showKeyModal = document.getElementById(
  "showKeyModal"
) as HTMLDialogElement;
export const keyText = document.getElementById("keyText");
export const keyModalBtn = document.getElementById("closeKeyModalBtn");
export const submitKeyBtn = document.getElementById("submitKeyBtn");

export const logsContainer = document.getElementById("logsContainer");

export const submittedHeadingText = "Excellent!";
export const submittedSupportingText =
  "You've signed up to be notified when this service is available";

export const submittedHeadingTextError = "Something went wrong!";
export const submittedSupportingTextError =
  "We may be having issues with our servers, please try again later.";

export const submittedLoginHeadingText = "Excellent!";
export const submittedLoginSupportingText =
  "We've sent you a magic link. Check your email and follow the link inside to get logged in";

export const submittedLoginHeadingTextError = "Something went wrong!";
export const submittedLoginSupportingTextError =
  "We may be having issues with our servers, please try again later.";

export const planDescriptions = {
  basic: "Basic: Supports 1 user, 1 key and up to 100 logs p/month",
  team: "Team: Supports up to 5 users, 3 keys and 300 logs p/month",
  pro: "Pro: Unlimited users, and keys. Up to 500 logs p/month",
};

export const JSON_HEADERS = {
  "Content-Type": "application/json",
};
