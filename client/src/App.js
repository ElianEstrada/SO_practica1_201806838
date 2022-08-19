import React from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";


const App = () => {
  return (
    <BrowserRouter>
      <Switch>
        <Route exact path="/" />
        <Route exact path="**" />
      </Switch>
    </BrowserRouter>
  )
}

export default App;