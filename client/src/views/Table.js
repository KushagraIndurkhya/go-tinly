import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import Paper from "@material-ui/core/Paper";

const useStyles = makeStyles({
  table: {
    minWidth: 650,
    "& .MuiTableCell-root": {
      color: "black",
      fontWeight: "600",
      fontSize: "1rem",
      fontStyle: "normal",
    },
    tableCell: {
      color: "red",
    },
  },
});

// function createData(name, calories, fat, carbs, protein) {
//   return { name, calories, fat, carbs, protein };
// }

// const rows = [
//   createData("Frozen yoghurt", 159, 6.0, 24, 4.0),
//   createData("Ice cream sandwich", 237, 9.0, 37, 4.3),
//   createData("Eclair", 262, 16.0, 24, 6.0),
//   createData("Cupcake", 305, 3.7, 67, 4.3),
//   createData("Gingerbread", 356, 16.0, 49, 3.9),
// ];

export default function DenseTable(props) {
  const classes = useStyles();
  const BASE =
    // process.env.BASE_URL
    "https://go-tinly.onrender.com" + "/r/";

  return (
    <TableContainer component={Paper}>
      <Table className={classes.table} size="small" aria-label="a dense table">
        <TableHead>
          <TableRow>
            <TableCell>URL</TableCell>
            <TableCell align="right">Short URL</TableCell>
            <TableCell align="right">Hits</TableCell>
            <TableCell align="right">Comments</TableCell>
            <TableCell align="right">Medium</TableCell>
            <TableCell align="right">Source</TableCell>
            <TableCell align="right">Campaign</TableCell>
            <TableCell align="right">Keyword</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {props.rows.map((row) => (
            <TableRow key={row.name}>
              <TableCell component="th" scope="row">
                {row.url}
              </TableCell>
              <TableCell align="right">
                {" "}
                <a href={`${BASE}` + row.short_url}>{row.short_url}</a>
              </TableCell>
              <TableCell align="right">{row.Hits}</TableCell>
              <TableCell align="right">{row.Comments}</TableCell>
              <TableCell align="right">{row.Medium}</TableCell>
              <TableCell align="right">{row.Source}</TableCell>
              <TableCell align="right">{row.Campaign}</TableCell>
              <TableCell align="right">{row.Keyword}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
