import React,{Link}from 'react'
import GoogleButton from 'react-google-button';
import './../css/App.css'
export default function Login() {
    return (
        <div className="App">
        <p style={{
          color: "white",
          fontWeight: "bold",
          padding: "10px",
          fontSize: "80px",
          fontFamily: "Roboto"

        }} > Go-tinly</p>

        <p style={{
          color: "white",
          padding: "10px",
          fontSize: "30px",
          fontFamily: "Roboto",
        }} > Shorten your URLs</p>

        <div className="ButtonDiv">
          <a href="/login">
            <GoogleButton /></a>
        </div>
      </div>
    )
}
