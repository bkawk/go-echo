import React from "react";
import CheckHealth from "./components/checkHealth";
import CheckUsername from "./components/checkUsername";
import DoRegister from "./components/doRegister";

const App = () => {
  return (
    <div>
      <CheckHealth />
      <CheckUsername />
      <DoRegister />
    </div>
  );
};

export default App;
