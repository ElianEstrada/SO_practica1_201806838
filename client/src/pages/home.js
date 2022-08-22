import React from "react";
import Table from "../components/table";

import "./css/home.css";

export const Home = (props) => {
    return (
        <div className="container">
            <div className="content">
                <h1>Registered Vehicles </h1>
            </div>
            <div className="home_table">
                <Table />
            </div>
        </div>
    );
}