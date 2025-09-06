import { dashboardProjects } from "./utils/constants";

dashboardProjects?.addEventListener("click", (e) => {
  const target = e.target as HTMLElement | null;
  const btn = target?.closest(".btn");
  const projectId = btn?.getAttribute("data-id");

  if (projectId) {
    window.location.href = `/dashboard/logs/${projectId}`;
  }
});
