const s = (o) => {
    const e = document.createElement("div");
    e.classList.add("logs__card");
    const n = document.createElement("h1"),
      r = document.createElement("p");
    return (
      (n.innerText = o.type),
      (r.innerText = o.message),
      e.appendChild(n),
      e.appendChild(r),
      e
    );
  },
  a = (o) => o.split("/")[o.split("/").length - 1].replace(/\/$/, "");
let t;
function d() {
  console.log("Connecting to stream...");
  const o = a(window.location.href);
  (t = new EventSource(`/api/dashboard/logs/${o}/stream`)),
    (t.onopen = function () {
      console.log("✅ Connected to stream");
    }),
    (t.onmessage = function (e) {
      console.log("📨 Received:", e.data);
      try {
        const n = JSON.parse(e.data);
        if (n.type != "success") return;
        const r = document.getElementById("logsContainer"),
          c = s(n.data);
        r == null || r.prepend(c);
      } catch (n) {
        console.error("Error parsing message:", n);
      }
    }),
    (t.onerror = function (e) {
      console.error("❌ Stream error:", e),
        setTimeout(() => {
          t.readyState === EventSource.CONNECTING &&
            console.log("🔄 Reconnecting...");
        }, 1e3);
    });
}
function i() {
  t && (t.close(), console.log("🔌 Disconnected from stream"));
}
document.addEventListener("DOMContentLoaded", function () {
  d();
});
window.addEventListener("beforeunload", function () {
  i();
});
