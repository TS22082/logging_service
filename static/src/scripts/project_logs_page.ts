import { createLogCard } from "./utils/createLogCard";
import { getIdFromUrl } from "./utils/getIdFromUrl";

let eventSource: EventSource;

function connectToStream() {
  console.log("Connecting to stream...");

  const projectId = getIdFromUrl(window.location.href);
  eventSource = new EventSource(`/api/dashboard/logs/${projectId}/stream`);

  eventSource.onopen = function () {
    console.log("✅ Connected to stream");
  };

  eventSource.onmessage = function (event) {
    console.log("📨 Received:", event.data);

    try {
      const data = JSON.parse(event.data);

      if (data.type != "success") return;

      const logsContainer = document.getElementById("logsContainer");
      const logCard = createLogCard(data.data);
      logsContainer?.prepend(logCard);
    } catch (error) {
      console.error("Error parsing message:", error);
    }
  };

  eventSource.onerror = function (event) {
    console.error("❌ Stream error:", event);

    setTimeout(() => {
      if (eventSource.readyState === EventSource.CONNECTING) {
        console.log("🔄 Reconnecting...");
      }
    }, 1000);
  };
}

function disconnectFromStream() {
  if (eventSource) {
    eventSource.close();
    console.log("🔌 Disconnected from stream");
  }
}

document.addEventListener("DOMContentLoaded", function () {
  connectToStream();
});

window.addEventListener("beforeunload", function () {
  disconnectFromStream();
});
