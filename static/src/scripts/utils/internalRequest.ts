import { JSON_HEADERS } from "./constants";

export const internalRequest = {
  Post: async (url: string, data: any) => {
    const response = await fetch(url, {
      method: "POST",
      body: JSON.stringify(data),
      headers: JSON_HEADERS,
    });

    const responseData = await response.json();
    return responseData;
  },
  Put: async (url: string) => {
    const response = await fetch(url, {
      method: "PUT",
      headers: JSON_HEADERS,
    });

    const responseData = await response.json();
    return responseData;
  },
  Delete: async (url: string) => {
    const response = await fetch(url, {
      method: "DELETE",
      headers: JSON_HEADERS,
    });

    const responseData = await response.json();
    return responseData;
  },
};
