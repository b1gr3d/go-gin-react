import React, {useState, useEffect} from 'react';
import 'bootstrap/dist/css/bootstrap.css';
import {Card, Row, Col} from 'react-bootstrap'
const Change = ({changeData}) => {
    return (
        <Card>
            <Row>
                <Col>App:{ changeData !== undefined && changeData.app}</Col>
                <Col>Date:{ changeData !== undefined && changeData.date}</Col>
                <Col>Env:{ changeData !== undefined && changeData.env}</Col>
                <Col>Description:{ changeData !== undefined && changeData.desc}</Col>
                <Col>User:{ changeData !== undefined && changeData.user}</Col>
            </Row>
        </Card>
    )
}
export default Change