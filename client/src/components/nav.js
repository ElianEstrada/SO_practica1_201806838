import React from "react";
import {Menubar} from 'primereact/menubar';

import "./css/nav.css"

const Nav = () => {

    const start = <h1 className="h1_p1">Sopes 1</h1>
    const end = <h2>Elian Estrada</h2>

    return(
        <div className="card">
            <Menubar className="menubar" start={start} end={end}/>
        </div>
    )
}

export default Nav;