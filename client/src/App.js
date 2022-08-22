import React, { Fragment } from "react";
import { Routes, Route } from "react-router-dom";
import { Home } from "./pages/home";
import Nav from './components/nav';


const App = () => {
  return (
    <Fragment>
      <header>
        <Nav />
      </header>
      <Routes>
        <Route exact path="/" element={<Home />} />
      </Routes>
    </Fragment>
  )
}

export default App;