type Log = {
  type: string;
  message: string;
};

export const createLogCard = (log: Log) => {
  const logContainer = document.createElement("div");
  logContainer.classList.add("logs__card");

  const logType = document.createElement("h1");
  const logMsg = document.createElement("p");

  logType.innerText = log.type;
  logMsg.innerText = log.message;

  logContainer.appendChild(logType);
  logContainer.appendChild(logMsg);

  return logContainer;
};
