import React, { useState } from 'react';
import './../css/url_div.css'
import axios from 'axios';
import Button from '@material-ui/core/Button';
import Snackbar from '@material-ui/core/Snackbar';
import MuiAlert from '@material-ui/lab/Alert';
import { makeStyles } from '@material-ui/core/styles';



export default function Url_Inp(props) {

    function Alert(props) {
        return <MuiAlert elevation={6} variant="filled" {...props} />;
    }

    const useStyles = makeStyles((theme) => ({
        root: {
            width: '100%',
            '& > * + *': {
                marginTop: theme.spacing(2),
            },
        },
    }));


    const classes = useStyles();
    const [open, setOpen] = useState(true);
    const [message, setMessage] = useState("")
    const [status, setStatus] = useState("")
    const handleClose = (event, reason) => {
        if (reason === 'clickaway') {
            return;
        }

        setOpen(false);
    };


    const [long_url, setLongUrl] = useState("")

    const myChangeHandler = (event) => {
        setLongUrl(event.target.value);
    }

    return (

        <div className="Url_inp">

            <input type="text" onChange={myChangeHandler} className="Inp_holder" placeholder="Enter URL Here" />

            <button color="success" onClick={() => {

                const body = {
                    url: long_url
                };

                axios.post(`/api/create`, body, {
                    headers:
                    {
                        'Content-Type': 'application/json',
                    },
                    withCredentials: true
                }).then(response => response.data)
                    .then((data) => {
                        if (data.status != "fail") {
                            setOpen(true)
                            setMessage("Sucess!!")
                            setStatus("success")
                            props.setrefresh(true)
                        }

                        else {
                            setOpen(true)
                            setMessage("Something Went Wrong!!")
                            setStatus("error")
                            console.log(response)

                        }
                    })
                    .catch((error) => {
                        setOpen(true)
                        setMessage("Something Went Wrong!!")
                        setStatus("error")

                        console.log(error)

                    })



            }} className="Btn_holder">Go!</button>
            <Snackbar open={open} autoHideDuration={5000} onClose={handleClose}>
                <Alert onClose={handleClose} severity={status}>
                    {message}
                </Alert>
            </Snackbar>


        </div >
    )
}