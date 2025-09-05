const docsButtonGroup = document.getElementById("docs_btn_group");

docsButtonGroup &&
  docsButtonGroup.addEventListener("click", (e) => {
    const btn = e.target as HTMLElement;
    const selected = btn.getAttribute("data-selected");

    window.location.href = "/docs/" + selected;
  });
