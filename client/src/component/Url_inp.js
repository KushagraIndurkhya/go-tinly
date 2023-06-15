import React, { useState } from "react";
import "./../css/url_div.css";
import axios from "axios";
import Button from "@material-ui/core/Button";
import InputLabel from "@material-ui/core/InputLabel";
import MenuItem from "@material-ui/core/MenuItem";
import Select from "@material-ui/core/Select";
import FormControl from "@material-ui/core/FormControl";
import Grid from "@material-ui/core/Grid";
import Snackbar from "@material-ui/core/Snackbar";
import MuiAlert from "@material-ui/lab/Alert";
import { makeStyles } from "@material-ui/core/styles";
import { TextField } from "@material-ui/core";

export default function Url_Inp(props) {
  function Alert(props) {
    return <MuiAlert elevation={6} variant="filled" {...props} />;
  }

  const useStyles = makeStyles((theme) => ({
    root: {
      width: "100%",
      "& > * + *": {
        marginTop: theme.spacing(2),
      },
    },
  }));

  const classes = useStyles();
  const [open, setOpen] = useState(false);
  const [message, setMessage] = useState("");
  const [status, setStatus] = useState("error");
  const handleClose = (event, reason) => {
    if (reason === "clickaway") {
      return;
    }

    setOpen(false);
  };

  const [long_url, setLongUrl] = useState("");
  const [apiData, setApiData] = useState({
    Comments: "",
    Medium: "",
    Source: "",
    Campaign: "",
    Keyword: "",
    short: "",
  });

  const myChangeHandler = (event) => {
    setLongUrl(event.target.value);
  };

  const [expiry, setExpiry] = useState(10 * 12 * 30 * 24 * 60 * 60);
  const [openSelect, setOpenSelect] = React.useState(false);

  const handleChangeSelect = (event) => {
    setExpiry(event.target.value);
  };

  const handleCloseSelect = () => {
    setOpenSelect(false);
  };

  const handleOpenSelect = () => {
    setOpenSelect(true);
  };

  return (
    <div className="Url_inp">
      <Grid container spacing={2} alignItems="stretch">
        <Grid
          item
          style={{
            display: "flex",
          }}
          xs={8}
        >
          <input
            type="text"
            onChange={myChangeHandler}
            className="Inp_holder"
            placeholder="Enter URL Here"
          />
        </Grid>

        <Grid
          item
          style={{
            display: "flex",
            alignSelf: "stretch",
          }}
          xs={3}
        >
          <input
            label="Short"
            className="Inp_holder"
            value={apiData.short}
            onChange={(e) => {
              setApiData((prev) => {
                return {
                  ...prev,
                  short: e.target.value,
                };
              });
            }}
            placeholder="short"
          />
        </Grid>
        <Grid
          item
          style={{
            display: "flex",
          }}
          xs={1}
        ></Grid>
      </Grid>
      <div
        style={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          width: "100%",
          gap: "16px",
          paddingTop: "10px",
          flexWrap: "wrap",
        }}
      >
        <input
          label="Comments"
          className="Inp_holder"
          value={apiData.Comments}
          onChange={(e) => {
            setApiData((prev) => {
              return {
                ...prev,
                Comments: e.target.value,
              };
            });
          }}
          placeholder="Comments"
          style={{ width: "40%" }}
        />
        <input
          label="Medium"
          className="Inp_holder"
          value={apiData.Medium}
          onChange={(e) => {
            setApiData((prev) => {
              return {
                ...prev,
                Medium: e.target.value,
              };
            });
          }}
          placeholder="Medium"
          style={{ width: "40%" }}
        />
        <input
          label="Source"
          className="Inp_holder"
          value={apiData.Source}
          onChange={(e) => {
            setApiData((prev) => {
              return {
                ...prev,
                Source: e.target.value,
              };
            });
          }}
          placeholder="Source"
          style={{ width: "40%" }}
        />
        <input
          label="Campaign"
          className="Inp_holder"
          value={apiData.Campaign}
          onChange={(e) => {
            setApiData((prev) => {
              return {
                ...prev,
                Campaign: e.target.value,
              };
            });
          }}
          placeholder="Campaign"
          style={{ width: "40%" }}
        />
        <input
          label="Keyword"
          className="Inp_holder"
          value={apiData.Keyword}
          onChange={(e) => {
            setApiData((prev) => {
              return {
                ...prev,
                Keyword: e.target.value,
              };
            });
          }}
          placeholder="Keyword"
          style={{ width: "40%" }}
        />
        <FormControl className={classes.formControl} style={{ width: "100%" }}>
          <InputLabel
            id="select-label"
            style={{ backgroundColor: "None", color: "white" }}
          >
            Expiry
          </InputLabel>
          <Select
            id="select"
            open={openSelect}
            onClose={handleCloseSelect}
            onOpen={handleOpenSelect}
            value={expiry}
            onChange={handleChangeSelect}
            style={{ backgroundColor: "white", width: "100%" }}
          >
            {/* <MenuItem value=""><em>None</em></MenuItem> */}
            <MenuItem value={1 * 24 * 60 * 60}>1 Day(default)</MenuItem>
            <MenuItem value={7 * 24 * 60 * 60}>7 Days</MenuItem>
            <MenuItem value={30 * 24 * 60 * 60}>30 Days</MenuItem>
            <MenuItem value={10 * 12 * 30 * 24 * 60 * 60}>10 Years</MenuItem>
          </Select>
        </FormControl>
      </div>
      <div
        style={{
          display: "flex",
          justifyContent: "center",
          width: "100%",
          padding: "16px",
        }}
      >
        <Button
          variant="contained"
          color="secondary"
          onClick={() => {
            console.log(expiry);

            const body = {
              url: long_url,
              expiry: expiry,
              ...apiData,
            };

            axios
              .post(`/api/create`, body, {
                headers: {
                  "Content-Type": "application/json",
                },
                withCredentials: true,
              })
              .then((response) => response.data)
              .then((data) => {
                if (data.status != "fail") {
                  setOpen(true);
                  setMessage("Sucess!!");
                  setStatus("success");
                  props.setrefresh(true);
                } else {
                  setOpen(true);
                  setMessage("Something Went Wrong!!");
                  setStatus("error");
                  console.log(response);
                }
              })
              .catch((error) => {
                setOpen(true);
                setMessage("Something Went Wrong!!");
                setStatus("error");

                console.log(error);
              });
          }}
          className="Btn_holder"
          sx={{
            color: "white",
            
          }}
        >
          Go!
        </Button>
      </div>
      <Snackbar open={open} autoHideDuration={5000} onClose={handleClose}>
        <Alert onClose={handleClose} severity={status}>
          {message}
        </Alert>
      </Snackbar>
    </div>
  );
}
