import React, { useEffect, useState } from "react";
import { Card } from "primereact/card";
import { DataTable } from "primereact/datatable";
import { Column } from "primereact/column";
import "./css/table.css";

const Table = () => {

    const [vehicles, setVehicles] = useState([]);

    useEffect(() => {
        fetch('http://localhost:4000/vehicle').then(res => res.json()).then(data => setVehicles(data))
    }, []);

    const actions = () => {
        return (
            <div className="icons">
                <a href="##"><i className="pi pi-pencil mr-5" style={{ color: "green" }}></i></a>
                <a href="##"><i className="pi pi-trash" style={{ color: "red" }}></i></a>
            </div>
        )
    }

    return (
        <div>
            <Card className="div_table">
                <DataTable value={vehicles} responsiveLayout="scroll">
                    <Column field="Plate" header="Plate"></Column>
                    <Column field="Make" header="Make"></Column>
                    <Column field="Model" header="Model"></Column>
                    <Column field="Series" header="Series"></Column>
                    <Column field="Color" header="Color"></Column>
                    <Column field="actions" header="Actions" body={actions}></Column>
                </DataTable>
            </Card>
        </div>
    )
}

export default Table;