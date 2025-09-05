export const getIdFromUrl = (url: string) => {
  const projectId = url.split("/")[url.split("/").length - 1];
  return projectId.replace(/\/$/, "");
};
