import React, { useState, useEffect, Link } from "react";
import axios from "axios";
import Loading from "./component/Loading";
import Login from "./views/login";
import Url_Inp from "./component/Url_inp";
import Header from "./component/header";
import Dashboard from "./views/dashboard";

const BASE_URL = "http://localhost:8080/";

export default function App() {
  const [isLoggedIn, setisLoggedIn] = useState(1);
  const [info, setinfo] = useState([]);
  const [urli, seturli] = useState([]);
  const [refresh, setrefresh] = useState(true);

  useEffect(() => {
    setrefresh(false);
    const url =
      // `${BASE_URL}`+
      "api/dash";
    axios
      .get(url, {
        headers: {
          "Content-Type": "application/json",
        },
        withCredentials: true,
      })
      .then((response) => response.data)
      .then((data) => {
        console.log(data);
        if (data.status != "fail") {
          console.log(data);
          setisLoggedIn(1);
          setinfo(data.userInfo);
          seturli(data.urls);
        
        }
      })
      .catch((error) => {
        console.log("error");
        setisLoggedIn(1);
      });
  }, [refresh]);
  switch (isLoggedIn) {
    case -1:
      return <Loading />;
    case 1:
      return (
        <div>
          <Header />
          <Url_Inp setrefresh={setrefresh} />
          <Dashboard urls={urli} />{" "}
        </div>
      );

    case 0:
      return <Login />;
  }
}
