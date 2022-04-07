import React, {useState, useEffect} from 'react';
import axios from "axios";
import {Button, Form, Container, Modal } from 'react-bootstrap'
import Change from './single-change.component';
const Changes = () => {

    const [changes, setChanges] = useState
    const [refreshData, setRefreshData] = useState(false)


    const [addNewChange, setAddNewChange] = useState(false)
    const [newChange, setNewChange] = useState({
        "app": "",
        "date": "",
        "env": "",
        "desc": "",
        "user": ""
    })

    useEffect(() => {
        getAllChanges();
    }, [])

    if(refreshData){
        setRefreshData(false);
        getAllChanges();
    }

    return (
        <div>
            <Container>
                <Button onClick={() => setAddNewChange(true)}>
                    Add New Change
                </Button>
            </Container>
            <Container>
                {changes != null && changes.map((change, i) => (
                    <Change changeData={change}/>
                ))}
            </Container>
            <Modal show={addNewChange} onHide={() => setAddNewChange(false)} centered>
                <Modal.Header closeButton>
                    <Modal.Title>Add Order</Modal.Title>
                </Modal.Header>

                <Modal.Body>
                    <Form.Group>
                        <Form.Label >app</Form.Label>
                        <Form.Control onChange={(event) => {newChange.app = event.target.value}}/>
                        <Form.Label>date</Form.Label>
                        <Form.Control onChange={(event) => {newChange.date = event.target.value}}/>
                        <Form.Label >env</Form.Label>
                        <Form.Control onChange={(event) => {newChange.env = event.target.value}}/>
                        <Form.Label >desc</Form.Label>
                        <Form.Control onChange={(event) => {newChange.desc = event.target.value}}/>
                        <Form.Label >user</Form.Label>
                        <Form.Control onChange={(event) => {newChange.user = event.target.value}}/>
                    </Form.Group>
                    <Button onClick={() => addNewChange}>Add</Button>
                    <Button onClick={() => setAddNewChange(false)}>Cancel</Button>
                </Modal.Body>
            </Modal>

        </div>
    );

    function addSingleChange(){
        setAddNewChange(false)
        var url = "http://localhost:5000/change/create"
        axios.post(url, {
            "app": newChange.app,
            "date": newChange.date,
            "env": newChange.env,
            "desc": newChange.desc,
            "user": newChange.user
        }).then(response => {
            if(response.status == 200){
                setRefreshData(true)
            }
        })
    }

    function getAllChanges(){
        var url = "http://localhost:5000/changes"
        axios.get(url, {
            responseType: 'json'
        }).then(response => {
            if(response.status == 200){
                setChanges(response.data)
            }
        })
    }
}
export default Changes