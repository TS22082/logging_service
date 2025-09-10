type createProjectsRowProps = {
  name: string;
  plan: string;
  id: string;
};

export const createProjectRow = (props: createProjectsRowProps) => {
  const tr = document.createElement("tr");
  const td_1 = document.createElement("td");
  const anchor_1 = document.createElement("a");
  const td_2 = document.createElement("td");
  const td_3 = document.createElement("td");
  const td_4 = document.createElement("td");
  const deleteBtn_4 = document.createElement("button");
  const planTd = document.createElement("td");

  anchor_1.innerText = props.name;
  anchor_1.href = `/project/${props.id}`;
  anchor_1.classList.add("primary__link");
  td_2.innerText = "1";
  td_3.innerText = "0";
  deleteBtn_4.innerText = "Delete";
  deleteBtn_4.setAttribute("data-project-id", props.id);
  deleteBtn_4.classList.add("delete__link");
  planTd.innerText = props.plan;

  td_1.appendChild(anchor_1);
  td_4.appendChild(deleteBtn_4);

  tr.setAttribute("id", "project_" + props.id);
  tr.appendChild(td_1);
  tr.appendChild(planTd);
  tr.appendChild(td_2);
  tr.appendChild(td_3);
  tr.appendChild(td_4);

  return tr;
};
